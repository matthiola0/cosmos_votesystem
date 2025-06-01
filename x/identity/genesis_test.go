package identity_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/identity"
	"votesystem/x/identity/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RegisteredVoterList: []types.RegisteredVoter{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RegisteredVoterCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RegisteredVoterList, got.RegisteredVoterList)
	require.Equal(t, genesisState.RegisteredVoterCount, got.RegisteredVoterCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
