package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/karthikkalarikal/ecommerce-project/pkg/helper"
	"github.com/karthikkalarikal/ecommerce-project/pkg/mock/mockrepo"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type LoginKey string

const (
	SignUp LoginKey = "Signup"
)

func createSignUp(loginKey LoginKey) models.UserDetails {
	if loginKey == "Signup" {
		return models.UserDetails{
			Name:     helper.GenerateRandomString(7),
			Email:    helper.GenerateRandomString(12),
			Phone:    helper.GenerateRandomString(10),
			Password: helper.GenerateRandomString(12),
		}
	} else {
		return models.UserDetails{}
	}

}

func TestUserSignUp(t *testing.T) {
	tests := []struct {
		testName       string
		input          models.UserDetails
		expectedOutput models.TokenUsers
		buildStub      func(mockRepo *mockrepo.MockUserRepository, loginDetails models.UserDetails)
		expectedError  error
	}{
		{
			testName: "Test Sign Up Password is Different",
			input:    models.UserDetails{Password: "password", ConfirmPassword: "password2"},
			buildStub: func(mockRepo *mockrepo.MockUserRepository, loginDetails models.UserDetails) {
				// no mock is used

			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("password does not match"),
		},
		{
			testName: "Test Sign Up Email already exists",
			input:    createSignUp(SignUp),
			buildStub: func(mockRepo *mockrepo.MockUserRepository, loginDetails models.UserDetails) {
				// dbError := fmt.Errorf("user already exist, sign in")
				mockRepo.EXPECT().
					CheckUserAvailability(loginDetails.Email).
					Times(1).
					Return(false)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("user already exist, sign in"),
		},
		// {
		// 	testName: "Test Sigh Up Success",
		// 	input:    createSignUp(SignUp),
		// 	buildStub: func(mockRepo *mockrepo.MockUserRepository, loginDetails models.UserDetails) {
		// 		hashedPassword,err := helper.GenerateTokenClients(loginDetails)

		// 	},
		// },
	}

	for _, test := range tests {
		// var test = test
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mockrepo.NewMockUserRepository(ctrl)
			mockHelpRepo := mockrepo.NewMockHelperRepository(ctrl)

			test.buildStub(mockUserRepo, test.input)

			userUsecase := usecase.NewUserUseCase(mockUserRepo, mockHelpRepo)
			output, err := userUsecase.UserSignUp(test.input)
			fmt.Println("test input", test.input.ConfirmPassword)
			fmt.Println("err", err)
			// test.buildStub(mockRepo, test.input)

			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedOutput, output)
		})

	}
}

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockrepo.NewMockUserRepository(ctrl)
	helperRepo := mockrepo.NewMockHelperRepository(ctrl)

	userUseCase := usecase.NewUserUseCase(userRepo, helperRepo)
	sampleUser := models.UserDetailsResponse{
		Id:    1,
		Name:  "sample",
		Email: "sample@gmail.com",
		Phone: "0000000000",
	}
	sampleAdmin := models.AdminDetailsResponse{
		Id:    1,
		Name:  "admin",
		Email: "admin@gmail.com",
		Phone: "0000000000",
		Role:  true,
	}

	tokenStringUser, _ := helper.GenerateTokenClients(sampleUser)
	tokenStringAdmin, _ := helper.GenerateTokenAdmin(sampleAdmin)

	testData := map[string]struct {
		input          models.UserLogin
		stub           func(mockrepo.MockUserRepository, mockrepo.MockHelperRepository, models.UserLogin)
		expectedOutput interface{}
		expectedError  error
		expectedBool   bool
	}{
		"Test User Login Success": {
			input: models.UserLogin{
				Email:    "sample@gmail.com",
				Password: "password",
			},
			stub: func(mr mockrepo.MockUserRepository, mh mockrepo.MockHelperRepository, u models.UserLogin) {
				userRepo.EXPECT().CheckUserAvailability(u.Email).Times(1).Return(true)
				userRepo.EXPECT().UserBlockedStatus(u.Email).Times(1).Return(false, nil)
				userRepo.EXPECT().FindUserByEmail(u.Email).Times(1).Return(models.UserSignInResponse{
					Id:       1,
					UserID:   1,
					Name:     "sample",
					Email:    "sample@gmail.com",
					Phone:    "0000000000",
					Password: "password",
					Role:     false,
				}, nil)

			},

			expectedOutput: models.TokenUsers{
				Users: sampleUser,
				Token: tokenStringUser,
			},
			expectedError: nil,
			expectedBool:  false,
		},
		"Test User Login Fail": {
			input: models.UserLogin{
				Email:    "sample@gmail.com",
				Password: "password",
			},
			stub: func(mur mockrepo.MockUserRepository, mhr mockrepo.MockHelperRepository, ul models.UserLogin) {
				userRepo.EXPECT().CheckUserAvailability(ul.Email).Times(1).Return(false)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("the user does not exist"),
			expectedBool:   false,
		},

		"Test Admin Login Success": {
			input: models.UserLogin{
				Email:    "admin@gmail.com",
				Password: "password",
			},
			stub: func(mr mockrepo.MockUserRepository, mh mockrepo.MockHelperRepository, u models.UserLogin) {
				userRepo.EXPECT().CheckUserAvailability(u.Email).Times(1).Return(true)
				userRepo.EXPECT().UserBlockedStatus(u.Email).Times(1).Return(false, nil)
				userRepo.EXPECT().FindUserByEmail(u.Email).Times(1).Return(models.UserSignInResponse{
					Id:       1,
					UserID:   1,
					Name:     "admin",
					Email:    "admin@gmail.com",
					Phone:    "0000000000",
					Password: "password",
					Role:     true,
				}, nil)

			},

			expectedOutput: models.TokenAdmin{
				Users: sampleAdmin,
				Token: tokenStringAdmin,
			},
			expectedError: nil,
			expectedBool:  true,
		},
	}

	for _, test := range testData {
		test.stub(*userRepo, *helperRepo, test.input)

		tokenUsers, err, role := userUseCase.LoginHandler(test.input)

		assert.Equal(t, test.expectedOutput, tokenUsers)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedBool, role)
	}

}

func TestAddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockrepo.NewMockUserRepository(ctrl)

	userUseCase := usecase.NewUserUseCase(userRepo, nil)

	sampleAddress := models.Address{

		Name:      "sample",
		HouseName: "sample",
		Street:    "sample",
		City:      "sample",
		State:     "sample",
		Pin:       "sample",
	}

	testData := map[string]struct {
		input          models.Address
		stub           func(mockrepo.MockUserRepository, models.Address)
		expectedOutput []models.Address
		expectedError  error
	}{
		"success": {
			input: sampleAddress,
			stub: func(m mockrepo.MockUserRepository, a models.Address) {
				m.EXPECT().AddAddress(sampleAddress, 1).Times(1).Return(nil)
				m.EXPECT().FindAddress(1).Times(1).Return([]models.Address{sampleAddress}, nil)
			},
			expectedOutput: []models.Address{sampleAddress},
			expectedError:  nil,
		},
	}
	for _, test := range testData {

		test.stub(*userRepo, test.input)

		addresses, err := userUseCase.AddAddress(test.input, 1)

		assert.Equal(t, test.expectedOutput, addresses)
		assert.Equal(t, test.expectedError, err)
	}
}
