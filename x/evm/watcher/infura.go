package watcher

import "github.com/exfury/grbchain/x/evm/types"

type InfuraKeeper interface {
	OnSaveTransactionReceipt(TransactionReceipt)
	OnSaveBlock(types.Block)
	OnSaveTransaction(Transaction)
	OnSaveContractCode(address string, code []byte)
}
