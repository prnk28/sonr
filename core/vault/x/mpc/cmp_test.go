package mpc_test

import (
	"context"
	"testing"

	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/core/vault/x/mpc"
	"github.com/sonr-io/sonr/internal/node"
)

const (
	path     = "/Users/pradn/.sonr/wallet/a_sonr_id.json"
	id       = "a"
	password = "goodpassword"
)

var (
	ids = party.IDSlice{"a", "b", "c", "d", "e", "f"}
	msg = []byte("hello world")
	sig []byte
)

func TestCreateSaveLoadWallet(t *testing.T) {
	// Create 2 nodes
	n1, err := node.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, err = node.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Create a Channel for the nodes to communicate on
	c1, err := n1.Join("test")
	if err != nil {
		t.Fatal(err)
	}
	pids := c1.ListPeers()
	n, err := node.New(context.Background())
	w, err := mpc.NewWallet(n, pids)
	if err != nil {
		t.Fatal(err)
	}
	_, err = w.Save(password)
	if err != nil {
		t.Fatal(err)
	}
	_, err = mpc.NewWalletFromDisk(path, password)
	if err != nil {
		t.Fatal(err)
	}
}
