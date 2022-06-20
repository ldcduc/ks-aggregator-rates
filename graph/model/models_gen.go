// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type KyberSwapPrice struct {
	ID           string  `json:"Id"`
	Pair         string  `json:"Pair"`
	InputToken   string  `json:"InputToken"`
	OutputToken  string  `json:"OutputToken"`
	InputAmount  string  `json:"InputAmount"`
	OutputAmount string  `json:"OutputAmount"`
	TotalGas     float64 `json:"TotalGas"`
	GasPriceGwei string  `json:"GasPriceGwei"`
	GasUsd       float64 `json:"GasUsd"`
	Timestamp    int     `json:"Timestamp"`
}

type OneInchPrice struct {
	ID               string  `json:"Id"`
	Pair             string  `json:"Pair"`
	FromTokenAddress string  `json:"FromTokenAddress"`
	ToTokenAddress   string  `json:"ToTokenAddress"`
	ToTokenAmount    string  `json:"ToTokenAmount"`
	FromTokenAmount  string  `json:"FromTokenAmount"`
	EstimatedGas     float64 `json:"EstimatedGas"`
	Timestamp        int     `json:"Timestamp"`
}

type ParaSwapPrice struct {
	ID          string `json:"Id"`
	BlockNumber int    `json:"BlockNumber"`
	Pair        string `json:"Pair"`
	SrcToken    string `json:"SrcToken"`
	SrcAmount   string `json:"SrcAmount"`
	DestToken   string `json:"DestToken"`
	DestAmount  string `json:"DestAmount"`
	Timestamp   string `json:"Timestamp"`
}

type ZeroXPrice struct {
	ID               string `json:"Id"`
	Pair             string `json:"Pair"`
	BuyTokenAddress  string `json:"BuyTokenAddress"`
	SellTOkenAddress string `json:"SellTOkenAddress"`
	BuyAmount        string `json:"BuyAmount"`
	SellAmount       string `json:"SellAmount"`
	Timestamp        string `json:"Timestamp"`
}
