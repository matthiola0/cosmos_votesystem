package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/identity/keeper"
	"votesystem/x/identity/types"
)

func createNRegisteredVoter(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegisteredVoter {
	items := make([]types.RegisteredVoter, n)
	for i := range items {
		items[i].Id = keeper.AppendRegisteredVoter(ctx, items[i])
	}
	return items
}

func TestRegisteredVoterGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegisteredVoter(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRegisteredVoter(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRegisteredVoterRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegisteredVoter(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegisteredVoter(ctx, item.Id)
		_, found := keeper.GetRegisteredVoter(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRegisteredVoterGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegisteredVoter(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegisteredVoter(ctx)),
	)
}

func TestRegisteredVoterCount(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegisteredVoter(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRegisteredVoterCount(ctx))
}
