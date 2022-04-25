package requests

type ClientRequest interface {
	ParseRequest() string
	RequestInfo() string
	SetPairs(pair int)
}

type Token struct {
	Name     string
	Address  string
	Decimals int
}

type Tokens interface {
	ETH()
	USDC()
	USDT()
	APE()
}
