package distribution

import (
	abci "github.com/exfury/grbchain/libs/tendermint/abci/types"
	tmtypes "github.com/exfury/grbchain/libs/tendermint/types"
	"github.com/exfury/grbchain/x/common"

	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	"github.com/exfury/grbchain/x/distribution/keeper"
)

// BeginBlocker set the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// determine the total power signing the block
	var previousTotalPower int64
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		previousTotalPower += voteInfo.Validator.Power
	}

	// TODO this is Tendermint-dependent
	// ref https://github.com/cosmos/cosmos-sdk/issues/3095
	if ctx.BlockHeight() > tmtypes.GetStartBlockHeight()+1 {
		previousProposer := k.GetPreviousProposerConsAddr(ctx)
		/* allocate tokens by grbchain custom rule */
		if k.StakingKeeper().ParamsConsensusType(ctx) == common.PoA {
			k.PoAAllocateTokens(ctx, req.LastCommitInfo.GetVotes())
		} else {
			k.AllocateTokens(ctx, previousTotalPower, previousProposer, req.LastCommitInfo.GetVotes())
		}
	}

	// record the proposer for when we payout on the next block
	consAddr := sdk.ConsAddress(req.Header.ProposerAddress)
	k.SetPreviousProposerConsAddr(ctx, consAddr)
}