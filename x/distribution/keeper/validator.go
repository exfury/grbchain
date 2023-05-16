package keeper

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"

	"github.com/exfury/grbchain/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	k.initializeValidatorDistrProposal(ctx, val)
	return
}
