package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RegisteredVoterList: []RegisteredVoter{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in registeredVoter
	registeredVoterIdMap := make(map[uint64]bool)
	registeredVoterCount := gs.GetRegisteredVoterCount()
	for _, elem := range gs.RegisteredVoterList {
		if _, ok := registeredVoterIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for registeredVoter")
		}
		if elem.Id >= registeredVoterCount {
			return fmt.Errorf("registeredVoter id should be lower or equal than the last id")
		}
		registeredVoterIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
