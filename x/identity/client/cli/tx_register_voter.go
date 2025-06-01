package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"votesystem/x/identity/types"
)

var _ = strconv.Itoa(0)

func CmdRegisterVoter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-voter [identity-proof-hash]",
		Short: "Broadcast message registerVoter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIdentityProofHash := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterVoter(
				clientCtx.GetFromAddress().String(),
				argIdentityProofHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
