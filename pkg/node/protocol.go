package node

import (
	"log"

	brpt "berty.tech/berty/v2/go/pkg/bertyprotocol"
	brmd "berty.tech/berty/v2/go/pkg/protocoltypes"
)

// ^ Starts Protocol ^
func (n *Node) StartProtocol() error {
	// Start New Berty Client
	client, err := brpt.New(n.ctx, brpt.Opts{
		Host:   n.Host.Host,
		PubSub: n.Host.Pubsub,
	})
	if err != nil {
		return err
	}
	n.client = client

	// brpt.NewGroupMultiMember()
	resp, err := client.MultiMemberGroupCreate(n.ctx, &brmd.MultiMemberGroupCreate_Request{})
	if err != nil {
		return err
	}
	log.Println(string(resp.GroupPK))

	prresp, err := client.PeerList(n.ctx, &brmd.PeerList_Request{})
	if err != nil {
		return err
	}

	for _, v := range prresp.Peers {
		log.Println(v.String())
	}

	if err != nil {
		return err
	}
	return nil
}
