package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RewardContractList: []RewardContract{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in rewardContract
	rewardContractIdMap := make(map[uint64]bool)
	rewardContractCount := gs.GetRewardContractCount()
	for _, elem := range gs.RewardContractList {
		if _, ok := rewardContractIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for rewardContract")
		}
		if elem.Id >= rewardContractCount {
			return fmt.Errorf("rewardContract id should be lower or equal than the last id")
		}
		rewardContractIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
