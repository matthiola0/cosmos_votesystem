package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"votesystem/x/voting/types"
)

func CmdListHasVotedVoter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-has-voted-voter",
		Short: "list all HasVotedVoter",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHasVotedVoterRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HasVotedVoterAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHasVotedVoter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-has-voted-voter [index]",
		Short: "shows a HasVotedVoter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetHasVotedVoterRequest{
				Index: argIndex,
			}

			res, err := queryClient.HasVotedVoter(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
