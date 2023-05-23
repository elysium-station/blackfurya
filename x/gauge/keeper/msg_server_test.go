package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/blackfury-zone/blackfury/testutil/keeper"
	"github.com/blackfury-zone/blackfury/x/gauge/keeper"
	"github.com/blackfury-zone/blackfury/x/gauge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GaugeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
