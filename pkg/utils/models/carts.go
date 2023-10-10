package models

type Cart struct {
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
