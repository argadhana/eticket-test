package middleware

import (
	"eticket-test/internal/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// JWTAuth returns a middleware that secures routes with JWT authentication
func JWTAuth(jwtService jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "unauthorized, missing token",
				})
			}

			// Extract the token from the Bearer header
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "unauthorized, invalid token format",
				})
			}

			token := tokenParts[1]
			valid, err := jwtService.ValidateToken(token)
			if !valid || err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "unauthorized, invalid token",
				})
			}

			// Parse the token and extract claims
			claims, err := jwtService.ParseToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "unauthorized, failed to parse token",
				})
			}

			// Set user information in context
			c.Set("user", claims)
			return next(c)
		}
	}
}
