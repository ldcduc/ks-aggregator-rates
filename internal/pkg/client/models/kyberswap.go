package models

import (
	"gorm.io/gorm"
)

type KyberSwap struct {
	Pair         string
	InputToken   string
	OutputToken  string
	InputAmount  string
	OutputAmount string
	TotalGas     float64
	GasPriceGwei string
	GasUsd       float64
	AmountInUsd  float64
	AmountOutUsd float64
	ReceivedUsd  float64
	Timestamp    int64
	// Swaps        []interface{}
}

type KyberSwapPrice struct {
	Id           uint `gorm:"primaryKey"`
	Pair         string
	InputToken   string
	OutputToken  string
	InputAmount  string
	OutputAmount string
	TotalGas     float64
	GasPriceGwei string
	GasUsd       float64
	Timestamp    int64
}

func KyberSwapPriceTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("kyberswap")
	}
}

func (KyberSwapPrice) TableName() string {
	return "kyberswap"
}

func (price KyberSwap) ToPrice() interface{} {
	return &KyberSwapPrice{
		Pair:         price.Pair,
		InputToken:   price.InputToken,
		OutputToken:  price.OutputToken,
		InputAmount:  price.InputAmount,
		OutputAmount: price.OutputAmount,
		TotalGas:     price.TotalGas,
		GasPriceGwei: price.GasPriceGwei,
		GasUsd:       price.GasUsd,
		Timestamp:    price.Timestamp,
	}
}
