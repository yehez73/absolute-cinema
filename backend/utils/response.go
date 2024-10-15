package utils

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Status  bool        `json:"status"`
    Data    interface{} `json:"data,omitempty"`
}

// Code 200
func SuccessResponse(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusOK, Response{
        Code:    http.StatusOK,
        Message: message,
        Status:  true,
        Data:    data,
    })
}

// Code 201
func CreatedResponse(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusCreated, Response{
        Code:    http.StatusCreated,
        Message: message,
        Status:  true,
        Data:    data,
    })
}

// Code 400
func BadRequestResponse(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusBadRequest, Response{
        Code:    http.StatusBadRequest,
        Message: message,
        Status:  false,
        Data:    data,
    })
}

// Code 404
func NotFoundResponse(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusNotFound, Response{
        Code:    http.StatusNotFound,
        Message: message,
        Status:  false,
        Data:    data,
    })
}

// Code 500
func InternalServerErrorResponse(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusInternalServerError, Response{
        Code:    http.StatusInternalServerError,
        Message: message,
        Status:  false,
        Data:    data,
    })
}

// Custom Error Response
func ErrorResponse(c echo.Context, statusCode int, message string, data interface{}) error {
    return c.JSON(statusCode, Response{
        Code:    statusCode,
        Message: message,
        Status:  false,
        Data:    data,
    })
}