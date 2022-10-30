package client

import (
	"github.com/sonr-io/sonr/pkg/client/api"
	"github.com/sonr-io/sonr/pkg/client/config"
	"github.com/sonr-io/sonr/pkg/crypto/wallet"
)

type Client struct {
	*config.Config
	api.BucketAPI
	api.RegistryAPI
	api.SchemaAPI
	api.PaymentsAPI
}

func NewClient(options ...config.Option) *Client {
	cnfg := config.New(options...)
	return &Client{
		Config:      cnfg,
		BucketAPI:   api.NewBucketAPI(cnfg),
		RegistryAPI: api.NewRegistryAPI(cnfg),
		SchemaAPI:   api.NewSchemaAPI(cnfg),
		PaymentsAPI: api.NewPaymentsAPI(cnfg),
	}
}

func (c *Client) GetFaucetAddress() string {
	return c.Config.GetFaucetAddress()
}

func (c *Client) GetRPCAddress() string {
	return c.Config.GetRPCAddress()
}

func (c *Client) GetAPIAddress() string {
	return c.Config.GetAPIAddress()
}

func (c *Client) GetIPFSAddress() string {
	return c.Config.GetIPFSAddress()
}

func (c *Client) GetIPFSApiAddress() string {
	return c.Config.GetIPFSApiAddress()
}

func (c *Client) SetWallet(w wallet.EDCSAWallet) {
	c.Config.Wallet = w
	c.Config.HasWallet = true
}
