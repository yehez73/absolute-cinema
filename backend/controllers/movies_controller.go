package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func GetMovies(c echo.Context) error {
	movies, err := services.GetMovies()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, movies)
}

func GetSpecMovie(c echo.Context) error {
	movie, err := services.GetSpecMovie(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, movie)
}

func CreateMovie(c echo.Context) error {
	movie := new(models.Movie)
	if err := c.Bind(movie); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}
	
	validate := validator.New()
	if err := validate.Struct(movie); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.CreateMovie(movie)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.CreatedResponse(c, "Movie created", nil)
}

func UpdateMovie(c echo.Context) error {
	movie := new(models.Movie)
	if err := c.Bind(movie); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(movie); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	err := services.UpdateMovie(c.Param("id"), movie)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Movie updated", nil)
}

func DeleteMovie(c echo.Context) error {
	err := services.DeleteMovie(c.Param("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return utils.SuccessResponse(c, "Movie deleted", nil)
}