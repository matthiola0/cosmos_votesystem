package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterVoter = "register_voter"

var _ sdk.Msg = &MsgRegisterVoter{}

func NewMsgRegisterVoter(voter string, identityProofHash string) *MsgRegisterVoter {
	return &MsgRegisterVoter{
		Voter:             voter,
		IdentityProofHash: identityProofHash,
	}
}

func (msg *MsgRegisterVoter) Route() string {
	return RouterKey
}

func (msg *MsgRegisterVoter) Type() string {
	return TypeMsgRegisterVoter
}

func (msg *MsgRegisterVoter) GetSigners() []sdk.AccAddress {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voter}
}

func (msg *MsgRegisterVoter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterVoter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address (%s)", err)
	}
	return nil
}
