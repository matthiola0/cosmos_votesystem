package keeper

import (
	"context"

	"votesystem/x/voting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetElectionResult(goCtx context.Context, req *types.QueryGetElectionResultRequest) (*types.QueryGetElectionResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.ElectionID == "" {
		return nil, status.Error(codes.InvalidArgument, "election ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	allVoteRecords := k.GetAllVoteRecord(ctx) // 確保 GetAllVoteRecord 存在並正確工作
	electionResults := make(map[string]uint64)

	for _, record := range allVoteRecords {
		if record.ElectionID == req.ElectionID {
			electionResults[record.CandidateID]++
		}
	}

	return &types.QueryGetElectionResultResponse{
		ElectionId: req.ElectionID,
		Results:    electionResults,
	}, nil
}
