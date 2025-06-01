package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"votesystem/x/voting/keeper"
	"votesystem/x/voting/types"
)

func SimulateMsgCastVote(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCastVote{
			Voter: simAccount.Address.String(),
		}

		// TODO: Handling the CastVote simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CastVote simulation not implemented"), nil, nil
	}
}
