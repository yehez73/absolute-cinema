package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error { // For all users
	users, err := services.GetUsers()

	if err != nil {
		log.Print(err)
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, users)
}

func GetSpecUser(c echo.Context) error { // Just one user
	id := c.Param("id")
	user, err := services.GetSpecUser(id)

	if err != nil {
		log.Print(err)
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	if err := services.CreateUser(*user); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user", err)
	}

	return utils.SuccessResponse(c, "User created successfully", nil)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	err := services.UpdateUser(id, *user)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user", err)
	}

	return utils.SuccessResponse(c, "User updated successfully", nil)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user", err)
	}

	return utils.SuccessResponse(c, "User deleted successfully", nil)
}