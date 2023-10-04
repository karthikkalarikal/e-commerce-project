package usecase

import (
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

func (usecase *adminUseCaseImpl) UserList() ([]models.UserDetails, error) {
	userList, err := usecase.repo.UserList()
	if err != nil {
		return []models.UserDetails{}, err
	}

	return userList, nil

}
