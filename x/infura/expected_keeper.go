package infura

import evm "github.com/exfury/grbchain/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
