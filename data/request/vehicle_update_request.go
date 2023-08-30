package request

import "github.com/shopspring/decimal"

type VehicleUpdateRequest struct {
	Id int
	Brand string  `validade:"required min=1,max=100" json:"brand"`
	Model string  `validade:"required min=1,max=100" json:"model"`
	Price decimal.Decimal `json:"price"`
	Year  int  `validade:"required min=4" json:"year"`
}