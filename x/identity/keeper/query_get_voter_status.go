package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"votesystem/x/identity/types"
)

func (k Keeper) GetVoterStatus(goCtx context.Context, req *types.QueryGetVoterStatusRequest) (*types.QueryGetVoterStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(req.VoterAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid voter address format")
	}

	allRegisteredVoters := k.GetAllRegisteredVoter(ctx)
	for _, registeredVoter := range allRegisteredVoters {
		if registeredVoter.VoterAddress == req.VoterAddress {
			// 找到選民，返回其資訊
			return &types.QueryGetVoterStatusResponse{
				IsRegistered:      true,
				IdentityProofHash: registeredVoter.IdentityProofHash,
			}, nil
		}
	}

	// 未找到選民
	return &types.QueryGetVoterStatusResponse{
		IsRegistered:      false,
		IdentityProofHash: "",
	}, nil
}