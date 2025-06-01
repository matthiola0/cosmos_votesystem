package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/voting module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInsufficientVoteCoin = sdkerrors.Register(ModuleName, 1200, "insufficient vote coin")
	ErrAlreadyVoted         = sdkerrors.Register(ModuleName, 1201, "voter has already voted in this election")
)
