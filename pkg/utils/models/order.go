package models

type CombinedOrderDetails struct {
	OrderId        string  `json:"order_id"`
	Amount         float64 `json:"amount"`
	ShipmentStatus string  `json:"order_status"`
	PaymentStatus  bool    `json:"payment_status"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
}
