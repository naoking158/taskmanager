package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/auth"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// by default token is stored under `user` key
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "JWT token missing or invalid")
		}

		// by default claims is of type `jwt.MapClaims`
		claims, ok := token.Claims.(*auth.JwtCustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to cast claims as jwt.MapClaims")
		}

		userID := claims.UserID
		c.Set("userID", userID)

		return next(c)
	}
}
