package bind

import (
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/network"
	tpc "github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	u "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type Node struct {
	// Properties
	call    Callback
	config  nodeConfig
	connreq *md.ConnectionRequest

	// Client
	client *sc.Client
	user   *u.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *Node {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "http://8f37928df15e41318ebc28770270da05@ec2-34-201-54-61.compute-1.amazonaws.com/2",
	})

	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
		return nil
	}

	// Check Device
	if req.Device.Platform == md.Platform_IOS || req.Device.Platform == md.Platform_Android {
		// Create Mobile Node
		mn := &Node{
			call:    call,
			config:  newNodeConfig(),
			connreq: req,
			topics:  make(map[string]*tpc.TopicManager, 10),
		}

		// Create New User
		mn.user = u.NewUser(req, mn.callbackNode())

		// Create Client
		mn.client = sc.NewClient(mn.contextNode(), req, mn.callbackNode())
		return mn
	} else {
		// Get Location by IP
		geoIP := md.GeoIP{}
		err := network.Location(&geoIP)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Finding Geolocated IP"))
			return nil
		}

		// Modify Request
		req.AttachGeoToRequest(&geoIP)

		// Create Mobile Node
		mn := &Node{
			call:    call,
			config:  newNodeConfig(),
			connreq: req,
			topics:  make(map[string]*tpc.TopicManager, 10),
		}

		// Create New User
		mn.user = u.NewUser(req, mn.callbackNode())

		// Create Client
		mn.client = sc.NewClient(mn.contextNode(), req, mn.callbackNode())
		return mn
	}
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host and Connect
func (mn *Node) Connect() {
	if !mn.config.HasConnected {
		// Connect Host
		err := mn.client.Connect(mn.user.PrivateKey())
		if err != nil {
			mn.handleError(err)
			mn.setConnected(false)
			return
		} else {
			// Update Status
			mn.setConnected(true)
		}

		// Bootstrap Node
		err = mn.client.Bootstrap()
		if err != nil {
			mn.handleError(err)
			mn.setBootstrapped(false)
			return
		} else {
			mn.setBootstrapped(true)
		}

		// Join Local Lobby
		mn.local, err = mn.client.JoinLocal()
		if err != nil {
			mn.handleError(err)
			mn.setJoinedLocal(false)
			return
		} else {
			mn.setJoinedLocal(true)
		}
	}

}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
func (n *Node) NetworkSwitch() {

}

// @ Close Ends All Network Communication
func (mn *Node) Pause() {
	md.GetState().Pause()
}

// @ Close Ends All Network Communication
func (mn *Node) Resume() {
	md.GetState().Resume()
}

// @ Close Ends All Network Communication
func (mn *Node) Stop() {
	mn.client.Close()
	mn.config.Ctx.Done()
}
