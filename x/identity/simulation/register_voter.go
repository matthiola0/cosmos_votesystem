package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"votesystem/x/identity/keeper"
	"votesystem/x/identity/types"
)

func SimulateMsgRegisterVoter(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterVoter{
			Voter: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterVoter simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterVoter simulation not implemented"), nil, nil
	}
}
