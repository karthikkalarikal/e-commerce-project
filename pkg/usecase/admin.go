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
	adminrepo repo.AdminRepository
}

func NewAdminUseCase(repo repo.AdminRepository) interfaces.AdminUseCase {
	return &adminUseCaseImpl{
		adminrepo: repo,
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

// user block value logic
func (usecase *adminUseCaseImpl) BlockUser(id int, block bool) (domain.Users, error) {
	fmt.Println("here")
	// var user domain.Users

	user, err := usecase.adminrepo.BlockUser(id, block)
	if err != nil {
		return domain.Users{}, err
	}
	fmt.Println("user", user)
	return user, err
}

// ------------------------------------------search user by email----------------------------- \\
func (usecase *adminUseCaseImpl) FindUserByEmail(email string) ([]domain.Users, error) {
	var user []domain.Users

	user, err := usecase.adminrepo.FindUserByEmail(email)
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
