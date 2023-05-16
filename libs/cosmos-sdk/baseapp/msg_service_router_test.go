package baseapp_test

import (
	chaincodec "github.com/exfury/grbchain/app/codec"
	"github.com/exfury/grbchain/libs/ibc-go/testing/simapp"
	"github.com/exfury/grbchain/x/evm"
	"os"
	"testing"

	"github.com/exfury/grbchain/libs/tendermint/libs/log"

	dbm "github.com/exfury/grbchain/libs/tm-db"
	"github.com/stretchr/testify/require"

	"github.com/exfury/grbchain/libs/cosmos-sdk/baseapp"

	"github.com/exfury/grbchain/x/evm/types/testdata"
)

func TestRegisterMsgService(t *testing.T) {
	db := dbm.NewMemDB()

	// Create an encoding config that doesn't register testdata Msg services.
	codecProxy, interfaceRegistry := chaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Register testdata Msg services, and rerun `RegisterService`.
	testdata.RegisterInterfaces(interfaceRegistry)
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}

func TestRegisterMsgServiceTwice(t *testing.T) {
	// Setup baseapp.
	db := dbm.NewMemDB()
	codecProxy, interfaceRegistry := chaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	testdata.RegisterInterfaces(interfaceRegistry)

	// First time registering service shouldn't panic.
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Second time should panic.
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}
