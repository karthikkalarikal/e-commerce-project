package helper

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/tealeg/xlsx"
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

func ConvertToExel(sales []models.OrderDetails) (*xlsx.File, error) {
	// Create a new Excel file
	file := xlsx.NewFile()

	// Create a new sheet
	sheet, err := file.AddSheet("Sales Report")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return nil, err
	}
	// Add headers to the sheet
	headerRow := sheet.AddRow()
	headerRow.AddCell().SetString("Item")
	// headerRow.AddCell().SetString("Quantity Sold")
	headerRow.AddCell().SetString("Total Amount")

	// Add data rows to the sheet
	fmt.Println("sales :", sales)
	for _, sale := range sales {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetString(sale.ProductName)
		// dataRow.AddCell().SetInt(sale.QuantitySold)
		dataRow.AddCell().SetFloatWithFormat(sale.TotalAmount, "#,##0.00")
		fmt.Println("date", dataRow)
	}

	buffer := new(bytes.Buffer)
	if err := file.Write(buffer); err != nil {
		fmt.Println("Error writing Excel data to buffer:", err)
		return nil, err
	}
	fmt.Println("buffer", buffer)

	// Save the Excel file
	err = file.Save("salesReport/sales_report.xlsx")
	if err != nil {
		fmt.Println("Error saving Excel file:", err)
		return nil, err
	}
	log.Print("file", file)
	return file, nil
}
