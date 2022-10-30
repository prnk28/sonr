package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
	"github.com/sonr-io/sonr/pkg/crypto/wallet"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

const (
	// -- Local Blockchain --
	BLOCKCHAIN_REST_LOCAL   = "http://0.0.0.0:26657"
	BLOCKCHAIN_FAUCET_LOCAL = "http://0.0.0.0:4500"
	BLOCKCHAIN_RPC_LOCAL    = "127.0.0.1:9090"

	// -- Dev Blockchain --
	BLOCKCHAIN_REST_DEV   = "http://143.198.29.209:1317"
	BLOCKCHAIN_FAUCET_DEV = "http://143.198.29.209:8000"
	BLOCKCHAIN_RPC_DEV    = "143.198.29.209:9090"

	// -- Beta Blockchain --
	BLOCKCHAIN_REST_BETA   = "http://137.184.190.146:1317"
	BLOCKCHAIN_FAUCET_BETA = "http://137.184.190.146:8000"
	BLOCKCHAIN_RPC_BETA    = "137.184.190.146:9090"

	// -- Services --
	IPFS_ADDRESS      = "https://ipfs.sonr.ws"
	IPFS_API_ADDRESS  = "https://api.ipfs.sonr.ws"
	VAULT_API_ADDRESS = "http://164.92.99.233"
)

type Config struct {
	ClientMode mt.ClientMode
	Wallet     wallet.EDCSAWallet
	HasWallet  bool
}

func (c *Config) GetFaucetAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.ClientMode {
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_FAUCET_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_FAUCET_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_FAUCET_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_FAUCET")
}

func (c *Config) GetRPCAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.ClientMode {
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_RPC_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_RPC_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_RPC_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_RPC")
}

func (c *Config) GetAPIAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		// .env does not exist, use preset client mode
		switch c.ClientMode {
		case mt.ClientMode_ENDPOINT_BETA:
			return BLOCKCHAIN_REST_BETA
		case mt.ClientMode_ENDPOINT_DEV:
			return BLOCKCHAIN_REST_DEV
		case mt.ClientMode_ENDPOINT_LOCAL:
			return BLOCKCHAIN_REST_LOCAL
		}
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("BLOCKCHAIN_REST")
}

func (c *Config) GetIPFSAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		return IPFS_ADDRESS
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("IPFS_ADDRESS")
}

func (c *Config) GetIPFSApiAddress() string {
	env_path := filepath.Join(projectpath.Root, ".env")

	// by default use .env if it exists
	_, err := os.Stat(env_path)
	if errors.Is(err, os.ErrNotExist) {
		return IPFS_API_ADDRESS
	}

	err = godotenv.Load(env_path)
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("IPFS_API_ADDRESS")
}

func (c *Config) SignTxWithWallet(typeUrl string, msgs ...sdk.Msg) ([]byte, error) {
	if !c.HasWallet {
		return nil, errors.New("no wallet was passed in client configuration")
	}
	return wallet.SignTxWithWallet(c.Wallet, typeUrl, msgs...)
}
