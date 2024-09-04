package handler

import (
	"api-product/internal/infrastructure/container"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Handler struct {
	postgres           *pgxpool.Pool
	healthCheckHandler *healthCheckHandler
	productHandler     *productHandler
}

func SetupHandler(container *container.Container) *Handler {
	return &Handler{
		postgres:           container.PostgresDB,
		healthCheckHandler: NewHealthCheckHandler().SetHealthCheckService(container.HealthCheckService).Validate(),
		productHandler:     NewProductHandler().SetProductService(container.ProductService).Validate(),
	}
}

func (h *Handler) Validate() *Handler {
	if h.healthCheckHandler == nil {
		panic("healthCheckHandler is nil")
	}
	if h.productHandler == nil {
		panic("productHandler is nil")
	}
	return h
}
