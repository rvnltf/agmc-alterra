package routes

import (
	"agmc-api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// e.GET("/v1/users", controllers.GetUserControllers)
	v1 := e.Group("/v1")
	v1User := v1.Group("/users")

	v1User.GET("", controllers.V1GetUserControllers)
	v1User.GET("/:id", controllers.V1GetUserByIdControllers)
	v1User.POST("", controllers.V1CreateUserController)
	v1User.PUT("/:id", controllers.V1UpdateUserByIdControllers)
	v1User.DELETE("/:id", controllers.V1DeleteUserControllers)

	return e
}
