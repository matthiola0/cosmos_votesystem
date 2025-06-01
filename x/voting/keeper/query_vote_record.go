package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"votesystem/x/voting/types"
)

func (k Keeper) VoteRecordAll(goCtx context.Context, req *types.QueryAllVoteRecordRequest) (*types.QueryAllVoteRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var voteRecords []types.VoteRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	voteRecordStore := prefix.NewStore(store, types.KeyPrefix(types.VoteRecordKey))

	pageRes, err := query.Paginate(voteRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var voteRecord types.VoteRecord
		if err := k.cdc.Unmarshal(value, &voteRecord); err != nil {
			return err
		}

		voteRecords = append(voteRecords, voteRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteRecordResponse{VoteRecord: voteRecords, Pagination: pageRes}, nil
}

func (k Keeper) VoteRecord(goCtx context.Context, req *types.QueryGetVoteRecordRequest) (*types.QueryGetVoteRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	voteRecord, found := k.GetVoteRecord(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVoteRecordResponse{VoteRecord: voteRecord}, nil
}
