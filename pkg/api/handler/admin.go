package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
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

// AddProduct is a function to add a new product by admin.
// @Summary Add product
// @Description Add product by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Security ApiKeyHeaderAuth
// @Success 200 {string}  domain.Product "Added product details"
// @Failure 400 {string}  response.ClientErrorResponse "Bad request"
// @Router /admin/product/addproduct [post]
func (u *AdminHandler) AddProduct(c *gin.Context) {
	var product domain.Product

	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	returnProduct, err := u.adminUseCase.AddProduct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added the product", returnProduct, nil)
	c.JSON(http.StatusOK, successRes)
}

// EditProduct is a function to edit product by admin.
// @Summary Edit product
// @Description Edit product by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Security ApiKeyHeaderAuth
// @Success 200 {string}  domain.Product "Edit product details"
// @Failure 400 {string}  response.ClientErrorResponse "Bad request"
// @Router /admin/product/editproduct [post]
func (u *AdminHandler) EditProduct(c *gin.Context) {
	var product domain.Product

	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	modProduct, err := u.adminUseCase.EditProduct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not edit the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "sucessfully edited product", modProduct, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Delete product
// @Description Delete product by id
// @Tags Product Management
// @Accept json
// @Produce json
//
//	@Param id query int true "product id"
//
// @Security ApiKeyHeaderAuth
// @Success 200 {array} domain.Product "Array of deleted product details "
// @Failure 400 {array} domain.Product  "Bad request"
// @Router /admin/product/deleteproduct [post]
func (u *AdminHandler) DeleteProduct(c *gin.Context) {

	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product, err := u.adminUseCase.DeleteProduct(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in deleting the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "sucessfully deleted the product", product, nil)
	c.JSON(http.StatusOK, succesRes)
}

// AddCategory is a function to add a new category by admin.
// @Summary Add category
// @Description Add category by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body domain.Category true "Category object"
// @Security ApiKeyHeaderAuth
// @Success 200 {string}  domain.Category "Added Category details"
// @Failure 400 {string}  response.ClientErrorResponse "Bad request"
// @Router /admin/product/addcategory [post]
func (u *AdminHandler) AddCategory(c *gin.Context) {
	var adCat domain.Category

	if err := c.BindJSON(&adCat); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	returnCategory, err := u.adminUseCase.AddCategory(adCat)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added the category", returnCategory, nil)
	c.JSON(http.StatusOK, successRes)
}
