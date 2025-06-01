package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/voting/keeper"
	"votesystem/x/voting/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNHasVotedVoter(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HasVotedVoter {
	items := make([]types.HasVotedVoter, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetHasVotedVoter(ctx, items[i])
	}
	return items
}

func TestHasVotedVoterGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNHasVotedVoter(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetHasVotedVoter(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestHasVotedVoterRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNHasVotedVoter(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHasVotedVoter(ctx,
			item.Index,
		)
		_, found := keeper.GetHasVotedVoter(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestHasVotedVoterGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNHasVotedVoter(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHasVotedVoter(ctx)),
	)
}
