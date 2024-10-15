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
        return utils.BadRequestResponse(c, "Invalid request data", nil)
    }

    if err := services.Register(*user); err != nil {
        return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user", err)
    }

    return utils.SuccessResponse(c, "User created successfully", nil)
}

func Login(c echo.Context) error {
    loginReq := new(LoginRequest)
    if err := c.Bind(loginReq); err != nil {
        return utils.BadRequestResponse(c, "Invalid request data", nil)
    }

    user, err := services.Login(loginReq.Email, loginReq.Password)
    if err != nil {
        return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password", nil)
    }

    token, err := utils.GenerateToken(user.ID.Hex(), user.Name, string(user.Role))
    if err != nil {
        return utils.InternalServerErrorResponse(c, "Failed to generate token", nil)
    }

    responseData := map[string]interface{}{
        "token": token,
        "user": map[string]interface{}{
            "name":  user.Name,
            "role":  user.Role,
        },
    }

    return utils.SuccessResponse(c, "User logged in successfully", responseData)
}

func Logout(c echo.Context) error {
    token := c.Request().Header.Get("Authorization")
    utils.InvalidateToken(token)

    return utils.SuccessResponse(c, "User logged out successfully", nil)
}
