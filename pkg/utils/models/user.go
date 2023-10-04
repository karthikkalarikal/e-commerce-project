package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Role            bool   `json:"role" default:"false"`
}

type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

// to be shown after logging in
type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  bool   `json:"role" default:"false"`
}

// token and details , token to access protected routes
type TokenUsers struct {
	Users UserDetailsResponse
	Token string
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     bool   `json:"role" default:"false"`
}
