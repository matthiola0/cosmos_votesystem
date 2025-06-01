package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"votesystem/x/voting/types"
)

// SetHasVotedVoter set a specific hasVotedVoter in the store from its index
func (k Keeper) SetHasVotedVoter(ctx sdk.Context, hasVotedVoter types.HasVotedVoter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HasVotedVoterKeyPrefix))
	b := k.cdc.MustMarshal(&hasVotedVoter)
	store.Set(types.HasVotedVoterKey(
		hasVotedVoter.Index,
	), b)
}

// GetHasVotedVoter returns a hasVotedVoter from its index
func (k Keeper) GetHasVotedVoter(
	ctx sdk.Context,
	index string,

) (val types.HasVotedVoter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HasVotedVoterKeyPrefix))

	b := store.Get(types.HasVotedVoterKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHasVotedVoter removes a hasVotedVoter from the store
func (k Keeper) RemoveHasVotedVoter(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HasVotedVoterKeyPrefix))
	store.Delete(types.HasVotedVoterKey(
		index,
	))
}

// GetAllHasVotedVoter returns all hasVotedVoter
func (k Keeper) GetAllHasVotedVoter(ctx sdk.Context) (list []types.HasVotedVoter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HasVotedVoterKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HasVotedVoter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
