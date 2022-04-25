package requests

import (
	"fmt"
	"math/big"
)

type KyberSwap struct {
	TokenIn    string
	TokenOut   string
	AmountIn   big.Int
	SaveGas    bool // 0 = maximize return, 1 = no splitting
	GasInclude bool // 0 = doesn't account for transaction fee in finding optimal path, 1= accounts for transaction fee in finding optimal path
	/*	dexes      string
	 */
	PairInfo string
}

func DefaultKyberSwapRequest(amount big.Int) *KyberSwap {
	res := KyberSwap{
		AmountIn:   amount,
		SaveGas:    false,
		GasInclude: false,
	}
	return &res
}

func (protocol *KyberSwap) ParseRequest() string {
	var saveGasString, gasIncludeString string
	if protocol.SaveGas {
		saveGasString = "1"
	} else {
		saveGasString = "0"
	}
	if protocol.GasInclude {
		gasIncludeString = "1"
	} else {
		gasIncludeString = "0"
	}
	res := fmt.Sprintf("https://aggregator-api.kyberswap.com/ethereum/route?tokenIn=%s&tokenOut=%s&amountIn=%s&saveGas=%s&gasInclude=%s", protocol.TokenIn, protocol.TokenOut, protocol.AmountIn.String(), saveGasString, gasIncludeString)
	return res
}

func (protocol *KyberSwap) RequestInfo() string {
	return fmt.Sprintf("KyberSwap %s:%s", protocol.AmountIn.String(), protocol.PairInfo)
}

func (protocol *KyberSwap) SetPairs(pair int) {
	switch pair {
	case -3:
		protocol.TokenIn = protocol.ETH().Address
		protocol.TokenOut = protocol.APE().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.APE().Name)
	case -2:
		protocol.TokenIn = protocol.USDT().Address
		protocol.TokenOut = protocol.USDC().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDT().Name, protocol.USDC().Name)
	case -1:
		protocol.TokenIn = protocol.USDC().Address
		protocol.TokenOut = protocol.ETH().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.ETH().Name)
	case 1:
		protocol.TokenIn = protocol.ETH().Address
		protocol.TokenOut = protocol.USDC().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.USDC().Name)
	case 2:
		protocol.TokenIn = protocol.USDC().Address
		protocol.TokenOut = protocol.USDT().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.USDT().Name)
	case 3:
		protocol.TokenIn = protocol.APE().Address
		protocol.TokenOut = protocol.ETH().Address
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.APE().Name, protocol.ETH().Name)
	}
}

func (protocol *KyberSwap) ETH() Token { // WETH
	return Token{
		Name:     "WETH",
		Address:  "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		Decimals: 18,
	}
}

func (protocol *KyberSwap) USDC() Token {
	return Token{
		Name:     "USDC",
		Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals: 6,
	}
}

func (protocol *KyberSwap) USDT() Token {
	return Token{
		Name:     "USDT",
		Address:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals: 6,
	}
}
func (protocol *KyberSwap) APE() Token {
	return Token{
		Name:     "APE",
		Address:  "0x4d224452801aced8b2f0aebe155379bb5d594381",
		Decimals: 18,
	}
}
