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

// @title Go + Gin E-Commerce API
// @version 1.0.0
// @description Stylezine is an E-commerce platform to purchase and sell Electronic itmes
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

// GetUserList godoc
// @Summary List the users you could specify page and no of users in one page
// @Description Retrive and display user list according to instructions
// @Tags User Management
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param per_page query int false "Results per page (default 10)"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "Array of user details "
// @Failure 400 {array} response.Response "Bad request"
// @Router /admin/users/userlist [get]
func (u *AdminHandler) UserList(c *gin.Context) {

	pageNo := c.DefaultQuery("page", "1")        // default 1
	pageList := c.DefaultQuery("per_page", "10") // default to 10

	pageNoInt, _ := strconv.Atoi(pageNo)
	pageListInt, _ := strconv.Atoi(pageList)

	user_list, err := u.adminUseCase.UserList(pageNoInt, pageListInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Users cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	message := "Userlist"

	successRes := response.ClientResponse(http.StatusOK, message, user_list, nil)
	c.JSON(http.StatusOK, successRes)
}

// Block/Unblock godoc
// @Summary Block/Unblock User
// @Description Block/Unblock on prompt
// @Tags User Management
// @Produce json
// @Param user_id path int true "user id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "The user details"
// @Failure 400 {object} response.Response "Bad request"
// @Router /admin/users/block/{user_id} [patch]
func (u *AdminHandler) BlockUser(c *gin.Context) {
	// fmt.Println("here")

	id := c.Param("user_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	// fmt.Println("here")

	user_details, err := u.adminUseCase.BlockUser(idInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully blocked/unblocked", user_details, nil)
	c.JSON(http.StatusOK, successRes)
}

// FindUsers godoc
// @Summary Search user by various criteria
// @Description Search for users based on various criteria with pagination.
// @Tags User Management
// @Accept json
// @Produce json
// @Param name query string false "Name to search for"
// @Param email query string false "Email address to search for"
// @Param id query int false "ID to search for"
// @Param page query int false "Page number (default 1)"
// @Param per_page query int false "Result per page (default 1)"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "Array of user details "
// @Failure 400 {array} response.Response "Bad request"
// @Failure 500 {array} response.Response "Error in server"
// @Router /admin/users/searchuser [get]
func (u *AdminHandler) FindUser(c *gin.Context) {

	email := c.Query("email")
	name := c.Query("name")
	id := c.Query("id")
	pageNo := c.DefaultQuery("page", "1")
	pageList := c.DefaultQuery("per_page", "10")
	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil {
		pageNoInt = 1
	}
	pageListInt, err := strconv.Atoi(pageList)
	if err != nil {
		pageListInt = 10
	}

	user, err := u.adminUseCase.FindUser(email, name, id, pageNoInt, pageListInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "The users", user, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete user by id
// @Tags User Management
// @Accept json
// @Produce json
// @Param user_id query int true "user id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Array of user details "
// @Failure 400 {object} response.Response "Bad request"
// @Router /admin/users/deleteuser/{user_id} [post]
func (u *AdminHandler) DeleteUser(c *gin.Context) {
	fmt.Println("*****Delete Handler*****")

	id := c.Query("user_id")
	fmt.Println("id:", id)
	id_int, err := strconv.Atoi(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in the user_id praram", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	user, err := u.adminUseCase.DeleteUser(id_int)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not delete the user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully deleted the user", user, nil)
	c.JSON(http.StatusOK, succesRes)
}

// AddProduct is a function to add a new product by admin.
// @Summary Add product
// @Description Add product by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Security BearerTokenAuth
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
// @Security BearerTokenAuth
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
// @Security BearerTokenAuth
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
// @Security BearerTokenAuth
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
