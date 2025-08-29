package handler

import (
	"eticket-test/internal/pkg/logger"
	"eticket-test/modules/auth/domain/service"
	"github.com/labstack/echo"
	"net/http"

)

type AuthHandler struct {
	logger      *logger.Logger
	authService *service.AuthService
}

func NewAuthHandler(log *logger.Logger, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{logger: log, authService: authService}
}

func (h *AuthHandler) RegisterRoutes(e *echo.Echo, basePath string) {
	e.POST(basePath+"/login", h.Login)
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	tokenResp, err := h.authService.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		h.logger.Error("login failed: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, tokenResp)
}
