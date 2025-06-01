package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCastVote = "cast_vote"

var _ sdk.Msg = &MsgCastVote{}

func NewMsgCastVote(voter string, candidateID string, electionID string) *MsgCastVote {
	return &MsgCastVote{
		Voter:       voter,
		CandidateID: candidateID,
		ElectionID:  electionID,
	}
}

func (msg *MsgCastVote) Route() string {
	return RouterKey
}

func (msg *MsgCastVote) Type() string {
	return TypeMsgCastVote
}

func (msg *MsgCastVote) GetSigners() []sdk.AccAddress {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voter}
}

func (msg *MsgCastVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCastVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address (%s)", err)
	}
	return nil
}
