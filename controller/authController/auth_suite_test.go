package authController

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service/mock"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteAuth struct {
	suite.Suite
	authController *AuthController
	mock           *mock.AuthMock
}

func (s *suiteAuth) SetupTest() {
	mock := &mock.AuthMock{}
	s.mock = mock
	s.authController = &AuthController{
		AuthService: s.mock,
	}
}

func (s *suiteAuth) TestLoginUsers() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               model.User
		HasReturnBody      bool
		ExpectedBody       model.User
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			},
			false,
			model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			},
		},
		{
			"failLogin",
			http.StatusInternalServerError,
			"GET",
			model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			},
			false,
			model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			},
		},
		// {
		// 	"failCreate",
		// 	http.StatusInternalServerError,
		// 	"POST",
		// 	model.User{
		// 		Name: "rafli",
		// 		// Email:    "rafli@gmail.com",
		// 		Password: "123456",
		// 	},
		// 	false,
		// 	"Failed",
		// },
	}

	for _, v := range testCase {
		var mockAuth model.User
		switch v.Name {
		case "success":
			mockAuth = model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			}
		case "failLogin":
			mockAuth = model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			}
		}
		var mockCall = s.mock.On("LoginUser", mockAuth)
		switch v.Name {
		case "success":
			mockCall.Return(model.User{
				Model: &gorm.Model{
					ID: 1,
				},
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			}, nil)
		case "failLogin":
			mockCall.Return(model.User{}, errors.New("error"))
		}
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/login", bytes.NewBuffer(res))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			err := s.authController.LoginUserController(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(t, v.ExpectedBody.Name, resp["user"].(map[string]interface{})["name"])
				s.Equal(t, v.ExpectedBody.Email, resp["user"].(map[string]interface{})["email"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteAuth(t *testing.T) {
	suite.Run(t, new(suiteAuth))
}
