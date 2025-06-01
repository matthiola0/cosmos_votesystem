package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VoteRecordList:    []VoteRecord{},
		HasVotedVoterList: []HasVotedVoter{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in voteRecord
	voteRecordIdMap := make(map[uint64]bool)
	voteRecordCount := gs.GetVoteRecordCount()
	for _, elem := range gs.VoteRecordList {
		if _, ok := voteRecordIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for voteRecord")
		}
		if elem.Id >= voteRecordCount {
			return fmt.Errorf("voteRecord id should be lower or equal than the last id")
		}
		voteRecordIdMap[elem.Id] = true
	}
	// Check for duplicated index in hasVotedVoter
	hasVotedVoterIndexMap := make(map[string]struct{})

	for _, elem := range gs.HasVotedVoterList {
		index := string(HasVotedVoterKey(elem.Index))
		if _, ok := hasVotedVoterIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for hasVotedVoter")
		}
		hasVotedVoterIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
