package controllers

import (
	"backend/models"
	"backend/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error { // For all users
	users, err := services.GetUsers()

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code": "500",
			"message": "Internal server error",
			"status": "error",
		})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error { // Just one user
	id := c.Param("id")
	user, err := services.GetUser(id)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code": "500",
			"message": "Internal server error",
			"status": "error",
		})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := services.CreateUser(*user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code": "201",
		"message": "User created successfully",
		"status": "success",
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := services.UpdateUser(id, *user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": "200",
		"message": "User updated successfully",
		"status": "success",
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": "200",
		"message": "User deleted successfully",
		"status": "success",
	})
}