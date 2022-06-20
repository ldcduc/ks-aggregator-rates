package models

import (
	"gorm.io/gorm"
)

type Token struct {
	Symbol   string
	Name     string
	Decimals float64
	Address  string
}

type OneInch struct {
	Pair            string
	FromToken       Token
	ToToken         Token
	ToTokenAmount   string
	FromTokenAmount string
	EstimatedGas    float64
	Timestamp       int64
	// Protocols []
}

func IntoToken(data map[string]interface{}) Token {
	return Token{
		Symbol:   data["symbol"].(string),
		Name:     data["name"].(string),
		Decimals: data["decimals"].(float64),
		Address:  data["address"].(string),
	}
}

type OneInchPrice struct {
	Id               uint `gorm:"primaryKey"`
	Pair             string
	FromTokenAddress string
	ToTokenAddress   string
	FromTokenAmount  string
	ToTokenAmount    string
	EstimatedGas     float64
	Timestamp        int64
}

func OneInchPriceTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("oneinch")
	}
}

func (OneInchPrice) TableName() string {
	return "oneinch"
}

func (price OneInch) ToPrice() interface{} {
	return &OneInchPrice{
		Pair:             price.Pair,
		FromTokenAddress: price.FromToken.Address,
		ToTokenAddress:   price.ToToken.Address,
		FromTokenAmount:  price.FromTokenAmount,
		ToTokenAmount:    price.ToTokenAmount,
		EstimatedGas:     price.EstimatedGas,
		Timestamp:        price.Timestamp,
	}
}
