package keeper

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	"github.com/exfury/grbchain/x/staking/types"
)

func calculateWeight(tokens sdk.Dec) types.Shares {
	return tokens
}

func SimulateWeight(tokens sdk.Dec) types.Shares {
	return calculateWeight(tokens)
}
