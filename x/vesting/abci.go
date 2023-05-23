package vesting

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	blackfury "github.com/blackfury-zone/blackfury/types"
	"github.com/blackfury-zone/blackfury/x/vesting/keeper"
	"github.com/blackfury-zone/blackfury/x/vesting/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if blackfury.IsPeriodLastBlock(ctx, types.ClaimVestedPeriod) {
		k.ClaimVested(ctx)
	}
}
