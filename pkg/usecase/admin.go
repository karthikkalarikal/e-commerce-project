package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repo "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type adminUseCaseImpl struct {
	adminrepo  repo.AdminRepository
	helperrepo repo.HelperRepository
}

func NewAdminUseCase(adrepo repo.AdminRepository, helper repo.HelperRepository) interfaces.AdminUseCase {
	return &adminUseCaseImpl{
		adminrepo:  adrepo,
		helperrepo: helper,
	}
}

// --------------------------------user list for admin--------------------------- \\
func (usecase *adminUseCaseImpl) UserList(pageNo int, pageList int) ([]models.UserDetailsResponse, error) {

	total, _ := usecase.adminrepo.CountUsers()
	fmt.Println("total", total)

	if pageNo <= 0 {
		pageNo = 1
	}
	offset := (pageNo - 1) * pageList

	userList, err := usecase.adminrepo.UserList(pageList, offset)
	if err != nil {
		return []models.UserDetailsResponse{}, err
	}

	return userList, nil

}

// ---------------------------------------block user----------------------------------------------\\
func (usecase *adminUseCaseImpl) BlockUser(id int) (domain.Users, error) {
	// fmt.Println("here")
	// var user domain.Users
	// blocked status
	users, err := usecase.helperrepo.GetUserDetailsThroughId(id)
	if err != nil {
		return domain.Users{}, err
	}

	user, err := usecase.adminrepo.BlockUser(id, !users.Blocked)
	if err != nil {
		return domain.Users{}, err
	}
	fmt.Println("user", user)
	return user, err
}

// ------------------------------------------search user ------------------------------------------ \\
func (usecase *adminUseCaseImpl) FindUser(email string, name string, id string, pageNo int, pageList int) ([]domain.Users, error) {
	var user []domain.Users

	offset := (pageNo - 1) * pageList
	fmt.Println("page", pageList, "offset", offset)
	user, err := usecase.adminrepo.FindUser(email, name, id, pageList, offset)
	if err != nil {
		return []domain.Users{}, err
	}
	return user, nil
}

// ------------------------------------------delete user ------------------------------------------------- \\
func (usecase *adminUseCaseImpl) DeleteUser(id int) (domain.Users, error) {
	// fmt.Println("****delete usecase ******")

	user, err := usecase.adminrepo.DeleteUser(id)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil

}

// ------------------------------------------ total sales by month ---------------------------------------- \\

func (usecase *adminUseCaseImpl) TotalSalesByMonth() (float64, error) {
	amount, err := usecase.adminrepo.SumRevenueByMonth()
	if err != nil {
		return 0, err
	}

	return amount, nil
}

// ----------------------------------------- sales by products by date --------------------------------------- \\

func (usecase *adminUseCaseImpl) GetSalesDetailsByDate(yearInt, monthInt, dayInt int) ([]models.OrderDetails, error) {

	// by year
	if yearInt > 0 {
		body, err := usecase.adminrepo.GetSalesDetailsByYear(yearInt)
		if err != nil {
			return []models.OrderDetails{}, err
		}
		return body, nil
	}

	// by month
	if monthInt > 0 {
		body, err := usecase.adminrepo.GetSalesDetailsByMonth(monthInt)
		if err != nil {
			return []models.OrderDetails{}, err
		}
		return body, nil
	}

	// by day
	if dayInt > 0 {
		body, err := usecase.adminrepo.GetSalesDetailsByDay(dayInt)
		if err != nil {
			return []models.OrderDetails{}, err
		}

		return body, nil
	}

	return []models.OrderDetails{}, errors.New("no value detected")
}

// ----------------------- print sales report ---------------------------------- \\

func (repo *adminUseCaseImpl) PrintSalesReport(sales []models.OrderDetails) (*gofpdf.Fpdf, error) {

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font and title
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Total Sales Report")
	pdf.Ln(10)

	// Add items to the PDF
	for _, item := range sales {
		pdf.Cell(0, 10, "Item: "+item.ProductName)
		pdf.Ln(10)
		amount := strconv.FormatFloat(item.TotalAmount, 'f', -1, 64)
		pdf.Cell(0, 10, "Sold: $ "+amount)
		pdf.Ln(10)

	}
	pdf.Ln(10)

	// Add the total amount to the PDF

	return pdf, nil
}
