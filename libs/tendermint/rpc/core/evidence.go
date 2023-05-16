package core

import (
	"github.com/exfury/grbchain/libs/tendermint/evidence"
	ctypes "github.com/exfury/grbchain/libs/tendermint/rpc/core/types"
	rpctypes "github.com/exfury/grbchain/libs/tendermint/rpc/jsonrpc/types"
	"github.com/exfury/grbchain/libs/tendermint/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.tendermint.com/master/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := env.EvidencePool.AddEvidence(ev)
	if _, ok := err.(evidence.ErrEvidenceAlreadyStored); err == nil || ok {
		return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
	}
	return nil, err
}
