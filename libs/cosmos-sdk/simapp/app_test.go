package simapp

import (
	"os"
	"testing"

	"github.com/exfury/grbchain/libs/tendermint/libs/log"
	dbm "github.com/exfury/grbchain/libs/tm-db"
	"github.com/stretchr/testify/require"

	"github.com/exfury/grbchain/libs/cosmos-sdk/codec"

	abci "github.com/exfury/grbchain/libs/tendermint/abci/types"
)

func TestSimAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, 0)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit(abci.RequestCommit{})

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, 0)
	_, _, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, 0)

	for acc := range maccPerms {
		require.Equal(t, !allowedReceivingModAcc[acc], app.BankKeeper.BlacklistedAddr(app.SupplyKeeper.GetModuleAddress(acc)))
	}
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}