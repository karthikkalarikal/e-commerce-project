package usecase

import (
	"fmt"

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

// user list for admin
func (usecase *adminUseCaseImpl) UserList() ([]models.UserDetails, error) {
	userList, err := usecase.adminrepo.UserList()
	if err != nil {
		return []models.UserDetails{}, err
	}

	return userList, nil

}

// user block value logic
func (usecase *adminUseCaseImpl) BlockUser(id int, block bool) (domain.Users, error) {
	fmt.Println("here")
	var user domain.Users

	user, err := usecase.adminrepo.BlockUser(id, block)
	if err != nil {
		return domain.Users{}, err
	}
	fmt.Println("user", user)
	return user, err
}

// search user by email
func (usecase *adminUseCaseImpl) FindUserByEmail(ctx *gin.Context) ([]domain.Users, error) {
	var user []domain.Users

	email := ctx.GetString("email")
	user, err := usecase.adminrepo.FindUserByEmail(email)
	if err != nil {
		return []domain.Users{}, err
	}
	return user, nil
}
