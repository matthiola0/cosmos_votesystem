package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"votesystem/x/votecoin/types"
)

var _ = strconv.Itoa(0)

func CmdIssueVoteCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-vote-coin [recipient] [amount]",
		Short: "Broadcast message issueVoteCoin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRecipient := args[0]
			argAmount, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueVoteCoin(
				clientCtx.GetFromAddress().String(),
				argRecipient,
				argAmount,
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
