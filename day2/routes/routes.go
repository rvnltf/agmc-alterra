package routes

import (
	"agmc-api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	v1User := v1.Group("/users")

	v1User.GET("", controllers.V1GetUserControllers)
	v1User.GET("/:id", controllers.V1GetUserByIdControllers)
	v1User.POST("", controllers.V1CreateUserController)
	v1User.PUT("/:id", controllers.V1UpdateUserByIdControllers)
	v1User.DELETE("/:id", controllers.V1DeleteUserControllers)

	v1Book := v1.Group("/books")

	v1Book.GET("", controllers.V1GetBookControllers)
	v1Book.GET("/:id", controllers.V1GetBookByIdControllers)
	v1Book.POST("", controllers.V1CreateBookController)
	v1Book.PUT("/:id", controllers.V1UpdateBookByIdControllers)
	v1Book.DELETE("/:id", controllers.V1DeleteBookControllers)

	return e
}
