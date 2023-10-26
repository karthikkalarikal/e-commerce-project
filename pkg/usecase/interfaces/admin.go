package interfaces

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminUseCase interface {
	UserList(int, int) ([]models.UserDetailsResponse, error)
	BlockUser(id int) (domain.Users, error)
	FindUser(email string, name string, id string, pageNo int, pageList int) ([]domain.Users, error)
	DeleteUser(id int) (domain.Users, error)
	TotalSalesByMonth() (float64, error)
	GetSalesDetailsByDate(yearInt, monthInt, dayInt int) ([]models.OrderDetails, error)
	PrintSalesReport(sales []models.OrderDetails) (*gofpdf.Fpdf, error)
}
