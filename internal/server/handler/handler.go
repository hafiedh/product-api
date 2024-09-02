package handler

import (
	"api-product/internal/infrastructure/container"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Handler struct {
	postgres           *pgxpool.Pool
	healthCheckHandler *healthCheckHandler
}

func SetupHandler(container *container.Container) *Handler {
	return &Handler{
		postgres:           container.PostgresDB,
		healthCheckHandler: NewHealthCheckHandler().SetHealthCheckService(container.HealthCheckService).Validate(),
	}
}

func (h *Handler) Validate() *Handler {
	if h.healthCheckHandler == nil {
		panic("healthCheckHandler is nil")
	}
	return h
}
