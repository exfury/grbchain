package system

const (
	Chain = "grbchain"
	TestnetPrefix = "grbchaintest"
	AppName = "OKBChain"
	Server = Chain+"d"
	Client = Chain+"cli"
	ServerHome = "$HOME/."+Server
	ClientHome = "$HOME/."+Client
	ServerLog = Server+".log"
	EnvPrefix = "GRBCHAIN"
	CoinType uint32 = 60
	Currency = "okb"
)
