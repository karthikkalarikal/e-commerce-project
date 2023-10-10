package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type HelperRepository interface {
	GetUserDetailsThroughId(id int) (domain.Users, error)
}
