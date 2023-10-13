package models

type Cart struct {
	CartId     int     `json:"cart_id"`
	UserId     int     `json:"user_id"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CartItems struct {
	CartId    int    `json:"cart_id"`
	ProductId int    `json:"product_id"`
	Quantity  string `json:"quantity"`
	Amount    string `json:"amount"`
}
