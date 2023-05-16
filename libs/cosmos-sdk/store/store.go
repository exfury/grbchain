package store

import (
	dbm "github.com/exfury/grbchain/libs/tm-db"

	"github.com/exfury/grbchain/libs/cosmos-sdk/store/cache"
	"github.com/exfury/grbchain/libs/cosmos-sdk/store/rootmulti"
	"github.com/exfury/grbchain/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
