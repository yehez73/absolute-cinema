package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func GetShowtimes(c echo.Context) error {
	showtimes, err := services.GetShowtimes()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, showtimes)
}

func GetSpecShowtime(c echo.Context) error {
	showtime, err := services.GetSpecShowtime(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, showtime)
}

func GetSpecShowtimeByMovieDate(c echo.Context) error {
	movie_id := c.Param("movie_id")
	showdate := c.Param("showdate")
	showtime, err := services.GetSpecShowtimeByMovieDate(movie_id, showdate)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNoContent, "No showtimes found for the given movie and date", nil)
	}
	return c.JSON(http.StatusOK, showtime)
}

func CreateShowtime(c echo.Context) error {
	showtime := new(models.Showtime)
	if err := c.Bind(showtime); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}
	
	validate := validator.New()
	if err := validate.Struct(showtime); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.CreateShowtime(showtime)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.CreatedResponse(c, "Showtime created", nil)
}

func UpdateShowtime(c echo.Context) error {
	showtime := new(models.Showtime)
	if err := c.Bind(showtime); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(showtime); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.UpdateShowtime(c.Param("id"), showtime)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Showtime updated", nil)
}

func DeleteShowtime(c echo.Context) error {
	err := services.DeleteShowtime(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Showtime deleted", nil)
}