package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIssueVoteCoin = "issue_vote_coin"

var _ sdk.Msg = &MsgIssueVoteCoin{}

func NewMsgIssueVoteCoin(admin string, recipient string, amount int32) *MsgIssueVoteCoin {
	return &MsgIssueVoteCoin{
		Admin:     admin,
		Recipient: recipient,
		Amount:    amount,
	}
}

func (msg *MsgIssueVoteCoin) Route() string {
	return RouterKey
}

func (msg *MsgIssueVoteCoin) Type() string {
	return TypeMsgIssueVoteCoin
}

func (msg *MsgIssueVoteCoin) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgIssueVoteCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueVoteCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}
