package node

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/exfury/grbchain/libs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
