package mpc

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
	"github.com/sonr-io/sonr/internal/node"
)

func handlerLoopChannel(id party.ID, h protocol.Handler, channel *node.Channel) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				channel.Close()
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			buf, err := msg.MarshalBinary()
			checkErr(err)
			go channel.Send(buf)

			// incoming messages
		case msg := <-channel.NextMessage():
			var m protocol.Message
			err := m.UnmarshalBinary(msg.Data)
			checkErr(err)
			h.Accept(&m)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		golog.Warnf("error: %v", err) // nolint: errcheck
	}
}
