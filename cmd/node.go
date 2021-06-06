package bind

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	tpc "github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type Node struct {
	md.NodeCallback
	// Properties
	call Callback
	ctx  context.Context

	// Client
	user   *md.User
	client *sc.Client

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager

	// Storage
	storageEnabled bool
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *Node {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})

	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
		return nil
	}
	// Create Mobile Node
	mn := &Node{
		call:   call,
		ctx:    context.Background(),
		topics: make(map[string]*tpc.TopicManager, 10),
	}

	// Create Users
	mn.user = md.NewUser(req)

	// Create Client
	mn.client = sc.NewClient(mn.ctx, mn.user, mn.callbackNode())
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host and Connect
func (mn *Node) Connect() {
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
	mn.local, err = mn.client.Bootstrap()
	if err != nil {
		mn.handleError(err)
		mn.setAvailable(false)
		return
	} else {
		mn.setAvailable(true)
	}
}

// @ Returns Node Location Protobuf as Bytes
func (mn *Node) Location() []byte {
	bytes, err := proto.Marshal(mn.user.Location)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Returns Node User Protobuf as Bytes
func (mn *Node) User() []byte {
	bytes, err := proto.Marshal(mn.user)
	if err != nil {
		return nil
	}
	return bytes
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
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
	mn.ctx.Done()
}