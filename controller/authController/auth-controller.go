package authController

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/middleware"
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService service.AuthService
}

// Login new user
func (a *AuthController) LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	
	userLogin, err := a.AuthService.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail Login",
			"error":   err,
		})
	}

	token, errToken := middleware.CreateToken(userLogin.ID, userLogin.Name)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	userResponse := dto.Response{Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}