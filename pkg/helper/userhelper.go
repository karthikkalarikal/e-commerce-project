package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// generate token for user

func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 48))
	issuedAt := jwt.NewNumericDate(time.Now())

	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// token for admin
func GenerateTokenAdmin(user models.AdminDetailsResponse) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 48))
	issuedAt := jwt.NewNumericDate(time.Now())

	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
