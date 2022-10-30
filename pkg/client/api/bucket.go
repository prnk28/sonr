package api

import (
	"context"

	"github.com/sonr-io/sonr/pkg/client/config"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc"
)

// BucketAPI is the interface for the bucket module which can be accessed from the client struct
type BucketAPI interface {
	// DefineBucket configure a bucket
	DefineBucket(request *types.MsgDefineBucket) (*types.MsgDefineBucketResponse, error)

	// AllocateBucket allocates a bucket
	AllocateBucket(request *types.MsgAllocateBucket) (*types.MsgAllocateBucketResponse, error)

	// DeactivateBucket deactivates a bucket
	DeactivateBucket(request *types.MsgDeactivateBucket) (*types.MsgDeactivateBucketResponse, error)

	// BurnBucket burns a bucket
	BurnBucket(request *types.MsgBurnBucket) (*types.MsgBurnBucketResponse, error)

	// QueryBucket queries a bucket
	QueryBucket(request *types.QueryGetBucketRequest) (*types.QueryGetBucketResponse, error)

	// QueryAllBuckets queries all buckets
	QueryAllBuckets(request *types.QueryAllBucketsRequest) (*types.QueryAllBucketsResponse, error)

	// QueryBucketByCreator queries buckets by creator
	QueryBucketByCreator(request *types.QueryGetBucketByCreatorRequest) (*types.QueryGetBucketByCreatorResponse, error)

	// QueryBucketByName queries buckets by name
	QueryBucketByName(request *types.QueryGetBucketByNameRequest) (*types.QueryGetBucketByNameResponse, error)
}

// bucketAPIImpl is the implementation of the bucket API
type bucketAPIImpl struct {
	config *config.Config
}

// NewBucketAPI creates a new bucket API
func NewBucketAPI(config *config.Config) BucketAPI {
	return &bucketAPIImpl{
		config: config,
	}
}

// DefineBucket configure a bucket
func (b *bucketAPIImpl) DefineBucket(request *types.MsgDefineBucket) (*types.MsgDefineBucketResponse, error) {
	txRaw, err := b.config.SignTxWithWallet("/sonrio.sonr.bucket.MsgDefineBucket", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(b.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgDefineBucketResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// AllocateBucket allocates a bucket
func (b *bucketAPIImpl) AllocateBucket(request *types.MsgAllocateBucket) (*types.MsgAllocateBucketResponse, error) {
	txRaw, err := b.config.SignTxWithWallet("/sonrio.sonr.bucket.MsgAllocateBucket", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(b.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgAllocateBucketResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeactivateBucket deactivates a bucket
func (b *bucketAPIImpl) DeactivateBucket(request *types.MsgDeactivateBucket) (*types.MsgDeactivateBucketResponse, error) {
	txRaw, err := b.config.SignTxWithWallet("/sonrio.sonr.bucket.MsgDeactivateBucket", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(b.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgDeactivateBucketResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BurnBucket burns a bucket
func (b *bucketAPIImpl) BurnBucket(request *types.MsgBurnBucket) (*types.MsgBurnBucketResponse, error) {
	txRaw, err := b.config.SignTxWithWallet("/sonrio.sonr.bucket.MsgBurnBucket", request)
	if err != nil {
		return nil, err
	}

	respRaw, err := BroadcastTx(b.config.GetRPCAddress(), txRaw)
	if err != nil {
		return nil, err
	}

	resp := &types.MsgBurnBucketResponse{}
	if err := DecodeTxResponseData(respRaw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryBucket queries a bucket
func (b *bucketAPIImpl) QueryBucket(request *types.QueryGetBucketRequest) (*types.QueryGetBucketResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		b.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).Bucket(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryAllBuckets queries all buckets
func (b *bucketAPIImpl) QueryAllBuckets(request *types.QueryAllBucketsRequest) (*types.QueryAllBucketsResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		b.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).BucketsAll(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryBucketByCreator queries buckets by creator
func (b *bucketAPIImpl) QueryBucketByCreator(request *types.QueryGetBucketByCreatorRequest) (*types.QueryGetBucketByCreatorResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		b.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).BucketByCreator(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}

// QueryBucketByName queries buckets by name
func (b *bucketAPIImpl) QueryBucketByName(request *types.QueryGetBucketByNameRequest) (*types.QueryGetBucketByNameResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		b.config.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),      // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// We then call the QueryWhoIs method on this client.
	res, err := types.NewQueryClient(grpcConn).BucketByName(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}
