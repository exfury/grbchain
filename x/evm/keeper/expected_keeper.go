package keeper

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	govtypes "github.com/exfury/grbchain/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
