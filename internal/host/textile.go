package host

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

// @ Initializes New Textile Instance
func (hn *hostNode) StartTextile(d *md.Device) *md.SonrError {
	// Check Textile Enabled
	if hn.tileOptions.GetIsEnabled() {
		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
		hn.tileClient, err = client.NewClient(util.TEXTILE_API_URL, opts...)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Get Identity
		hn.tileIdentity = getIdentity(hn.keyPair.PrivKey())

		// Create Auth Context
		hn.ctxTileAuth, err = newUserAuthCtx(context.Background(), hn.apiKeys)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Create Token Context
		hn.ctxTileToken, err = hn.newTokenCtx()
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Check Thread Enabled
		if hn.tileOptions.GetWithThreads() {
			// Generate a new thread ID
			threadID := thread.NewIDV1(thread.Raw, 32)

			// Create your new thread
			err = hn.tileClient.NewDB(hn.ctxTileToken, threadID)
			if err != nil {
				return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Get DB Info
			info, err := hn.tileClient.GetDBInfo(hn.ctxTileToken, threadID)
			if err != nil {
				return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Log DB Info
			log.Println("> Success!")
			log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), info.Addrs, info.Key.String(), info.Name))
		}
	}
	return nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (hn *hostNode) ReadMail() ([]*md.MailEntry, *md.SonrError) {
	// Check Mail Enabled
	if hn.tileOptions.GetWithMailbox() {
		// List the recipient's inbox
		inbox, err := hn.tileMailbox.ListInboxMessages(context.Background())

		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Initialize Entry List
		entries := make([]*md.MailEntry, len(inbox))

		// Iterate over Entries
		for i, v := range inbox {
			// Open decrypts the message body
			body, err := v.Open(context.Background(), hn.tileIdentity)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Unmarshal Body to entry
			entry := &md.MailEntry{}
			err = proto.Unmarshal(body, entry)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}
			entries[i] = entry
		}
		return entries, nil
	}
	return nil, nil
}

// @ Method Sends Mail Entry to Peer
func (hn *hostNode) SendMail(e *md.MailEntry) *md.SonrError {
	// Check Mail Enabled
	if hn.tileOptions.GetWithMailbox() {
		// Send Message to Mailbox
		_, err := hn.tileMailbox.SendMessage(context.Background(), e.ToPubKey(), e.Buffer())

		// Check Error
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
		return nil
	}
	return nil
}

// # Helper: Gets Thread Identity from Private Key
func getIdentity(privateKey crypto.PrivKey) thread.Identity {
	myIdentity := thread.NewLibp2pIdentity(privateKey)
	return myIdentity
}

// # Helper: Creates User Auth Context from API Keys
func newUserAuthCtx(ctx context.Context, keys *md.APIKeys) (context.Context, error) {
	// Add our user group key to the context
	ctx = common.NewAPIKeyContext(ctx, keys.TextileKey)

	// Add a signature using our user group secret
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Hour), keys.TextileSecret)
}

// # Helper: Creates Auth Token Context from AuthContext, Client, Identity
func (hn *hostNode) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := hn.tileClient.GetToken(hn.ctxTileAuth, hn.tileIdentity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(hn.ctxTileAuth, token), nil
}