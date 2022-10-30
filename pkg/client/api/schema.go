package api

import (
	"context"

	"github.com/sonr-io/sonr/pkg/client/config"
	"github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
)

// SchemaAPI is the interface for the schema module which can be accessed from the client struct
type SchemaAPI interface {
	// CreateSchema creates a schema
	CreateSchema(request *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error)

	// DeprecateSchema deprecates a schema
	DeprecateSchema(request *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error)

	// QuerySchema queries a schema
	QuerySchema(request *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error)

	// QueryWhatIs queries a whatIs document
	QueryWhatIs(request *types.QueryWhatIsRequest) (*types.QueryWhatIsResponse, error)

	// QueryWhatIsByCreator queries whatIs documents by creator
	QueryWhatIsByCreator(request *types.QueryWhatIsCreatorRequest) (*types.QueryWhatIsCreatorResponse, error)

	// QueryWhatIsByDid queries whatIs documents by did
	QueryWhatIsByDid(request *types.QueryWhatIsByDidRequest) (*types.QueryWhatIsByDidResponse, error)

	// QueryWhatIsAll queries all whatIs documents
	QueryWhatIsAll(request *types.QueryAllWhatIsRequest) (*types.QueryAllWhatIsResponse, error)
}

// schemaAPIImpl is the implementation of the schema API
type schemaAPIImpl struct {
	config *config.Config
}

// NewSchemaAPI creates a new schema API
func NewSchemaAPI(config *config.Config) SchemaAPI {
	return &schemaAPIImpl{
		config: config,
	}
}

// CreateSchema creates a schema
func (s *schemaAPIImpl) CreateSchema(request *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	txRaw, err := s.config.SignTxWithWallet("/sonrio.sonr.schema.MsgCreateSchema", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(s.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgCreateSchemaResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeprecateSchema deprecates a schema
func (s *schemaAPIImpl) DeprecateSchema(request *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	txRaw, err := s.config.SignTxWithWallet("/sonrio.sonr.schema.MsgDeprecateSchema", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(s.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgDeprecateSchemaResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QuerySchema queries a schema
func (s *schemaAPIImpl) QuerySchema(request *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		s.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).Schema(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhatIs queries a whatIs document
func (s *schemaAPIImpl) QueryWhatIs(request *types.QueryWhatIsRequest) (*types.QueryWhatIsResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		s.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhatIs(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhatIsByCreator queries whatIs documents by creator
func (s *schemaAPIImpl) QueryWhatIsByCreator(request *types.QueryWhatIsCreatorRequest) (*types.QueryWhatIsCreatorResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		s.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhatIsByCreator(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhatIsByDid queries whatIs documents by did
func (s *schemaAPIImpl) QueryWhatIsByDid(request *types.QueryWhatIsByDidRequest) (*types.QueryWhatIsByDidResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		s.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhatIsByDid(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryWhatIsAll queries all whatIs documents
func (s *schemaAPIImpl) QueryWhatIsAll(request *types.QueryAllWhatIsRequest) (*types.QueryAllWhatIsResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		s.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).WhatIsAll(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}
