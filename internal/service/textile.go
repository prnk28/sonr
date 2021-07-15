package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

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
	options     *md.ConnectionRequest_ServiceOptions
	onConnected md.OnConnected
	handler     ServiceHandler

	// Properties
	client  *client.Client
	mail    *local.Mail
	mailbox *local.Mailbox
}

// @ Starts New Textile Instance
func (sc *serviceClient) StartTextile() *md.SonrError {
	// Logging
	md.LogActivate("Textile Service")

	// Initialize
	textile := &TextileService{
		options:     sc.request.GetServiceOptions(),
		apiKeys:     sc.apiKeys,
		host:        sc.host,
		onConnected: sc.handler.OnConnected,
		device:      sc.user.GetDevice(),
		handler:     sc.handler,
	}
	sc.Textile = textile

	// Check Textile Enabled
	if textile.options.GetTextile() {
		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		textile.client, err = client.NewClient(util.TEXTILE_API_URL, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_START_CLIENT)
		}

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
	}
	return nil
}

// @ Returns Instance Host
func (ts *TextileService) PubKey() thread.PubKey {
	return ts.device.ThreadIdentity().GetPublic()
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

		// Log DB Info
		md.LogSuccess("Threads Activation")
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

			// Handle Mailbox Events
			ts.handleMailboxEvents()
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

			// Handle Mailbox Events
			ts.handleMailboxEvents()
		}
	}
	return nil
}

// @ Handle Mailbox Events
func (ts *TextileService) handleMailboxEvents() {
	// Handle mailbox events as they arrive
	events := make(chan local.MailboxEvent)
	defer close(events)
	go func() {
		for e := range events {
			switch e.Type {
			case local.NewMessage:
				ts.onNewMessage(e)
				continue
			}
		}
	}()

	// Start watching (the third param indicates we want to keep watching when offline)
	state, err := ts.mailbox.WatchInbox(context.Background(), events, true)
	if err != nil {
		md.NewError(err, md.ErrorMessage_MAILBOX_EVENT_STATE)
		return
	}

	// Handle Mailbox State
	md.LogSuccess("Mailbox State Handling")
	for s := range state {
		// handle connectivity state
		switch s.State {
		case cmd.Online:
			md.LogInfo(fmt.Sprintf("Mailbox is Online: %s", s.Err))
		case cmd.Offline:
			md.LogInfo(fmt.Sprintf("Mailbox is Offline: %s", s.Err))
		}
	}
}

// @ Handle New Mailbox Message
func (ts *TextileService) onNewMessage(e local.MailboxEvent) {
	// Logging Received
	inbox, err := ts.mailbox.ListInboxMessages(context.Background())
	if err != nil {
		md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_OPEN)
		return
	}

	// Logging and Open Body
	md.LogInfo(fmt.Sprintf("Received new message: %s", inbox[0].From))
	body, err := inbox[0].Open(context.Background(), ts.mailbox.Identity())
	if err != nil {
		md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_OPEN)
		return
	}

	// Log Valid Lobby Length
	md.LogInfo(fmt.Sprintf("Valid Body Length: %d", len(body)))

	// Create Mail Event
	mail := &md.MailEvent{
		To:        inbox[0].To.String(),
		From:      inbox[0].From.String(),
		CreatedAt: int32(inbox[0].CreatedAt.Unix()),
		ReadAt:    int32(inbox[0].ReadAt.Unix()),
		Body:      body,
		Signature: inbox[0].Signature,
	}

	// Callback Mail Event
	ts.handler.OnMail(mail)

	// Mark Item as Read
	err = ts.mailbox.ReadInboxMessage(context.Background(), inbox[0].ID)
	if err != nil {
		md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_READ)
	}
}

// @ Send Mail to Recipient
func (ts *TextileService) sendMail(to thread.PubKey, buf []byte) *md.SonrError {
	// Send Message to Mailbox
	_, err := ts.mailbox.SendMessage(context.Background(), to, buf)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_SEND)
	}

	// Return Message Info
	return nil
}

// # Helper: Creates New Textile Mailbox Config
func (ts *TextileService) defaultMailConfig() local.Config {
	return local.Config{
		Path:      ts.device.WorkingSupportDir(),
		Identity:  ts.device.ThreadIdentity(),
		APIKey:    ts.apiKeys.GetTextileKey(),
		APISecret: ts.apiKeys.GetTextileSecret(),
	}
}

// # Helper: Checks if Device Has Mailbox Directory
func (ts *TextileService) hasMailboxDirectory() bool {
	return ts.device.GetFileSystem().IsDirectory(ts.device.FileSystem.Support, util.TEXTILE_MAILBOX_DIR)
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
	token, err := ts.client.GetToken(ts.ctxAuth, ts.device.ThreadIdentity())
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(ts.ctxAuth, token), nil
}