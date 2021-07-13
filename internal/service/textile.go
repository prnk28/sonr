package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/phuslu/log"

	"github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/textile/v2/api/common"
	"github.com/textileio/textile/v2/cmd"
	"github.com/textileio/textile/v2/mail/local"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

var (
	isMailReady    = false
	isThreadsReady = false
	isBucketsReady = false
)

type TextileService struct {
	ctxAuth  context.Context
	ctxToken context.Context

	// Parameters
	apiKeys     *md.APIKeys
	device      *md.Device
	host        host.HostNode
	keyPair     *md.KeyPair
	options     *md.ConnectionRequest_TextileOptions
	onConnected md.OnConnected

	// Properties
	identity thread.Identity
	client   *client.Client
	mail     *local.Mail
	mailbox  *local.Mailbox
}

// @ Starts New Textile Instance
func (sc *serviceClient) StartTextile() *md.SonrError {
	// Logging
	md.LogActivate("Textile Service")

	// Initialize
	textile := &TextileService{
		keyPair:     sc.user.KeyPair(),
		options:     sc.request.GetTextileOptions(),
		apiKeys:     sc.apiKeys,
		host:        sc.host,
		onConnected: sc.handler.OnConnected,
		device:      sc.user.GetDevice(),
	}
	sc.Textile = textile

	// Check Textile Enabled
	if textile.options.GetEnabled() {
		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		textile.client, err = client.NewClient(util.TEXTILE_API_URL, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_START_CLIENT)
		}

		// Get Identity
		textile.identity = getIdentity(textile.keyPair.PrivKey())

		// Create Auth Context
		textile.ctxAuth, err = newUserAuthCtx(context.Background(), textile.apiKeys)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_USER_CTX)
		}

		// Create Token Context
		textile.ctxToken, err = textile.newTokenCtx()
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_TOKEN_CTX)
		}

		// Initialize Threads
		serr := textile.InitThreads(sc)
		if err != nil {
			return serr
		}

		// Initialize Mailbox
		serr = textile.InitMail()
		if err != nil {
			return serr
		}

		// Read Existing Mail
		serr = sc.ReadMail()
		if err != nil {
			return serr
		}
	}
	return nil
}

// @ Returns Instance Host
func (ts *TextileService) PubKey() thread.PubKey {
	return ts.identity.GetPublic()
}

// @ Initializes Threads
func (ts *TextileService) InitThreads(sc *serviceClient) *md.SonrError {
	// Verify Ready to Init
	if ts.ctxToken != nil {
		// Generate a new thread ID
		threadID := thread.NewIDV1(thread.Raw, 32)
		err := ts.client.NewDB(ts.ctxToken, threadID, db.WithNewManagedName("Sonr-Users"))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_THREADS_START_NEW)
		}

		// Get DB Info
		allInfo, err := ts.client.ListDBs(ts.ctxToken)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_THREADS_LIST_ALL)
		}

		// Log DB Info
		md.LogSuccess("Threads Activation")
		for k, v := range allInfo {
			log.Info().Msg(fmt.Sprintf("Key: %s", k.String()))
			log.Info().Msg(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), v.Addrs, v.Key.String(), v.Name))
		}
		isThreadsReady = true
	}
	return nil
}

// @ Initializes Mailbox
func (ts *TextileService) InitMail() *md.SonrError {
	// Verify Ready to Initialize
	if ts.options.GetMailbox() {
		// Log
		md.LogActivate("Textile Mailbox")

		// Setup the mail lib
		ts.mail = local.NewMail(cmd.NewClients(util.TEXTILE_API_URL, true, util.TEXTILE_MINER_IDX), local.DefaultConfConfig())

		// Create New Mailbox
		if ts.hasMailboxDirectory() {
			// Return Existing Mailbox
			mailbox, err := ts.mail.GetLocalMailbox(context.Background(), ts.device.WorkingSupportDir())
			if err != nil {
				return md.NewError(err, md.ErrorMessage_MAILBOX_START_EXISTING)
			}

			// Set Mailbox and Update Status
			ts.mailbox = mailbox
			isMailReady = true
			md.LogSuccess("Mailbox Activation")
		} else {
			// Logging
			md.LogInfo("Mailbox not found, creating new one...")

			// Create a new mailbox with config
			mailbox, err := ts.mail.NewMailbox(context.Background(), ts.defaultMailConfig())
			if err != nil {
				return md.NewError(err, md.ErrorMessage_MAILBOX_START_NEW)
			}

			// Set Mailbox and Update Status
			ts.mailbox = mailbox
			isMailReady = true
			md.LogSuccess("Mailbox Activation")
		}
	}
	return nil
}

// @ Read Mailbox Mail
func (ts *TextileService) readMail() (*md.MailEvent, *md.SonrError) {
	// List the recipient's inbox
	inbox, err := ts.mailbox.ListInboxMessages(context.Background())
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_MAILBOX_LIST_ALL)
	}

	// Initialize Entry List
	hasNewMail := false
	count := len(inbox)

	// Set Vars
	entries := make([][]byte, count)
	if count > 0 {
		hasNewMail = true
	}

	// Iterate over Entries
	for i, v := range inbox {
		// Open decrypts the message body
		body, err := v.Open(context.Background(), ts.identity)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_OPEN)
		}
		entries[i] = body
	}
	return &md.MailEvent{
		Invites:    entries,
		HasNewMail: hasNewMail,
	}, nil
}

func (ts *TextileService) sendMail(to thread.PubKey, buf []byte) ([]byte, *md.SonrError) {
	// Send Message to Mailbox
	msg, err := ts.mailbox.SendMessage(context.Background(), to, buf)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_SEND)
	}

	// Marshal Response
	buf, err = proto.Marshal(&md.InviteResponse{
		Type: md.InviteResponse_Mailbox,
		MailInfo: &md.InviteResponse_MailInfo{
			To:        msg.To.String(),
			From:      msg.From.String(),
			CreatedAt: int32(msg.CreatedAt.Unix()),
			ReadAt:    int32(msg.ReadAt.Unix()),
			Body:      msg.Body,
			Signature: msg.Signature,
		},
	})
	if err != nil {
		return nil, md.NewMarshalError(err)
	}

	// Return Message Info
	return buf, nil
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
func (ts *TextileService) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := ts.client.GetToken(ts.ctxAuth, ts.identity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(ts.ctxAuth, token), nil
}

// # Helper: Creates New Textile Mailbox Config
func (ts *TextileService) defaultMailConfig() local.Config {
	return local.Config{
		Path:      ts.device.WorkingSupportDir(),
		Identity:  ts.identity,
		APIKey:    ts.apiKeys.GetTextileKey(),
		APISecret: ts.apiKeys.GetTextileSecret(),
	}
}

// # Helper: Checks if Device Has Mailbox Directory
func (ts *TextileService) hasMailboxDirectory() bool {
	return ts.device.GetFileSystem().IsDirectory(ts.device.FileSystem.Support, util.TEXTILE_MAILBOX_DIR)
}
