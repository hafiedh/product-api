package handler

import (
	"api-product/internal/infrastructure/container"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, cnt *container.Container) {
	h := SetupHandler(cnt).Validate()

	e.GET("/", h.healthCheckHandler.HealthCheck)

	product := e.Group("/v1/products")
	{
		product.POST("", h.productHandler.CreateProduct)
		product.GET("", h.productHandler.GetProducts)
		product.GET("/:id", h.productHandler.GetProduct)
		product.PUT("/:id", h.productHandler.UpdateProduct)
		product.DELETE("/:id", h.productHandler.DeleteProduct)
	}

}
