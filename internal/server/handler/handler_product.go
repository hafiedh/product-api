package handler

import (
	"api-product/internal/domain/entities"
	"api-product/internal/pkg"
	"api-product/internal/usecase/product"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	productHandler struct {
		productService product.ProductSvc
	}
)

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (h *productHandler) SetProductService(svc product.ProductSvc) *productHandler {
	h.productService = svc
	return h
}

func (h *productHandler) Validate() *productHandler {
	if h.productService == nil {
		panic("healthCheckService is nil")
	}

	return h
}

func (h *productHandler) CreateProduct(c echo.Context) error {
	ctx := c.Request().Context()
	defer recoveryPanicHandler()

	var product entities.Product

	if err := c.Bind(&product); err != nil {
		slog.Error("Failed to bind product: %v", err)

		if strings.ContainsAny(err.Error(), "Unmarshal type error") {
			field := strings.Split(err.Error(), "field=")
			field = strings.Split(field[1], ",")

			return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
				Code:    http.StatusBadRequest,
				Status:  "failed",
				Message: "Invalid request body",
				Data:    struct{}{},
				Errors:  fmt.Sprintf("Invalid data type for field %s", field[0]),
			})
		}
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request body",
			Data:    struct{}{},
			Errors:  err,
		})
	}

	if err := c.Validate(product); err != nil {
		slog.Error("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request body",
			Data:    struct{}{},
			Errors:  err,
		})
	}

	if err := h.productService.CreateProduct(ctx, product); err != nil {
		slog.Error("Failed to create product: %v", err)
		return c.JSON(http.StatusInternalServerError, pkg.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Status:  "failed",
			Message: "Failed to create product",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, pkg.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Status:  "success",
		Data:    struct{}{},
		Errors:  struct{}{},
	})
}

func (h *productHandler) UpdateProduct(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if id == "" {
		slog.Error("Product ID is empty")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("Product ID must be a number")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Status:  "failed",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	var product product.UpdateProductRequest

	if err := c.Bind(&product); err != nil {
		slog.Error("Failed to bind product: %v", err)
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request body",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}
	if err := c.Validate(product); err != nil {
		slog.Error("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request body",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}
	if err := h.productService.UpdateProduct(ctx, product, productID); err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Status:  "failed",
			Message: "Failed to update product",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, pkg.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Status:  "success",
		Data:    struct{}{},
		Errors:  struct{}{},
	})
}

func (h *productHandler) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()
	defer recoveryPanicHandler()

	id := c.Param("id")

	if id == "" {
		slog.Error("Product ID is empty")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("Product ID must be a number")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	product, err := h.productService.GetProduct(ctx, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Status:  "failed",
			Message: "Failed to get product",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, pkg.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Product retrieved successfully",
		Status:  "success",
		Data:    product,
		Errors:  struct{}{},
	})

}

func (h *productHandler) GetProducts(c echo.Context) error {
	ctx := c.Request().Context()

	var pagination pkg.PaginationRequest
	if err := c.Bind(&pagination); err != nil {
		slog.Error("Failed to bind pagination: %v", err)
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	if err := c.Validate(pagination); err != nil {
		slog.Error("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Invalid request",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	products, err := h.productService.GetProducts(ctx, pagination)
	if err != nil {
		slog.Error("Failed to get products: %v", err)
		return c.JSON(http.StatusInternalServerError, pkg.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Status:  "failed",
			Message: "Failed to get products",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, pkg.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Products retrieved successfully",
		Status:  "success",
		Data:    products,
		Errors:  struct{}{},
	})

}

func (h *productHandler) DeleteProduct(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if id == "" {
		slog.Error("Product ID is empty")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Status:  "failed",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("Product ID must be a number")
		return c.JSON(http.StatusBadRequest, pkg.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Status:  "failed",
			Data:    struct{}{},
			Errors:  "invalid product ID",
		})
	}

	if err := h.productService.DeleteProduct(ctx, productID); err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete product",
			Status:  "failed",
			Data:    struct{}{},
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, pkg.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
		Status:  "success",
		Data:    struct{}{},
		Errors:  struct{}{},
	})
}

func recoveryPanicHandler() {
	if r := recover(); r != nil {
		slog.Error("Panic Recovered: %v", r)
	}
}
