package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"io"
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

func GetNowShowing(c echo.Context) error {
	movies, err := services.GetNowShowing()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, movies)
}

func GetUpcoming(c echo.Context) error {
	movies, err := services.GetUpcoming()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}
	return c.JSON(http.StatusOK, movies)
}

func CreateMovie(c echo.Context) error {
	movie := new(models.Movie)
	if err := c.Bind(movie); err != nil {
		return utils.BadRequestResponse(c, "Invalid request data", nil)
	}

	movie.Title = c.FormValue("title")
	movie.Description = c.FormValue("description")
	movie.Genre = c.FormValue("genre")
	movie.Language = c.FormValue("language")
	movie.ReleaseDate = c.FormValue("release_date")
	movie.Rating = c.FormValue("rating")
	movie.Duration = c.FormValue("duration")

	validate := validator.New()
	if err := validate.Struct(movie); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return utils.BadRequestResponse(c, "Field "+err.Field()+" cannot be empty", nil)
		}
	}

	file, err := c.FormFile("image")
	if err != nil {
		return utils.BadRequestResponse(c, "Image is required", nil)
	}

	src, err := file.Open()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to process image", nil)
	}

	err = services.CreateMovie(movie, fileBytes)
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