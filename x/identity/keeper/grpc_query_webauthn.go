package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WebauthnRegistrationBegin(goCtx context.Context, req *types.QueryWebauthnRegisterBeginRequest) (*types.QueryWebauthnRegisterBeginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	

	return &types.QueryWebauthnRegisterBeginResponse{}, nil
}

func (k Keeper) WebauthnRegistrationFinish(goCtx context.Context, req *types.QueryWebauthnRegisterFinishRequest) (*types.QueryWebauthnRegisterFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryWebauthnRegisterFinishResponse{}, nil
}

func (k Keeper) WebauthnLoginBegin(goCtx context.Context, req *types.QueryWebauthnLoginBeginRequest) (*types.QueryWebauthnLoginBeginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryWebauthnLoginBeginResponse{}, nil
}

func (k Keeper) WebauthnLoginFinish(goCtx context.Context, req *types.QueryWebauthnLoginFinishRequest) (*types.QueryWebauthnLoginFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryWebauthnLoginFinishResponse{}, nil
}
