package requests

import (
	"fmt"
	"ks-aggregator-rates/internal/pkg/client/models"
	"math/big"
)

type OneInch struct {
	FromTokenAddress string
	ToTokenAddress   string
	Amount           big.Int
	PairInfo         string
}

func DefaultOneInchRequest(amount big.Int) *OneInch {
	return &OneInch{
		Amount: amount,
	}
}

func (protocol *OneInch) ParseRequest() string {
	res := fmt.Sprintf("https://api.1inch.io/v4.0/1/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s", protocol.FromTokenAddress, protocol.ToTokenAddress, protocol.Amount.String())
	return res
}

func (protocol *OneInch) ParseResponse(body interface{}, timestamp int64) models.Response {
	data := body.(map[string]interface{})
	return models.OneInch{
		Pair:            protocol.PairInfo,
		FromToken:       models.IntoToken(data["fromToken"].(map[string]interface{})),
		ToToken:         models.IntoToken(data["toToken"].(map[string]interface{})),
		ToTokenAmount:   data["toTokenAmount"].(string),
		FromTokenAmount: data["fromTokenAmount"].(string),
		EstimatedGas:    data["estimatedGas"].(float64),
		Timestamp:       timestamp,
	}
}

func (protocol *OneInch) RequestInfo() string {
	return fmt.Sprintf("OneInch %s:%s", protocol.Amount.String(), protocol.PairInfo)
}

func (protocol *OneInch) SetPairs(pair int) {
	switch pair {
	case -3:
		protocol.FromTokenAddress = protocol.ETH().Address // ETH
		protocol.ToTokenAddress = protocol.APE().Address   // APE
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.APE().Name)
	case -2:
		protocol.FromTokenAddress = protocol.USDT().Address // USDT
		protocol.ToTokenAddress = protocol.USDC().Address   // USDC
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDT().Name, protocol.USDC().Name)
	case -1:
		protocol.FromTokenAddress = protocol.USDC().Address // USDC
		protocol.ToTokenAddress = protocol.ETH().Address    // ETH
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.ETH().Name)
	case 1:
		protocol.FromTokenAddress = protocol.ETH().Address // ETH
		protocol.ToTokenAddress = protocol.USDC().Address  // USDC
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.ETH().Name, protocol.USDC().Name)
	case 2:
		protocol.FromTokenAddress = protocol.USDC().Address // USDC
		protocol.ToTokenAddress = protocol.USDT().Address   // USDT
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.USDC().Name, protocol.USDT().Name)
	case 3:
		protocol.FromTokenAddress = protocol.APE().Address // APE
		protocol.ToTokenAddress = protocol.ETH().Address   // ETH
		protocol.PairInfo = fmt.Sprintf("%s:%s", protocol.APE().Name, protocol.ETH().Name)
	}
}

func (protocol *OneInch) ETH() Token { // ETH
	return Token{
		Name:     "ETH",
		Address:  "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		Decimals: 18,
	}
}

func (protocol *OneInch) USDC() Token {
	return Token{
		Name:     "USDC",
		Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals: 6,
	}
}

func (protocol *OneInch) USDT() Token {
	return Token{
		Name:     "USDT",
		Address:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals: 6,
	}
}

func (protocol *OneInch) APE() Token {
	return Token{
		Name:     "APE",
		Address:  "0x4d224452801aced8b2f0aebe155379bb5d594381",
		Decimals: 18,
	}
}
