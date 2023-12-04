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
	}

	for _, test := range tests {
		// var test = test
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockrepo.NewMockUserRepository(ctrl)
			test.buildStub(mockRepo, test.input)

			userUsecase := usecase.NewUserUseCase(mockRepo, nil)
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
