package types

import (
	amino "github.com/tendermint/go-amino"

	"github.com/exfury/grbchain/libs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
