package rmiddleware

import (
	tokenjwt "backend/pkg/token_jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Middleware для JWT аутентификации
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header format"})
		}

		tokenStr := parts[1]
		claims, err := tokenjwt.DecodeJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}

		// Добавляем UserID в context, чтобы использовать в хендлерах
		c.Set("userID", claims.UserID)

		return next(c)
	}
}
