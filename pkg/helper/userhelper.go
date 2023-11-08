package helper

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
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

// for exel conversion

func ConvertToExel(sales []models.OrderDetails) (*excelize.File, error) {
	// Create a new Excel file
	filename := "salesReport/sales_report.xlsx"
	file := excelize.NewFile()

	// create headers for excel
	file.SetCellValue("Sheet1", "A1", "Item")
	file.SetCellValue("Sheet1", "B1", "Total Amount Sold")

	// Add data rows to the sheet
	// fmt.Println("sales :", sales)
	for i, sale := range sales {
		col1 := fmt.Sprintf("A%d", i+1)
		col2 := fmt.Sprintf("B%d", i+1)
		// dataRow := sheet.AddRow()
		file.SetCellValue("Sheet1", col1, sale.ProductName)
		file.SetCellValue("Sheet1", col2, sale.TotalAmount)

	}

	// save excel
	if err := file.SaveAs(filename); err != nil {
		return nil, err
	}
	// fmt.Println(file)

	return file, nil
}
