package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"votesystem/x/identity/types"
)

var _ = strconv.Itoa(0)

func CmdGetVoterStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-voter-status [voter-address]",
		Short: "Query getVoterStatus",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqVoterAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetVoterStatusRequest{

				VoterAddress: reqVoterAddress,
			}

			res, err := queryClient.GetVoterStatus(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
