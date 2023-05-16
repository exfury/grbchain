package proto

import (
	"os"
	"testing"

	"github.com/exfury/grbchain/libs/cosmos-sdk/store"
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	abci "github.com/exfury/grbchain/libs/tendermint/abci/types"
	"github.com/exfury/grbchain/libs/tendermint/libs/log"
	dbm "github.com/exfury/grbchain/libs/tm-db"
	"github.com/stretchr/testify/require"
)

func createTestInput(t *testing.T) (sdk.Context, ProtocolKeeper) {
	keyMain := sdk.NewKVStoreKey("main")

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyMain, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewTMLogger(os.Stdout))

	keeper := NewProtocolKeeper(keyMain)

	return ctx, keeper
}
