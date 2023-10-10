package domain

type Product struct {
	ProductId     int      `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryId    int      `json:"category_id" gorm:"foreignkey:CategoryId"`
	Category      Category `json:"-" gorm:"foreignkey:CategoryId;constraint:OnDelete:CASCADE"`
	ProductName   string   `json:"product_name"`
	Product_image []byte   `json:"product_image"`
	Colour        string   `json:"colour"`
	Stock         string   `json:"stock"`
	Price         string   `json:"price"`
}

type Category struct {
	CategoryId   int    `json:"id" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name"`
	// SubCategory_id int    `json:"subcategory_id"`
}
