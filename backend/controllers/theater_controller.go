package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetTheaters(c echo.Context) error {
	theaters, err := services.GetTheaters()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, theaters)
}

func GetSpecTheater(c echo.Context) error {
	theater, err := services.GetSpecTheater(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, theater)
}

func CreateTheater(c echo.Context) error {
	theater := new(models.Theater)
	if err := c.Bind(theater); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(theater); err != nil {
		for _, err := range err.(validator.	ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.CreateTheater(theater)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.CreatedResponse(c, "Theater created", nil)
}

func UpdateTheater(c echo.Context) error {
	theater := new(models.Theater)
	if err := c.Bind(theater); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(theater); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.UpdateTheater(c.Param("id"), theater)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Theater updated", nil)
}

func DeleteTheater(c echo.Context) error {
	err := services.DeleteTheater(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Theater deleted", nil)
}