package transfer

import (
	"io"
	"os"
	sync "sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"google.golang.org/protobuf/proto"
)

// onInviteResponse response handler
func (p *TransferProtocol) onInviteResponse(s network.Stream) {
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Error("Failed to Read Invite RESPONSE buffer.", err)
		return
	}
	s.Close()

	// Unmarshal response
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite RESPONSE buffer.", err)
		return
	}
	p.emitter.Emit(Event_RESPONDED, resp.ToEvent())

	// Locate request data and remove it if found
	entry, err := p.queue.Validate(resp)
	if err != nil {
		logger.Error("Failed to Validate Invite RESPONSE buffer.", err)
		return
	}

	// Create a new stream
	stream, err := p.host.NewStream(p.ctx, remotePeer, SessionPID)
	if err != nil {
		logger.Error("Failed to create new stream.", err)
		return
	}

	// Call Outgoing Transfer
	go p.onOutgoingTransfer(entry, msgio.NewWriter(stream))
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransferProtocol) onOutgoingTransfer(entry *Session, wc msgio.WriteCloser) {
	// Initialize Params
	logger.Info("Beginning Outgoing Transfer Stream")
	wg := sync.WaitGroup{}

	// Write All Files
	err := entry.MapItems(func(m *common.Payload_Item, i int, count int) error {
		// Add to WaitGroup
		wg.Add(1)

		// Create New Writer
		w, err := NewWriter(m, i, count, p.emitter)
		if err != nil {
			return logger.Error("Failed to create new writer.", err)
		}

		// Write File to Stream
		if err := w.WriteTo(wc); err != nil {
			return logger.Error("Error writing stream", err)
		}

		// Complete Writing
		wg.Done()
		return nil
	})
	if err != nil {
		logger.Error("Error writing stream", err)
		return
	}

	// Complete the transfer
	wg.Wait()
	event, err := p.queue.Done()
	if err != nil {
		logger.Error("Failed to Complete Transfer", err)
		return
	}

	// Emit Event
	p.emitter.Emit(Event_COMPLETED, event)
}

// itemWriter is a Writer for FileItems
type itemWriter struct {
	chunker *config.Chunker
	emitter *state.Emitter
	file    *os.File
	item    *common.FileItem
	index   int
	count   int
	size    int64
}

// NewWriter Returns a new Writer for the given FileItem
func NewWriter(pi *common.Payload_Item, index int, count int, em *state.Emitter) (*itemWriter, error) {
	// Properties
	size := pi.GetSize()
	item := pi.GetFile()

	// Define Chunker Opts
	var avgSize int
	if size < ITEM_INTERVAL {
		avgSize = int(size)
	} else {
		avgSize = int(size / ITEM_INTERVAL)
	}

	// Open Os File
	f, err := os.Open(item.Path)
	if err != nil {
		return nil, logger.Error("Error opening item for Write stream", err)
	}

	// Create New Chunker
	chunker, err := config.NewChunker(f, config.ChunkerOptions{
		AverageSize: avgSize, // Only Average Required
	})
	if err != nil {
		return nil, logger.Error("Failed to create new chunker.", err)
	}

	// Create New Writer
	return &itemWriter{
		item:    item,
		size:    size,
		file:    f,
		emitter: em,
		index:   index,
		count:   count,
		chunker: chunker,
	}, nil
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress(i int) {
	// Create Progress Event
	event := &common.ProgressEvent{
		Progress: (float64(i) / float64(p.size)),
		Current:  int32(p.index),
		Total:    int32(p.count),
	}

	// Push ProgressEvent to Emitter
	p.emitter.Emit(Event_PROGRESS, event)
}

// Write Item to Stream
func (iw *itemWriter) WriteTo(writer msgio.WriteCloser) error {
	// Loop through File
	for i := 0; i < int(iw.size); {
		c, err := iw.chunker.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Create Block Protobuf from Chunk
		data, err := encodeChunk(c)
		if err != nil {
			return logger.Error("Error Encoding chunk", err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(data)
		if err != nil {
			return logger.Error("Error Writing data to msgio.Writer", err)
		}

		// Unexpected Error
		if err != nil && err != io.EOF {
			return logger.Error("Unexpected Error occurred on Write Stream", err)
		}

		// Update Progress
		i += c.Length

		// Emit Progress
		if (i % ITEM_INTERVAL) == 0 {
			iw.Progress(i)
		}
	}

	// Close File
	if err := iw.file.Close(); err != nil {
		return logger.Error("Failed to Close item on Write Stream", err)
	}
	return nil
}

// encodeChunk Encodes a Chunk into a Protobuf
func encodeChunk(c config.Chunk) ([]byte, error) {
	// Create Block Protobuf from Chunk
	data, err := proto.Marshal(&Chunk{
		Offset:      int32(c.Offset),
		Length:      int32(c.Length),
		Data:        c.Data,
		Fingerprint: int64(c.Fingerprint),
	})

	if err != nil {
		return nil, logger.Error("Error Marshalling Chunk Proto.", err)
	}
	return data, nil
}