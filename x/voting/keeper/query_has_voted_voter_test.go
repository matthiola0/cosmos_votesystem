package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/voting/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHasVotedVoterQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHasVotedVoter(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetHasVotedVoterRequest
		response *types.QueryGetHasVotedVoterResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetHasVotedVoterRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetHasVotedVoterResponse{HasVotedVoter: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetHasVotedVoterRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetHasVotedVoterResponse{HasVotedVoter: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetHasVotedVoterRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.HasVotedVoter(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestHasVotedVoterQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHasVotedVoter(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHasVotedVoterRequest {
		return &types.QueryAllHasVotedVoterRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HasVotedVoterAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HasVotedVoter), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HasVotedVoter),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HasVotedVoterAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HasVotedVoter), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HasVotedVoter),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HasVotedVoterAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HasVotedVoter),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HasVotedVoterAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
