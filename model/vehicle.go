package model

import "github.com/shopspring/decimal"

type Vehicle struct {
	Id int 
	Brand string
	Model string
	Price decimal.Decimal
	Year int
} 