package userController

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func (us *UserController) GetAllUsers(c echo.Context) error {
	users, err := us.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (us *UserController) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err := us.UserService.CreateUser(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"message": "success",
	})
}
