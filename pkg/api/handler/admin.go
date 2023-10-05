package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type AdminHandler struct {
	adminUseCase interfaces.AdminUseCase
}

func NewAdminHandler(usecase interfaces.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary ListProducts
// @Description Retrive and display user list
// @Tags User Management
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {array} models.UserDetails "Array of user details "
// @Failure 400 {array} models.UserDetails "Bad request"
// @Router /admin/userlist [post]
func (u *AdminHandler) UserList(c *gin.Context) {
	product_list, err := u.adminUseCase.UserList()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Users cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	message := "Userlist"

	successRes := response.ClientResponse(http.StatusOK, message, product_list, nil)
	// fmt.Println(product_list)
	c.JSON(http.StatusOK, successRes)
}
