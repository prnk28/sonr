package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-io/sonr/pkg/client/config"
	"google.golang.org/grpc"
)

// PaymentsAPI is the interface for managing the Bank module on the Blockchain
type PaymentsAPI interface {
	// SendTokens sends tokens to an address
	SendTokens(request *types.MsgSend) (*types.MsgSendResponse, error)

	// SendTokensToMany sends tokens to multiple addresses
	SendTokensToMany(request *types.MsgMultiSend) (*types.MsgMultiSendResponse, error)

	// GetBalance gets the balance of an address
	GetBalance(address string) (sdk.Coins, error)

	// RequestFaucet funds an address with tokens
	RequestFaucet(address string) error
}

// paymentsAPIImpl is the implementation of the payments API
type paymentsAPIImpl struct {
	config *config.Config
}

// NewPaymentsAPI creates a new payments API
func NewPaymentsAPI(config *config.Config) PaymentsAPI {
	return &paymentsAPIImpl{
		config: config,
	}
}

// SendTokens sends tokens to an address
func (p *paymentsAPIImpl) SendTokens(request *types.MsgSend) (*types.MsgSendResponse, error) {
	txRaw, err := p.config.SignTxWithWallet("/cosmos.bank.v1beta1.MsgSend", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(p.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgSendResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// SendTokensToMany sends tokens to multiple addresses
func (p *paymentsAPIImpl) SendTokensToMany(request *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	txRaw, err := p.config.SignTxWithWallet("/cosmos.bank.v1beta1.MsgMultiSend", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(p.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgMultiSendResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetBalance gets the balance of an address
func (p *paymentsAPIImpl) GetBalance(address string) (sdk.Coins, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		p.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	resp, err := types.NewQueryClient(grpcConn).AllBalances(context.Background(), &types.QueryAllBalancesRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetBalances(), nil
}

// RequestFaucet requests a faucet from the Sonr network
func (p *paymentsAPIImpl) RequestFaucet(address string) error {
	values := faucetRequest{
		Address: address,
		Coins:   []string{"12000snr"},
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(p.config.GetFaucetAddress(), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

type faucetRequest struct {
	Address string   `json:"address"`
	Coins   []string `json:"coins"`
}
