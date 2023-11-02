package handler_test

// func TestUserSignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserUsecase := mockusecase.NewMockUserUseCase(ctrl)
// 	handler := handler.NewUserHandler(mockUserUsecase)

// 	// create a test router
// 	r := gin.Default()
// 	r.POST("/users/signup", handler.UserSignUp)

// 	// define a sample user and expected response
// 	sampleUser := models.UserDetails{
// 		Name:            "Jose",
// 		Email:           "jose@example.com",
// 		Phone:           "1234567889",
// 		Password:        "password",
// 		ConfirmPassword: "password",
// 	}

// 	expectedUser := models.UserDetailsResponse{
// 		Id:      gomock.Eq(0),
// 		Address: "",
// 	}
// }
