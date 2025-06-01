package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "votesystem/testutil/keeper"
	"votesystem/x/votecoin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VotecoinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
