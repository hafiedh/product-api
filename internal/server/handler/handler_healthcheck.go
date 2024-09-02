package handler

import (
	"net/http"

	"api-product/internal/usecase/healthcheck"

	"github.com/labstack/echo/v4"
)

type healthCheckHandler struct {
	healthCheckService healthcheck.Service
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) SetHealthCheckService(service healthcheck.Service) *healthCheckHandler {
	h.healthCheckService = service
	return h
}

func (h *healthCheckHandler) Validate() *healthCheckHandler {
	if h.healthCheckService == nil {
		panic("healthCheckService is nil")
	}

	return h
}

func (h *healthCheckHandler) HealthCheck(c echo.Context) (err error) {
	ctx := c.Request().Context()

	res, err := h.healthCheckService.HealthCheck(ctx)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}
