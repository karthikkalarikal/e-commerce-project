package handler

import (
	"net/http"
	"strconv"

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

// @Summary Address
// @Description Enter and save userAdress along with userId
// @Tags Address Management
// @Accept json
// @Produce json
// @Param user_id query int true "User Id"
// @Param user body models.Address true "User details"
// @Security BearerTokenAuth
// @Success 201 {array} models.UserDetails "User details and token"
// @Failure 400 {array} models.UserSignInResponse{} "Bad request"
// @Router /users/user/address [post]
func (u *UserHandler) UserAddress(c *gin.Context) {
	var address models.Address
	userId := c.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//bind user details to struct
	if err := c.BindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//checking validity

	err = validator.New().Struct(address)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	//buisness logic
	// fmt.Println(user)
	userAddress, err := u.userUseCase.AddAddress(address, userIdInt)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "address added succesfully ", nil, err.Error())
		c.JSON(http.StatusCreated, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "address added succesfully", userAddress, nil)
	// fmt.Println(userCreated)

	c.JSON(http.StatusCreated, successRes)
}

// @Summary SelectAddress
// @Description Address selected for cash on delivery
// @Tags Address Management
// @Produce json
// @Param address_id query int true "Address Id"
// @Param selection query bool true "selection "
// @Security BearerTokenAuth
// @Success 201 {object} response.Response "changed addres"
// @Failure 400 {object} response.Response "Bad request"
// @Router /users/user/select [patch]
func (u *UserHandler) SelectAddress(c *gin.Context) {
	addressId := c.Query("address_id")
	addressIdInt, err := strconv.Atoi(addressId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	val := c.Query("selection")
	valBol, err := strconv.ParseBool(val)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userAddress, err := u.userUseCase.SelectAddress(addressIdInt, valBol)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "address added succesfully", userAddress, nil)
	// fmt.Println(userCreated)

	c.JSON(http.StatusCreated, successRes)

}

// @Summary ViewUser
// @Description UserDetails
// @Tags User Profile
// @Produce json
// @Param user_id query int true "User Id"
// @Security BearerTokenAuth
// @Success 201 {object} models.UserDetails "changed addres"
// @Failure 400 {object} models.UserSignInResponse{} "Bad request"
// @Router /users/user/viewdetails [get]
func (u *UserHandler) ViewUser(c *gin.Context) {
	userId := c.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	user, err := u.userUseCase.FindUserById(userIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "user details fetched succesfully", user, nil)
	// fmt.Println(userCreated)

	c.JSON(http.StatusCreated, successRes)

}

// @Summary ViewAddress
// @Description get Address by user id
// @Tags User Profile
// @Produce json
// @Param user_id query int true "User Id"
// @Security BearerTokenAuth
// @Success 201 {object} response.Response "changed addres"
// @Failure 400 {object} response.Response "Bad request"
// @Router /users/user/addresses [get]
func (u *UserHandler) GetAddress(c *gin.Context) {
	userId := c.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	user, err := u.userUseCase.FindAddressByUI(userIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "addresses fetched succesfully", user, nil)
	// fmt.Println(userCreated)

	c.JSON(http.StatusCreated, successRes)
}
