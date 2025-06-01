package keeper

import (
	"votesystem/x/identity/types"
)

var _ types.QueryServer = Keeper{}
