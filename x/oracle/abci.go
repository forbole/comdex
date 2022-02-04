package oracle

import (
	"fmt"
	"github.com/comdex-official/comdex/x/oracle/keeper"
	"github.com/comdex-official/comdex/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)


func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	a := k.GetAssets(ctx)
	fmt.Println(a)
	fmt.Println("printing Assets")
	if ctx.BlockHeight()>=349{
		assets := k.GetAssets(ctx)
		for _, asset := range assets{
			k.SetRates(ctx, asset.Name)
			k.SetMarketForAsset(ctx, asset.Id, asset.Name)
			rate,_ := k.GetRates(ctx, asset.Name)
			var (
				market = types.Market{
					Symbol:   asset.Name,
					ScriptID: 112,
					Rates:    rate,
				}
			)
			k.SetMarket(ctx, market)
		}
	}
}
