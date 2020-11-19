package sonr

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	CTX        context.Context
	Host       host.Host
	Lobby      lobby.Lobby
	Profile    pb.Profile
	Contact    pb.Contact
	AuthStream authStreamConn
	Callback   Callback
}

func (sn *Node) GetPeerInfo() *pb.PeerInfo {
	return &pb.PeerInfo{
		PeerId:     sn.Host.ID().String(),
		Device:     sn.Profile.Device,
		FirstName:  sn.Contact.FirstName,
		LastName:   sn.Contact.LastName,
		ProfilePic: sn.Contact.ProfilePic,
		Direction:  sn.Profile.Direction,
	}
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() []byte {
	// Create User Object
	user := &pb.ConnectedMessage{
		HostId:  sn.Profile.HostId,
		Profile: &sn.Profile,
		Contact: &sn.Contact,
	}

	// Marshal to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Send Callback with Available Peers
	sn.Callback.OnRefreshed(sn.Lobby.GetAllPeers())

	// Return as JSON String
	return data
}

// ^ Message Emitter ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(dir float64) bool {
	// Update User Values
	sn.Profile.Direction = math.Round(dir*100) / 100

	// Create Message with Updated Info
	notif := &pb.LobbyMessage{
		Event:  "Update",
		Sender: sn.Profile.HostId,
		Data:   sn.GetPeerInfo(),
	}

	// Inform Lobby
	err := sn.Lobby.Publish(notif)
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}

	// Send Callback with Available Peers
	sn.Callback.OnRefreshed(sn.Lobby.GetAllPeers())

	// Return Success
	return true
}

// Invite an available peer to transfer
func (sn *Node) Invite(data []byte) bool {
	// Initialize
	var wg sync.WaitGroup
	invite := pb.InviteEvent{}
	err := proto.Unmarshal(data, &invite)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}

	// ** Get Required Data **
	peerID, err := sn.Lobby.GetPeerID(invite.PeerId)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}

	// Create Metadata
	meta := file.GetMetadata(invite.FilePath)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		return false
	}

	// Initialize
	var thumb []byte
	wg.Add(1)

	// Create Thumb Nail
	go func(thumb []byte) {
		defer wg.Done()
		thumb = file.GetThumbnail(&wg, meta)
		wg.Wait()
		fmt.Println("Raw Thumbnail: ", thumb)
	}(thumb)

	// ** Create New Auth Stream **
	stream, err := sn.Host.NewStream(sn.CTX, peerID, protocol.ID("/sonr/auth"))
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}
	// Establish Auth Stream
	sn.NewAuthStream(stream)

	// Create Request Message
	authPbf := &pb.AuthMessage{
		Subject:   pb.AuthMessage_REQUEST,
		PeerInfo:  sn.GetPeerInfo(),
		Metadata:  meta,
		Thumbnail: thumb,
	}

	// ** Send Invite Message **
	err = sn.AuthStream.Write(authPbf)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject: pb.AuthMessage_ACCEPT,
	}

	// Send Message
	err := sn.AuthStream.Write(authMsg)

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Decline an Invite from a Peer
func (sn *Node) Decline() bool {
	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject: pb.AuthMessage_DECLINE,
	}

	// Send Message
	err := sn.AuthStream.Write(authMsg)

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}
