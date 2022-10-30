package config

import (
	"github.com/sonr-io/sonr/pkg/crypto/wallet"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

func New(options ...Option) *Config {
	c := &Config{
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

type Option func(*Config)

func WithClientMode(mode mt.ClientMode) Option {
	return func(c *Config) {
		c.ClientMode = mode
		if c.Wallet == nil {
			c.HasWallet = false
		}
	}
}

func WithWallet(w wallet.EDCSAWallet) Option {
	return func(c *Config) {
		c.Wallet = w
		c.HasWallet = true
	}
}
