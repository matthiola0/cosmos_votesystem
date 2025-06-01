package voting_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "votesystem/testutil/keeper"
	"votesystem/testutil/nullify"
	"votesystem/x/voting"
	"votesystem/x/voting/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VoteRecordList: []types.VoteRecord{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VoteRecordCount: 2,
		HasVotedVoterList: []types.HasVotedVoter{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotingKeeper(t)
	voting.InitGenesis(ctx, *k, genesisState)
	got := voting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VoteRecordList, got.VoteRecordList)
	require.Equal(t, genesisState.VoteRecordCount, got.VoteRecordCount)
	require.ElementsMatch(t, genesisState.HasVotedVoterList, got.HasVotedVoterList)
	// this line is used by starport scaffolding # genesis/test/assert
}
