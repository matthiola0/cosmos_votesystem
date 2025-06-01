package votecoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/votecoin"
	"votesystem/x/votecoin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotecoinKeeper(t)
	votecoin.InitGenesis(ctx, *k, genesisState)
	got := votecoin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
