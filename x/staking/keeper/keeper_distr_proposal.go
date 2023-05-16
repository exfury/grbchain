package keeper

import (
	"github.com/exfury/grbchain/x/staking/exported"

	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
)

func (k Keeper) Delegation(ctx sdk.Context, delAddr sdk.AccAddress, address2 sdk.ValAddress) exported.DelegatorI {
	delegator, found := k.GetDelegator(ctx, delAddr)
	if !found {
		return nil
	}

	return delegator
}
