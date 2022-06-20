package models

import (
	"gorm.io/gorm"
)

type ZeroX struct {
	ChainId              int
	Pair                 string
	Price                string
	GuaranteedPrice      string
	EstimatedPriceImpact string
	Value                string
	Gas                  string
	EstimatedGas         string
	GasPrice             string
	ProtocolFee          string
	MinimumProtocolFee   string
	BuyTokenAddress      string
	SellTokenAddress     string
	BuyAmount            string
	SellAmount           string
	// Source []
	// Orders []
	SellTokenToEthRate string
	BuyTokenToEthRate  string
	Timestamp          int64
}

type ZeroXPrice struct {
	Id               uint `gorm:"primaryKey"`
	Pair             string
	BuyTokenAddress  string
	SellTokenAddress string
	BuyAmount        string
	SellAmount       string
	Timestamp        int64
}

func ZeroXPriceTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("zerox")
	}
}

func (ZeroXPrice) TableName() string {
	return "zerox"
}

func (price ZeroX) ToPrice() interface{} {
	return &ZeroXPrice{
		Pair:             price.Pair,
		BuyTokenAddress:  price.BuyTokenAddress,
		SellTokenAddress: price.SellTokenAddress,
		BuyAmount:        price.BuyAmount,
		SellAmount:       price.SellAmount,
		Timestamp:        price.Timestamp,
	}
}
