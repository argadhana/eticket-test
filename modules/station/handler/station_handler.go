package handler

import (
	"eticket-test/internal/pkg/logger"
	"eticket-test/modules/station/domain/service"
	"eticket-test/modules/station/dto/request"
	"eticket-test/modules/station/dto/response"
	"github.com/labstack/echo"
	"net/http"
)

type StationHandler struct {
	log            *logger.Logger
	stationService *service.StationService
}

func NewStationHandler(log *logger.Logger, service *service.StationService) *StationHandler {
	return &StationHandler{log: log, stationService: service}
}

// sekarang register ke group (sudah protected middleware)
func (h *StationHandler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateStation)
}

func (h *StationHandler) CreateStation(c echo.Context) error {
	var req request.CreateStationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}

	station, err := h.stationService.CreateStation(c.Request().Context(), req.Name, req.Location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := response.StationResponse{
		ID:       station.ID,
		Name:     station.Name,
		Location: station.Location,
	}

	return c.JSON(http.StatusCreated, resp)
}
