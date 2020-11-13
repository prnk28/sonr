package sonr

import (
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
)

// Node contains all values for user
type Node struct {
	PeerID  string
	Host    host.Host
	Lobby   lobby.Lobby
	Profile user.Profile
	Contact user.Contact
}

// Send publishes a message to the SonrNode lobby
func (sn *Node) Send(data string) bool {
	// Log Send
	fmt.Println("Sonr P2P General-Send: ", data)

	// Create Message Type
	cm := new(lobby.Message)
	err := json.Unmarshal([]byte(data), cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Publish to Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() string {
	// Initialize Map
	m := make(map[string]string)
	m["profile"] = sn.Profile.String()
	m["contact"] = sn.Contact.String()
	m["id"] = sn.PeerID

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		println(err)
	}

	// Return String
	return string(msgBytes)
}

// SetUser from connection request
func (sn *Node) SetUser(cm lobby.ConnectRequest) error {
	// Set Profile
	profile := user.NewProfile(sn.Host.ID().String(), cm.OLC, cm.Device)
	sn.Profile = profile

	// Set Contact
	contact := user.NewContact(cm.Contact)
	sn.Contact = contact

	return nil
}

// Update occurs when status or direction changes
func (sn *Node) Update(data string) bool {
	// Get Update from Json
	notif := new(lobby.Notification)
	notif.ID = sn.PeerID
	err := json.Unmarshal([]byte(data), notif)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Repackage with graph ID
	renotif, err := json.Marshal(notif)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Update User Values
	sn.Profile.Update(notif.Direction)

	// Create Message
	cm := new(lobby.Message)
	cm.Event = "Update"
	cm.SenderID = sn.PeerID
	cm.Data = string(renotif)

	// Inform Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}
