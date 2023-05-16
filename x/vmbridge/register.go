package vmbridge

import (
	"github.com/exfury/grbchain/libs/cosmos-sdk/codec"
	"github.com/exfury/grbchain/libs/cosmos-sdk/types/module"
	"github.com/exfury/grbchain/x/vmbridge/keeper"
	"github.com/exfury/grbchain/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
