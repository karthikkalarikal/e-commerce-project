package models

type CombinedOrderDetails struct {
	OrderId       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	OrderStatus   string  `json:"order_status"`
	PaymentStatus bool    `json:"payment_status"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
	Street        string  `json:"street"`
	City          string  `json:"city"`
}

type ItemDetails struct {
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}

// to watch our for if you dont give the name similar to the gorm/json name there might be some errors

type OrderDetails struct {
	TotalAmount float64 `gorm:"column:sum"`
	ProductName string  `gorm:"column:product_name"`
}
