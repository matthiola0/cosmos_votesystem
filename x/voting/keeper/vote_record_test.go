package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/voting/keeper"
	"votesystem/x/voting/types"
)

func createNVoteRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VoteRecord {
	items := make([]types.VoteRecord, n)
	for i := range items {
		items[i].Id = keeper.AppendVoteRecord(ctx, items[i])
	}
	return items
}

func TestVoteRecordGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteRecord(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVoteRecord(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVoteRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVoteRecord(ctx, item.Id)
		_, found := keeper.GetVoteRecord(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVoteRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVoteRecord(ctx)),
	)
}

func TestVoteRecordCount(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteRecord(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVoteRecordCount(ctx))
}
