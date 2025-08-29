package middleware

import (
	"eticket-test/internal/pkg/config"
	"eticket-test/internal/pkg/jwt"
	"github.com/labstack/echo"
)

// MiddlewareProvider provides middlewares for the application
type MiddlewareProvider struct {
	jwtService jwt.JWT
}

// NewMiddlewareProvider creates a new middleware provider
func NewMiddlewareProvider() *MiddlewareProvider {
	signatureKey := config.GetString("jwt.signature_key")
	dayExpired := config.GetInt("jwt.day_expired")

	jwtService := jwt.NewJWTImpl(signatureKey, dayExpired)

	return &MiddlewareProvider{
		jwtService: jwtService,
	}
}

// GetJWTAuth returns JWT authentication middleware
func (p *MiddlewareProvider) GetJWTAuth() echo.MiddlewareFunc {
	return JWTAuth(p.jwtService)
}
