package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"votesystem/x/voting/types"
)

// GetVoteRecordCount get the total number of voteRecord
func (k Keeper) GetVoteRecordCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VoteRecordCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVoteRecordCount set the total number of voteRecord
func (k Keeper) SetVoteRecordCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VoteRecordCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVoteRecord appends a voteRecord in the store with a new id and update the count
func (k Keeper) AppendVoteRecord(
	ctx sdk.Context,
	voteRecord types.VoteRecord,
) uint64 {
	// Create the voteRecord
	count := k.GetVoteRecordCount(ctx)

	// Set the ID of the appended value
	voteRecord.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteRecordKey))
	appendedValue := k.cdc.MustMarshal(&voteRecord)
	store.Set(GetVoteRecordIDBytes(voteRecord.Id), appendedValue)

	// Update voteRecord count
	k.SetVoteRecordCount(ctx, count+1)

	return count
}

// SetVoteRecord set a specific voteRecord in the store
func (k Keeper) SetVoteRecord(ctx sdk.Context, voteRecord types.VoteRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteRecordKey))
	b := k.cdc.MustMarshal(&voteRecord)
	store.Set(GetVoteRecordIDBytes(voteRecord.Id), b)
}

// GetVoteRecord returns a voteRecord from its id
func (k Keeper) GetVoteRecord(ctx sdk.Context, id uint64) (val types.VoteRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteRecordKey))
	b := store.Get(GetVoteRecordIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVoteRecord removes a voteRecord from the store
func (k Keeper) RemoveVoteRecord(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteRecordKey))
	store.Delete(GetVoteRecordIDBytes(id))
}

// GetAllVoteRecord returns all voteRecord
func (k Keeper) GetAllVoteRecord(ctx sdk.Context) (list []types.VoteRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VoteRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVoteRecordIDBytes returns the byte representation of the ID
func GetVoteRecordIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVoteRecordIDFromBytes returns ID in uint64 format from a byte array
func GetVoteRecordIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
