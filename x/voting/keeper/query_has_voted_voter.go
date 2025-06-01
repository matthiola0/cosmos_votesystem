package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"votesystem/x/voting/types"
)

func (k Keeper) HasVotedVoterAll(goCtx context.Context, req *types.QueryAllHasVotedVoterRequest) (*types.QueryAllHasVotedVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hasVotedVoters []types.HasVotedVoter
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	hasVotedVoterStore := prefix.NewStore(store, types.KeyPrefix(types.HasVotedVoterKeyPrefix))

	pageRes, err := query.Paginate(hasVotedVoterStore, req.Pagination, func(key []byte, value []byte) error {
		var hasVotedVoter types.HasVotedVoter
		if err := k.cdc.Unmarshal(value, &hasVotedVoter); err != nil {
			return err
		}

		hasVotedVoters = append(hasVotedVoters, hasVotedVoter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHasVotedVoterResponse{HasVotedVoter: hasVotedVoters, Pagination: pageRes}, nil
}

func (k Keeper) HasVotedVoter(goCtx context.Context, req *types.QueryGetHasVotedVoterRequest) (*types.QueryGetHasVotedVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetHasVotedVoter(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetHasVotedVoterResponse{HasVotedVoter: val}, nil
}
