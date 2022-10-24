package routes

import (
	"belajar-go-echo/constants"
	"belajar-go-echo/controller/authController"
	"belajar-go-echo/controller/userController"
	"belajar-go-echo/repository"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	userRepository := repository.NewUserGormSql(db)
	userService := service.NewUserService(userRepository)
	// userController := controller.NewUserController(userService)
	userController := userController.UserController{
		UserService: userService,
	}
	
	authService := service.NewAuthService(userRepository)
	authController := authController.AuthController{
		AuthService: authService,
	}

	app := echo.New()

	app.GET("/login", authController.LoginUserController)

	app.GET("/users", userController.GetAllUsers, middleware.JWT([]byte(constants.SECRET_JWT)))
	app.POST("/users", userController.CreateUser)

	app.GET("/hallo", func(c echo.Context) error {
		return c.String(200, "Hallo")
	})

	return app
}