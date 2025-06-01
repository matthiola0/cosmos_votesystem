package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identity module sentinel errors
var (
	ErrVoterAlreadyRegistered = sdkerrors.Register(ModuleName, 1100, "voter already registered")
    ErrVoterNotRegistered = sdkerrors.Register(ModuleName, 1101, "voter not registered")
	ErrInsufficientVoteCoin = sdkerrors.Register(ModuleName, 1200, "insufficient vote coin")
	ErrAlreadyVoted         = sdkerrors.Register(ModuleName, 1201, "voter has already voted in this election")

)
