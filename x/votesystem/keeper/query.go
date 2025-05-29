package keeper

import (
	"votesystem/x/votesystem/types"
)

var _ types.QueryServer = Keeper{}
