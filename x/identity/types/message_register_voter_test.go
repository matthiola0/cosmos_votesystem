package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"votesystem/testutil/sample"
)

func TestMsgRegisterVoter_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterVoter
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterVoter{
				Voter: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRegisterVoter{
				Voter: sample.AccAddress(),
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
