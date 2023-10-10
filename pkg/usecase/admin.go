package usecase

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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

// delete user
func (usecase *adminUseCaseImpl) DeleteUser(ctx *gin.Context) (string, error) {
	// fmt.Println("****delete usecase ******")
	id_str := ctx.Query("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return "error in atoi", err
	}
	_, err = usecase.adminrepo.DeleteUser(id)
	if err != nil {
		return "could not delete", err
	}

	return "succesfuly deleted", nil

}

// add products
func (usecase *adminUseCaseImpl) AddProduct(products domain.Product) (domain.Product, error) {
	product, err := usecase.adminrepo.AddProduct(products)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil

}

// edit products
func (usecase *adminUseCaseImpl) EditProduct(product domain.Product) (domain.Product, error) {
	modProduct, err := usecase.adminrepo.EditProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	return modProduct, nil
}

// delete products
func (usecase *adminUseCaseImpl) DeleteProduct(id int) (domain.Product, error) {

	delProduct, err := usecase.adminrepo.DeleteProduct(id)
	if err != nil {
		return domain.Product{}, nil
	}

	return delProduct, nil
}

// add category
func (usecase *adminUseCaseImpl) AddCategory(category domain.Category) (domain.Category, error) {
	adCat, err := usecase.adminrepo.AddCategory(category)
	if err != nil {
		return adCat, err
	}
	return adCat, nil

}
