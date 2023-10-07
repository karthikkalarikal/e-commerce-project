package domain

type Product struct {
	ProductID     int      `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryID    int      `json:"category_id" gorm:"foreignkey:CategoryID"`
	Category      Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName   string   `json:"product_name"`
	Product_image []byte   `json:"product_image"`
	Colour        string   `json:"colour"`
	Stock         string   `json:"stock"`
	Price         string   `json:"price"`
}

type Category struct {
	CategoryID   int    `json:"id" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name"`
	// SubCategory_id int    `json:"subcategory_id"`
}
