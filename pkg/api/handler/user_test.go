package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/helper"
	"github.com/karthikkalarikal/ecommerce-project/pkg/mock/mockusecase"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserSignUp(t *testing.T) {

	// define a sample user and expected response
	sampleUser := models.UserDetails{
		Name:            "Jose",
		Email:           "jose@example.com",
		Phone:           "1234567889",
		Password:        "password",
		ConfirmPassword: "password",
	}

	tokenUserResponse := models.UserDetailsResponse{
		Id:    1,
		Name:  "Jose",
		Email: "jose@example.com",
		Phone: "1234567889",
	}
	tokenUser, _ := helper.GenerateTokenClients(tokenUserResponse)

	testCase := map[string]struct {
		input         models.UserDetails
		stub          func(useCaseMock *mockusecase.MockUserUseCase, signup interface{})
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"successful": {
			input: sampleUser,
			stub: func(useCaseMock *mockusecase.MockUserUseCase, signup interface{}) {

				_ = validator.New().Struct(signup)
				useCaseMock.EXPECT().UserSignUp(sampleUser).Times(1).Return(models.TokenUsers{
					Users: tokenUserResponse,
					Token: tokenUser,
				})

			},
			checkResponse: func(t *testing.T, rsp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, rsp.Code)
			},
		},
	}
	for testName, test := range testCase {
		test := test

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserUsecase := mockusecase.NewMockUserUseCase(ctrl)
			handler := handler.NewUserHandler(mockUserUsecase)

			// create a test router
			r := gin.New()
			url := "/users/signup"
			r.POST(url, handler.UserSignUp)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, url, body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			r.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)
		})
	}
}
