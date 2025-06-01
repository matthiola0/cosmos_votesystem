package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"votesystem/testutil/sample"
)

func TestMsgIssueVoteCoin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgIssueVoteCoin
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgIssueVoteCoin{
				Admin: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgIssueVoteCoin{
				Admin: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
