package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"votesystem/x/voting/types"
)

var _ = strconv.Itoa(0)

func CmdGetElectionResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-election-result [election-id]",
		Short: "Query getElectionResult",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqElectionID := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetElectionResultRequest{

				ElectionID: reqElectionID,
			}

			res, err := queryClient.GetElectionResult(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
