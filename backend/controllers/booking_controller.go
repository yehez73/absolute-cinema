package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetBookings(c echo.Context) error {
	bookings, err := services.GetBookings()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, bookings)
}

func GetSpecBooking(c echo.Context) error {
	booking, err := services.GetSpecBooking(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, booking)
}

func CreateBooking(c echo.Context) error {
	booking := new(models.Booking)
	if err := c.Bind(booking); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(booking); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}	
	}

	token := c.Request().Header.Get("Authorization")
	ID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}
	
	err = services.CreateBooking(booking, ID)
	if err != nil {
		if strings.Contains(err.Error(), "is already booked") {
			return utils.BadRequestResponse(c, err.Error(), nil)
		}
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.CreatedResponse(c, "Booking created", nil)
}

func UpdateBooking(c echo.Context) error {
	booking := new(models.Booking)
	if err := c.Bind(booking); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	err := services.UpdateBooking(c.Param("id"), booking)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Booking updated", nil)
}

func DeleteBooking(c echo.Context) error {
	err := services.DeleteBooking(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Booking deleted", nil)
}