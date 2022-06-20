package models

import (
	"gorm.io/gorm"
)

type ParaSwap struct {
	BlockNumber  int
	Network      int
	Pair         string
	SrcToken     string
	SrcDecimals  float64
	SrcAmount    string
	SrcUSD       string
	DestToken    string
	DestDecimals float64
	DestAmount   string
	DestUSD      string
	Side         string
	Timestamp    int64
}

type ParaSwapPrice struct {
	Id          uint `gorm:"PrimaryKey"`
	BlockNumber int
	Pair        string
	SrcToken    string
	SrcAmount   string
	DestToken   string
	DestAmount  string
	Timestamp   int64
}

func ParaSwapPriceTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("paraswap")
	}
}

func (ParaSwapPrice) TableName() string {
	return "paraswap"
}

func (price ParaSwap) ToPrice() interface{} {
	return &ParaSwapPrice{
		BlockNumber: price.BlockNumber,
		Pair:        price.Pair,
		SrcToken:    price.SrcToken,
		SrcAmount:   price.SrcAmount,
		DestToken:   price.DestToken,
		DestAmount:  price.DestAmount,
		Timestamp:   price.Timestamp,
	}
}
