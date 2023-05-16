package token

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	"github.com/exfury/grbchain/x/common/perf"
	"github.com/exfury/grbchain/x/token/types"
)

// BeginBlocker is called when dapp handles with abci::BeginBlock
func beginBlocker(ctx sdk.Context, keeper Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
