package requests

import (
	"fmt"
	"math/big"
	"strconv"
)

type ParaSwap struct {
	SrcToken     string
	DestToken    string
	Amount       big.Int
	SrcDecimals  int
	DestDecimals int
	Side         string
	Network      int
	PairInfo     string
}

func DefaultParaSwapRequest(amount big.Int) *ParaSwap {
	return &ParaSwap{
		Amount:  amount,
		Side:    "SELL",
		Network: 1,
	}
}

func (protocol *ParaSwap) ParseRequest() string {
	res := fmt.Sprintf("https://apiv5.paraswap.io/prices/?srcToken=%s&destToken=%s&amount=%s&srcDecimals=%s&destDecimals=%s&side=%s&network=%s", protocol.SrcToken, protocol.DestToken, protocol.Amount.String(), strconv.Itoa(protocol.SrcDecimals), strconv.Itoa(protocol.DestDecimals), protocol.Side, strconv.Itoa(protocol.Network))
	return res
}

func (protocol *ParaSwap) RequestInfo() string {
	return fmt.Sprintf("ParaSwap %s:%s", protocol.Amount.String(), protocol.PairInfo)
}

func (protocol *ParaSwap) SetPairs(pair int) {
	switch pair {
	case -3:
		protocol.SrcToken = protocol.ETH().Address
		protocol.DestToken = protocol.APE().Address
		protocol.SrcDecimals = protocol.ETH().Decimals
		protocol.DestDecimals = protocol.APE().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.APE().Name)
	case -2:
		protocol.SrcToken = protocol.USDT().Address
		protocol.DestToken = protocol.USDC().Address
		protocol.SrcDecimals = protocol.USDT().Decimals
		protocol.DestDecimals = protocol.USDC().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDT().Name, protocol.USDC().Name)
	case -1:
		protocol.SrcToken = protocol.USDC().Address
		protocol.DestToken = protocol.ETH().Address
		protocol.SrcDecimals = protocol.USDC().Decimals
		protocol.DestDecimals = protocol.ETH().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.ETH().Name)
	case 1:
		protocol.SrcToken = protocol.ETH().Address
		protocol.DestToken = protocol.USDC().Address
		protocol.SrcDecimals = protocol.ETH().Decimals
		protocol.DestDecimals = protocol.USDC().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.USDC().Name)
	case 2:
		protocol.SrcToken = protocol.USDC().Address
		protocol.DestToken = protocol.USDT().Address
		protocol.SrcDecimals = protocol.USDC().Decimals
		protocol.DestDecimals = protocol.USDT().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.USDT().Name)
	case 3:
		protocol.SrcToken = protocol.APE().Address
		protocol.DestToken = protocol.ETH().Address
		protocol.SrcDecimals = protocol.APE().Decimals
		protocol.DestDecimals = protocol.ETH().Decimals
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.APE().Name, protocol.ETH().Name)
	}
}

func (protocol *ParaSwap) ETH() Token { // ETH
	return Token{
		Name:     "ETH",
		Address:  "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		Decimals: 18,
	}
}

func (protocol *ParaSwap) USDC() Token {
	return Token{
		Name:     "USDC",
		Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals: 6,
	}
}

func (protocol *ParaSwap) USDT() Token {
	return Token{
		Name:     "USDT",
		Address:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals: 6,
	}
}

func (protocol *ParaSwap) APE() Token {
	return Token{
		Name:     "APE",
		Address:  "0x4d224452801aced8b2f0aebe155379bb5d594381",
		Decimals: 18,
	}
}
