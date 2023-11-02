package domain

type Users struct {
	UserID   int    `json:"id" gorm:"primarykey;autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" validate:"min=8,max=20" gorm:"not null"`
	Phone    string `json:"phone" gorm:"not null"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
	Role     bool   `json:"role" gorm:"default:false"`
}

type Address struct {
	AddressId int    `json:"address_id" gorm:";primarykey;autoIncrement"`
	Selection bool   `json:"selection" gorm:"default:false"`
	UserID    int    `json:"user_id"`
	Users     Users  `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}
