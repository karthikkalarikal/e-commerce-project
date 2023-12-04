package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminRepository interface {
	UserList(int, int) ([]models.UserDetailsResponse, error)
	BlockUser(id int, block bool) (domain.Users, error)
	FindUser(email string, name string, id string, pageList int, offset int) ([]domain.Users, error)
	DeleteUser(id int) (domain.Users, error)

	CountUsers() (int, error)
	SumRevenueByMonth() (float64, error)
	GetSalesDetailsByYear(year int) ([]models.OrderDetails, error)
	GetSalesDetailsByMonth(month int) ([]models.OrderDetails, error)
	GetSalesDetailsByDay(day int) ([]models.OrderDetails, error)
	GetSalesReport(year int) (models.OrderDetails, error)

	GetImgagesById(productId int) ([]string, error)
}
