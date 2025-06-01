package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"votesystem/x/identity/types"
)

func (k Keeper) RegisteredVoterAll(goCtx context.Context, req *types.QueryAllRegisteredVoterRequest) (*types.QueryAllRegisteredVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registeredVoters []types.RegisteredVoter
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	registeredVoterStore := prefix.NewStore(store, types.KeyPrefix(types.RegisteredVoterKey))

	pageRes, err := query.Paginate(registeredVoterStore, req.Pagination, func(key []byte, value []byte) error {
		var registeredVoter types.RegisteredVoter
		if err := k.cdc.Unmarshal(value, &registeredVoter); err != nil {
			return err
		}

		registeredVoters = append(registeredVoters, registeredVoter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegisteredVoterResponse{RegisteredVoter: registeredVoters, Pagination: pageRes}, nil
}

func (k Keeper) RegisteredVoter(goCtx context.Context, req *types.QueryGetRegisteredVoterRequest) (*types.QueryGetRegisteredVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	registeredVoter, found := k.GetRegisteredVoter(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRegisteredVoterResponse{RegisteredVoter: registeredVoter}, nil
}
