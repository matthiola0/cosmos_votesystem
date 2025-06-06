package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"votesystem/x/voting/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated voteRecord",
			genState: &types.GenesisState{
				VoteRecordList: []types.VoteRecord{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid voteRecord count",
			genState: &types.GenesisState{
				VoteRecordList: []types.VoteRecord{
					{
						Id: 1,
					},
				},
				VoteRecordCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated hasVotedVoter",
			genState: &types.GenesisState{
				HasVotedVoterList: []types.HasVotedVoter{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
