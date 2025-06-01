package keeper

import (
	"votesystem/x/votecoin/types"
)

var _ types.QueryServer = Keeper{}
