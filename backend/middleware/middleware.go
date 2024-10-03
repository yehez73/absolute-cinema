package middleware

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        _, exists := utils.InvalidTokens[token]
		if exists {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak valid atau Anda telah logout",
				Status:  false,
			})
		}

        if token == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
        }

        claims, err := utils.ValidateToken(token)
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
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
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak valid atau Anda telah logout",
				Status:  false,
			})
		}

        if token == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
        }

        claims, err := utils.ValidateToken(token)
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        if claims["role"] != "admin" {
            return echo.NewHTTPError(http.StatusForbidden, "Unauthorized")
        }

        c.Set("userId", claims["userId"])
        return next(c)
    }
}