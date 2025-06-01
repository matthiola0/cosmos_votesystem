package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"votesystem/x/voting/types"
)

var _ = strconv.Itoa(0)

func CmdCastVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cast-vote [candidate-id] [election-id]",
		Short: "Broadcast message castVote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCandidateID := args[0]
			argElectionID := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCastVote(
				clientCtx.GetFromAddress().String(),
				argCandidateID,
				argElectionID,
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
