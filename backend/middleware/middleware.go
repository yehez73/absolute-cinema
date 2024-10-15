package middleware

import (
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        _, exists := utils.InvalidTokens[token]
		if exists {
			return utils.ErrorResponse(c, http.StatusUnauthorized, "Token tidak valid atau Anda telah logout", nil)
		}

        if token == "" {
            return utils.ErrorResponse(c, http.StatusUnauthorized, "Missing or invalid token", nil)
        }

        claims, err := utils.ValidateToken(token)
        if err != nil {
            return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", nil)
        }

        c.Set("userId", claims["userId"])
        return next(c)
    }
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")

        _, exists := utils.InvalidTokens[token]
		if exists {
			return utils.ErrorResponse(c, http.StatusUnauthorized, "Token tidak valid atau Anda telah logout", nil)
		}

        if token == "" {
            return utils.ErrorResponse(c, http.StatusUnauthorized, "Missing or invalid token", nil)
        }

        claims, err := utils.ValidateToken(token)
        if err != nil {
            return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", nil)
        }

        if claims["role"] != "admin" {
            return utils.ErrorResponse(c, http.StatusForbidden, "You don't have permission to access this resource", nil)
        }

        c.Set("userId", claims["userId"])
        return next(c)
    }
}