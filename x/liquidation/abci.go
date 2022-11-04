package liquidation

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	"github.com/comdex-official/comdex/x/liquidation/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)

	err := k.LiquidateVaults(ctx)
	if err != nil {
		ctx.Logger().Error("error in LiquidateVaults")
	}
	err = k.LiquidateBorrows(ctx)
	if err != nil {
		ctx.Logger().Error("error in LiquidateBorrows")
	}
}
