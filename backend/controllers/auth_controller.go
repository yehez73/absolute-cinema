package controllers

import (
    "net/http"
    "backend/models"
    "backend/services"
    "backend/utils"
    "github.com/labstack/echo/v4"
)

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

func RegisterUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    if err := services.Register(*user); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }   

    return c.JSON(http.StatusCreated, map[string]interface{}{
        "code": "201",
        "message": "User registered successfully",
        "status": "success",
    })
}

func Login(c echo.Context) error {
    loginReq := new(LoginRequest)
    if err := c.Bind(loginReq); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    user, err := services.Login(loginReq.Email, loginReq.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
    }

    token, err := utils.GenerateToken(user.ID.Hex(), user.Name, string(user.Role))
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{
        "code": "200",
        "message": "User logged in successfully",
        "status": "success",
        "token": token,
    })
}

func Logout(c echo.Context) error {
    token := c.Request().Header.Get("Authorization")
    utils.InvalidateToken(token)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "code": "200",
        "message": "User logged out successfully",
        "status": "success",
    })
}
