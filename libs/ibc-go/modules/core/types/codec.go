package types

import (
	"github.com/exfury/grbchain/libs/cosmos-sdk/codec"
	codectypes "github.com/exfury/grbchain/libs/cosmos-sdk/codec/types"
	clienttypes "github.com/exfury/grbchain/libs/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/exfury/grbchain/libs/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/exfury/grbchain/libs/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/exfury/grbchain/libs/ibc-go/modules/core/23-commitment/types"
	solomachinetypes "github.com/exfury/grbchain/libs/ibc-go/modules/light-clients/06-solomachine/types"
	ibctmtypes "github.com/exfury/grbchain/libs/ibc-go/modules/light-clients/07-tendermint/types"
	localhosttypes "github.com/exfury/grbchain/libs/ibc-go/modules/light-clients/09-localhost/types"
)

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)

	solomachinetypes.RegisterInterfaces(registry)
	ibctmtypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}

func RegisterCodec(cdc *codec.Codec) {
	clienttypes.RegisterCodec(cdc)
	ibctmtypes.RegisterCodec(cdc)
}
