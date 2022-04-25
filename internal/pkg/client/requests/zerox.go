package requests

import (
	"fmt"
	"math/big"
)

type ZeroX struct {
	BuyToken   string
	SellToken  string
	SellAmount big.Int
	PairInfo   string
}

func DefaultZeroXRequest(amount big.Int) *ZeroX {
	return &ZeroX{
		SellAmount: amount,
	}
}

func (protocol *ZeroX) ParseRequest() string {
	res := fmt.Sprintf("https://api.0x.org/swap/v1/quote?buyToken=%s&sellToken=%s&sellAmount=%s", protocol.BuyToken, protocol.SellToken, protocol.SellAmount.String())
	return res
}

func (protocol *ZeroX) RequestInfo() string {
	return fmt.Sprintf("ZeroX %s:%s", protocol.SellAmount.String(), protocol.PairInfo)
}

func (protocol *ZeroX) SetPairs(pair int) {
	switch pair {
	case -3:
		protocol.SellToken = protocol.ETH().Name
		protocol.BuyToken = protocol.APE().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.APE().Name)
	case -2:
		protocol.SellToken = protocol.USDT().Name
		protocol.BuyToken = protocol.USDC().Name
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDT().Name, protocol.USDC().Name)
	case -1:
		protocol.SellToken = protocol.USDC().Name
		protocol.BuyToken = protocol.ETH().Name
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.ETH().Name)
	case 1:
		protocol.SellToken = protocol.ETH().Name
		protocol.BuyToken = protocol.USDC().Name
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.USDC().Name)
	case 2:
		protocol.SellToken = protocol.USDC().Name
		protocol.BuyToken = protocol.USDT().Name
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.USDT().Name)
	case 3:
		protocol.SellToken = protocol.APE().Address
		protocol.BuyToken = protocol.ETH().Name
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.APE().Name, protocol.ETH().Name)
	}
}

func (protocol *ZeroX) ETH() Token { // ETH
	return Token{
		Name:     "ETH",
		Address:  "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		Decimals: 18,
	}
}

func (protocol *ZeroX) USDC() Token {
	return Token{
		Name:     "USDC",
		Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals: 6,
	}
}

func (protocol *ZeroX) USDT() Token {
	return Token{
		Name:     "USDT",
		Address:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals: 6,
	}
}
func (protocol *ZeroX) APE() Token {
	return Token{
		Name:     "APE",
		Address:  "0x4d224452801aced8b2f0aebe155379bb5d594381",
		Decimals: 18,
	}
}
