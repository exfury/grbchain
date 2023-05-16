package rpc

import (
	"github.com/exfury/grbchain/libs/tendermint/crypto/merkle"
)

func defaultProofRuntime() *merkle.ProofRuntime {
	prt := merkle.NewProofRuntime()
	prt.RegisterOpDecoder(
		merkle.ProofOpSimpleValue,
		merkle.SimpleValueOpDecoder,
	)
	return prt
}
