package userController

import (
	"belajar-go-echo/dto"
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
)

type suiteUsers struct {
	suite.Suite
	userController *UserController
	mock           *mock.UserMock
}

func (s *suiteUsers) SetupSuite() {
	mock := &mock.UserMock{}
	s.mock = mock
	s.userController = &UserController{
		UserService: s.mock,
	}

}

func (s *suiteUsers) TestGetAllUsers() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// Body               model.User
		HasReturnBody bool
		ExpectedBody  dto.UserDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			// model.User{
			// 	Name: "bimo",
			// },
			true,
			dto.UserDTO{
				Name:  "rafli",
				Email: "rafli@gmail.com",
			},
		},
		{
			"failGetAll",
			http.StatusInternalServerError,
			"GET",
			// model.Users{},
			false,
			dto.UserDTO{},
		},
	}

	for _, v := range testCase {
		var mockCall = s.mock.On("GetAllUsers")
		switch v.Name {
		case "success":
			mockCall.Return([]dto.UserDTO{
				{
					Name:  "rafli",
					Email: "rafli@gmail.com",
				},
			}, nil)
		case "failGetAll":
			mockCall.Return([]dto.UserDTO{
				{},
			}, errors.New("Failed"))
		}
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.userController.GetAllUsers(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody.Name, resp["data"].([]interface{})[0].(map[string]interface{})["name"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteUsers) TestCreateUsers() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               model.User
		HasReturnBody      bool
		ExpectedMesaage    string
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			model.User{
				Name:     "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			},
			true,
			"success",
		},
		{
			"failCreate",
			http.StatusInternalServerError,
			"POST",
			model.User{
				Name: "rafli",
				// Email:    "rafli@gmail.com",
				Password: "123456",
			},
			false,
			"Failed",
		},
	}

	for _, v := range testCase {
		var mockUser model.User
		switch v.Name {
		case "success":
			mockUser = model.User{
				Name: "rafli",
				Email:    "rafli@gmail.com",
				Password: "123456",
			}
		case "failCreate":
			mockUser = model.User{
				Name: "rafli",
				// Email:    "rafli@gmail.com",
				Password: "123456",
			}
		}
		var mockCall = s.mock.On("CreateUser", mockUser)
		switch v.Name {
		case "success":
			mockCall.Return(nil)
		case "failCreate":
			mockCall.Return(errors.New("Failed"))
		}
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/users", bytes.NewBuffer(res))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.userController.CreateUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedMesaage, resp["message"].(string))
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
