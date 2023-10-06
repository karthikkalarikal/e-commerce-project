package usecase

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repo "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type adminUseCaseImpl struct {
	repo repo.AdminRepository
}

func NewAdminUseCase(repo repo.AdminRepository) interfaces.AdminUseCase {
	return &adminUseCaseImpl{
		repo: repo,
	}
}

// user list for admin
func (usecase *adminUseCaseImpl) UserList() ([]models.UserDetails, error) {
	userList, err := usecase.repo.UserList()
	if err != nil {
		return []models.UserDetails{}, err
	}

	return userList, nil

}

// user block value logic
func (usecase *adminUseCaseImpl) BlockUser(id int, block bool) (domain.Users, error) {
	fmt.Println("here")
	var user domain.Users

	user, err := usecase.repo.BlockUser(id, block)
	if err != nil {
		return domain.Users{}, err
	}
	fmt.Println("user", user)
	return user, err
}
