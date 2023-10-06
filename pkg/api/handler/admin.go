package handler

import (
	"fmt"
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
// @Router /admin/users/userlist [post]
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

type userBlock struct {
	Id      int  `json:"id"`
	Blocked bool `json:"blocked"`
}

// @Summary Block/Unblock-User
// @Description Edit block collumn of user
// @Tags User Management
// @Accept json
// @Produce json
// @Param user body userBlock true "blocked user id"
// @Security ApiKeyHeaderAuth
// @Success 200 {array} domain.Users "Array of user details "
// @Failure 400 {array} domain.Users "Bad request"
// @Router /admin/users/block [post]
func (u *AdminHandler) BlockUser(c *gin.Context) {
	// fmt.Println("here")
	var user userBlock

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	// fmt.Println("here")

	user_details, err := u.adminUseCase.BlockUser(user.Id, user.Blocked)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully blocked", user_details, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Search user by email
// @Description find user by email
// @Tags User Management
// @Accept json
// @Produce json
//
//	@Param email query string true "User's email address"
//
// @Security ApiKeyHeaderAuth
// @Success 200 {array} domain.Users "Array of user details "
// @Failure 400 {array} domain.Users "Bad request"
// @Router /admin/users/searchbyemail [post]
func (u *AdminHandler) FindUserByEmail(c *gin.Context) {
	// fmt.Println("here")

	user, err := u.adminUseCase.FindUserByEmail(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "The users", user, nil)
	c.JSON(http.StatusOK, successRes)
}

// Delete user

// @Summary Delete user
// @Description Delete user by id
// @Tags User Management
// @Accept json
// @Produce json
//
//	@Param id query int true "User's id"
//
// @Security ApiKeyHeaderAuth
// @Success 200 {string}  "Array of user details "
// @Failure 400 {string}  "Bad request"
// @Router /admin/users/searchbyemail [post]
func (u *AdminHandler) DeleteUser(c *gin.Context) {
	fmt.Println("*****Delete Handler*****")

	message, err := u.adminUseCase.DeleteUser(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, message, nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, message, nil, nil)
	c.JSON(http.StatusOK, succesRes)
}
