package keeper

import (
	"votesystem/x/voting/types"
)

var _ types.QueryServer = Keeper{}
