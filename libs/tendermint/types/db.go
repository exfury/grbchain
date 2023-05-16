package types

import dbm "github.com/exfury/grbchain/libs/tm-db"

// DBBackend This is set at compile time.
var DBBackend = string(dbm.GoLevelDBBackend)
