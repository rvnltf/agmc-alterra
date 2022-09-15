package controllers

import (
	"agmc-api/lib/database"

	// "agmc-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func V1GetUserControllers(c echo.Context) error {
	user, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully get data.",
		"data":    user,
	})
}

func V1GetUserByIdControllers(c echo.Context) error {
	user, e := database.GetUserById(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully get data.",
		"data":    user,
	})
}

func V1CreateUserController(c echo.Context) error {
	user, e := database.CreateUser(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully insert data.",
		"data":    user,
	})
}

func V1UpdateUserByIdControllers(c echo.Context) error {
	user, e := database.UpdateUser(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully update data.",
		"data":    user,
	})
}

func V1DeleteUserControllers(c echo.Context) error {
	_, e := database.DeleteById(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully delete data.",
	})
}
