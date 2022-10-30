package api

import (
	"context"

	"github.com/sonr-io/sonr/pkg/client/config"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc"
)

// RegistryAPI is the interface for the registry module which can be accessed from the client struct
type RegistryAPI interface {
	// CreateWhoIs creates a new whois record
	CreateWhoIs(request *types.MsgCreateWhoIs) (*types.MsgCreateWhoIsResponse, error)

	// UpdateWhoIs updates a whois record
	UpdateWhoIs(request *types.MsgUpdateWhoIs) (*types.MsgUpdateWhoIsResponse, error)

	// DeleteWhoIs deletes a whois record
	DeactivateWhoIs(request *types.MsgDeactivateWhoIs) (*types.MsgDeactivateWhoIsResponse, error)

	// BuyAlias buys an alias
	BuyAlias(request *types.MsgBuyAlias) (*types.MsgBuyAliasResponse, error)

	// SellAlias sells an alias
	SellAlias(request *types.MsgSellAlias) (*types.MsgSellAliasResponse, error)

	// TransferAlias transfers an alias
	TransferAlias(request *types.MsgTransferAlias) (*types.MsgTransferAliasResponse, error)

	// QueryAllWhoIs queries all whois records
	QueryAllWhoIs(request *types.QueryAllWhoIsRequest) (*types.QueryAllWhoIsResponse, error)

	// QueryWhoIs queries a whois record
	QueryWhoIs(request *types.QueryWhoIsRequest) (*types.QueryWhoIsResponse, error)

	// QueryWhoIsByAlias queries a whois record by alias
	QueryWhoIsByAlias(request *types.QueryWhoIsAliasRequest) (*types.QueryWhoIsAliasResponse, error)

	// QueryWhoIsByOwner queries a whois record by owner
	QueryWhoIsByOwner(request *types.QueryWhoIsControllerRequest) (*types.QueryWhoIsControllerResponse, error)
}

// registryAPIImpl is the implementation of the registry API
type registryAPIImpl struct {
	config *config.Config
}

// NewRegistryAPI creates a new registry API
func NewRegistryAPI(config *config.Config) RegistryAPI {
	return &registryAPIImpl{
		config: config,
	}
}

// CreateWhoIs creates a new whois record
func (r *registryAPIImpl) CreateWhoIs(request *types.MsgCreateWhoIs) (*types.MsgCreateWhoIsResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgCreateWhoIs", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgCreateWhoIsResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateWhoIs updates a whois record
func (r *registryAPIImpl) UpdateWhoIs(request *types.MsgUpdateWhoIs) (*types.MsgUpdateWhoIsResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgUpdateWhoIs", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgUpdateWhoIsResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeactivateWhoIs deletes a whois record
func (r *registryAPIImpl) DeactivateWhoIs(request *types.MsgDeactivateWhoIs) (*types.MsgDeactivateWhoIsResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgDeactivateWhoIs", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgDeactivateWhoIsResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BuyAlias buys an alias
func (r *registryAPIImpl) BuyAlias(request *types.MsgBuyAlias) (*types.MsgBuyAliasResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgBuyAlias", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgBuyAliasResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// SellAlias sells an alias
func (r *registryAPIImpl) SellAlias(request *types.MsgSellAlias) (*types.MsgSellAliasResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgSellAlias", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgSellAliasResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// TransferAlias transfers an alias
func (r *registryAPIImpl) TransferAlias(request *types.MsgTransferAlias) (*types.MsgTransferAliasResponse, error) {
	txRaw, err := r.config.SignTxWithWallet("/sonrio.sonr.registry.MsgTransferAlias", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(r.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgTransferAliasResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryAllWhoIs queries all whois records
func (r *registryAPIImpl) QueryAllWhoIs(request *types.QueryAllWhoIsRequest) (*types.QueryAllWhoIsResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		r.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhoIsAll(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhoIs queries a whois record
func (r *registryAPIImpl) QueryWhoIs(request *types.QueryWhoIsRequest) (*types.QueryWhoIsResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		r.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhoIs(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhoIsByAlias queries a whois record by alias
func (r *registryAPIImpl) QueryWhoIsByAlias(request *types.QueryWhoIsAliasRequest) (*types.QueryWhoIsAliasResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		r.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhoIsAlias(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhoIsByOwner queries a whois record by owner
func (r *registryAPIImpl) QueryWhoIsByOwner(request *types.QueryWhoIsControllerRequest) (*types.QueryWhoIsControllerResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		r.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhoIsController(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}
