package keeper

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	"github.com/exfury/grbchain/libs/cosmos-sdk/x/supply/exported"

	"github.com/exfury/grbchain/x/distribution/types"
)

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}
