package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"votesystem/x/identity/keeper"
	"votesystem/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the registeredVoter
	for _, elem := range genState.RegisteredVoterList {
		k.SetRegisteredVoter(ctx, elem)
	}

	// Set registeredVoter count
	k.SetRegisteredVoterCount(ctx, genState.RegisteredVoterCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RegisteredVoterList = k.GetAllRegisteredVoter(ctx)
	genesis.RegisteredVoterCount = k.GetRegisteredVoterCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
