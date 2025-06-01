package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// HasVotedVoterKeyPrefix is the prefix to retrieve all HasVotedVoter
	HasVotedVoterKeyPrefix = "HasVotedVoter/value/"
)

// HasVotedVoterKey returns the store key to retrieve a HasVotedVoter from the index fields
func HasVotedVoterKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
