package types

import (
	interfacetypes "github.com/exfury/grbchain/libs/cosmos-sdk/codec/types"
	txmsg "github.com/exfury/grbchain/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/exfury/grbchain/libs/cosmos-sdk/types/msgservice"
)

func RegisterInterface(registry interfacetypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*txmsg.Msg)(nil),
		&MsgSendToEvm{},
		&MsgCallToEvm{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
