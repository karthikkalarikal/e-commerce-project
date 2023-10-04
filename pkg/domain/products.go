package domain

type Product struct {
	Id            int    `json:"id" gorm:"unique;not null"`
	Category_id   int    `json:"category_id"`
	ProductName   string `json:"product_name"`
	Product_image []byte `json:"product_image"`
	Colour        string `json:"colour"`
	Stock         string `json:"stock"`
	Price         string `json:"price"`
}
