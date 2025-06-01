package voting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"votesystem/x/voting/keeper"
	"votesystem/x/voting/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the voteRecord
	for _, elem := range genState.VoteRecordList {
		k.SetVoteRecord(ctx, elem)
	}

	// Set voteRecord count
	k.SetVoteRecordCount(ctx, genState.VoteRecordCount)
	// Set all the hasVotedVoter
	for _, elem := range genState.HasVotedVoterList {
		k.SetHasVotedVoter(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VoteRecordList = k.GetAllVoteRecord(ctx)
	genesis.VoteRecordCount = k.GetVoteRecordCount(ctx)
	genesis.HasVotedVoterList = k.GetAllHasVotedVoter(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
