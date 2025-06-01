package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"votesystem/x/identity/types"
)

// GetRegisteredVoterCount get the total number of registeredVoter
func (k Keeper) GetRegisteredVoterCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegisteredVoterCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRegisteredVoterCount set the total number of registeredVoter
func (k Keeper) SetRegisteredVoterCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegisteredVoterCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRegisteredVoter appends a registeredVoter in the store with a new id and update the count
func (k Keeper) AppendRegisteredVoter(
	ctx sdk.Context,
	registeredVoter types.RegisteredVoter,
) uint64 {
	// Create the registeredVoter
	count := k.GetRegisteredVoterCount(ctx)

	// Set the ID of the appended value
	registeredVoter.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisteredVoterKey))
	appendedValue := k.cdc.MustMarshal(&registeredVoter)
	store.Set(GetRegisteredVoterIDBytes(registeredVoter.Id), appendedValue)

	// Update registeredVoter count
	k.SetRegisteredVoterCount(ctx, count+1)

	return count
}

// SetRegisteredVoter set a specific registeredVoter in the store
func (k Keeper) SetRegisteredVoter(ctx sdk.Context, registeredVoter types.RegisteredVoter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisteredVoterKey))
	b := k.cdc.MustMarshal(&registeredVoter)
	store.Set(GetRegisteredVoterIDBytes(registeredVoter.Id), b)
}

// GetRegisteredVoter returns a registeredVoter from its id
func (k Keeper) GetRegisteredVoter(ctx sdk.Context, id uint64) (val types.RegisteredVoter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisteredVoterKey))
	b := store.Get(GetRegisteredVoterIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegisteredVoter removes a registeredVoter from the store
func (k Keeper) RemoveRegisteredVoter(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisteredVoterKey))
	store.Delete(GetRegisteredVoterIDBytes(id))
}

// GetAllRegisteredVoter returns all registeredVoter
func (k Keeper) GetAllRegisteredVoter(ctx sdk.Context) (list []types.RegisteredVoter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisteredVoterKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegisteredVoter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRegisteredVoterIDBytes returns the byte representation of the ID
func GetRegisteredVoterIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRegisteredVoterIDFromBytes returns ID in uint64 format from a byte array
func GetRegisteredVoterIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
