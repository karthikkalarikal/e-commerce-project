package models

type AdminDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  bool   `json:"role"`
}

type TokenAdmin struct {
	Users AdminDetailsResponse
	Token string
}
