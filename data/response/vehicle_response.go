package response

import "github.com/shopspring/decimal"


type VehicleResponse struct {
	Id    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price decimal.Decimal `json:"price"`
	Year int `json:"year"`
}