package product

import (
	"api-product/internal/domain/entities"
	"api-product/internal/domain/repositories"
	"api-product/internal/pkg"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
)

type (
	ProductSvc interface {
		CreateProduct(ctx context.Context, req entities.Product) (err error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest, id int) (err error)
		GetProduct(ctx context.Context, id int) (res entities.Product, err error)
		GetProducts(ctx context.Context, req pkg.PaginationRequest) (res pkg.PaginationResponse, err error)
		DeleteProduct(ctx context.Context, id int) (err error)
	}

	productSvc struct {
		productRepo repositories.ProductRepo
	}
)

func NewProductSvc(productRepo repositories.ProductRepo) ProductSvc {
	if productRepo == nil {
		panic("ProductRepo is required")
	}

	return &productSvc{
		productRepo: productRepo,
	}
}

func (s *productSvc) CreateProduct(ctx context.Context, req entities.Product) (err error) {
	err = s.productRepo.CreateProduct(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvc.CreateProduct] Error: %v", err)
		err = fmt.Errorf("failed to create product")
		return
	}
	return
}

func (s *productSvc) UpdateProduct(ctx context.Context, req UpdateProductRequest, id int) (err error) {
	productEntity := entities.Product{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Variety:     req.Variety,
		Stock:       req.Stock,
		Rating:      req.Rating,
		Category:    req.Category,
	}

	err = s.productRepo.UpdateProduct(ctx, productEntity)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvc.UpdateProduct] Error: %v", err)
		return
	}
	return

}

func (s *productSvc) GetProduct(ctx context.Context, id int) (res entities.Product, err error) {
	res, err = s.productRepo.FindProduct(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvc.GetProduct] Error: %v", err)
		err = fmt.Errorf("product not found")
		return
	}
	return
}

func (s *productSvc) GetProducts(ctx context.Context, req pkg.PaginationRequest) (res pkg.PaginationResponse, err error) {
	var dataResp pkg.DefaultResponse
	{
		{
			dataResp.Code = http.StatusOK
			dataResp.Message = "Success"
		}
	}
	offset := 0
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page > 1 {
		offset = int((req.Page - 1) * req.Limit)
	}

	products, total, err := s.productRepo.FindAllProduct(ctx, req.Limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvc.GetProducts] Error: %v", err)
		err = fmt.Errorf("failed to get products")
		return
	}

	{
		res.Page = uint(req.Page)
		res.Limit = uint(req.Limit)

		totalPage := math.Ceil(float64(total) / float64(req.Limit))
		res.TotalPages = uint(totalPage)

		res.TotalItems = uint(total)
		res.HasNext = req.Page < int(totalPage)
		res.HasPrevious = req.Page > 1
		res.Results = products
	}
	return
}

func (s *productSvc) DeleteProduct(ctx context.Context, id int) (err error) {
	err = s.productRepo.SoftDeleteProduct(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvc.DeleteProduct] Error: %v", err)
		return
	}
	return
}
