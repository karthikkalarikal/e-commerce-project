package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCase
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary UserSignUP
// @Description Retrive UserDetails stored in DB and a auth token with success message
// @Tags User Auth
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User details"
// @Success 201 {array} models.UserDetails "User details and token"
// @Failure 400 {array} models.UserSignInResponse{} "Bad request"
// @Router /users/signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {
	var user models.UserDetails

	//bind user details to struct
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//checking validity

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	//buisness logic
	// fmt.Println(user)
	userCreated, err := u.userUseCase.UserSignUp(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User couldnt sign up", nil, err.Error())
		c.JSON(http.StatusCreated, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "user succesfully signed up", userCreated, nil)
	// fmt.Println(userCreated)

	c.JSON(http.StatusCreated, successRes)
}

// signin handler
// @Summary UserSignIN
// @Description Sign in a user and return user details and a token
// @Tags User Auth
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User login details"
// @Success 200 {object} models.UserSignInResponse "User details and role"
// @Failure 400 {array} models.UserSignInResponse "Bad request"
// @Router /users/login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {
	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err, role := u.userUseCase.LoginHandler(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var message string
	if role {
		message = "admin succesfully logged in"
	} else {
		message = "User succesfully logged in"
	}
	successRes := response.ClientResponse(http.StatusOK, message, user_details, nil)
	c.JSON(http.StatusOK, successRes)
}
