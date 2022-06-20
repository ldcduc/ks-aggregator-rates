package requests

import "ks-aggregator-rates/internal/pkg/client/models"

type ClientRequest interface {
	ParseRequest() string
	ParseResponse(body interface{}, timestamp int64) models.Response
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
