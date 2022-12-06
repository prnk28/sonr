package node

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"

	ps "github.com/libp2p/go-libp2p-pubsub"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/sonr-io/sonr/pkg/common"
)

var (
	libp2pRendevouz = "/sonr/rendevouz/0.9.2"
)

// Node returns the Node for the Main Node
type Node interface {
	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// ID returns the node.ID of the host
	ID() (ID, error)

	// Host returns the Host
	Host() host.Host

	// NewStream opens a new stream to a peer
	NewStream(pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// Joins a Channel interface with an underlying pubsub topic and event handler
	Join(name string, opts ...ChannelOption) (*Channel, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Send(target string, data interface{}, protocol protocol.ID) error
	Send(id peer.ID, p protocol.ID, data []byte) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)
}

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	callback common.MotorCallback
	host     host.Host

	// Host and context
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	idht *dht.IpfsDHT
	ps   *pubsub.PubSub

	// State
	fsm *SFSM
}

// NewMotor Creates a Sonr libp2p Host with the given config
func New(ctx context.Context, options ...NodeOption) (Node, error) {
	// Default options
	var err error
	config := defaultNodeConfig()
	for _, o := range options {
		err = o(config)
		if err != nil {
			return nil, err
		}
	}

	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		fsm:          NewFSM(ctx),
		mdnsPeerChan: make(chan peer.AddrInfo),
	}

	hn.host, err = libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(config.GetPrivateKey()),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9000",      // regular tcp connections
			"/ip4/0.0.0.0/udp/9000/quic", // a UDP endpoint for the QUIC transport
		),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(config.ConnManager),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			hn.idht, err = dht.New(ctx, h)
			return hn.idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)
	if err != nil {
		return nil, err
	}

	// setup dht peer discovery
	err = hn.createDHTDiscovery()
	if err != nil {
		return nil, err
	}
	return hn, nil
}
