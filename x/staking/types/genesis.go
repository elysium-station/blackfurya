package types

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	blackfury "github.com/blackfury-zone/blackfury/types"
)

// DefaultGenesis gets the raw genesis raw message for testing
func DefaultGenesis() *stakingtypes.GenesisState {
	params := stakingtypes.DefaultParams()
	params.BondDenom = blackfury.BaseDenom
	return &stakingtypes.GenesisState{
		Params: params,
	}
}
