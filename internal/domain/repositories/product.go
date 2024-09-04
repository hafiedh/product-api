package repositories

import (
	"api-product/internal/domain/entities"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	ProductRepo interface {
		CreateProduct(ctx context.Context, product entities.Product) (err error)
		FindAllProduct(ctx context.Context, limit, offset int) (products []entities.Product, total int, err error)
		FindProduct(ctx context.Context, id int) (product entities.Product, err error)
		UpdateProduct(ctx context.Context, product entities.Product) (err error)
		SoftDeleteProduct(ctx context.Context, id int) (err error)
	}

	productRepo struct {
		db *pgxpool.Pool
	}
)

func NewProductRepo(db *pgxpool.Pool) ProductRepo {
	if db == nil {
		panic("ProductRepo: db is nil")
	}

	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(ctx context.Context, product entities.Product) (err error) {
	commandTag, err := r.db.Exec(ctx, QueryCreateProduct, product.Name, product.Description, product.Price, product.Variety, product.Rating, product.Stock, product.Category)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.CreateProduct]", "error when exec query", err.Error())
		err = fmt.Errorf("cannot create product")
		return
	}

	if commandTag.RowsAffected() == 0 {
		err = fmt.Errorf("cannot create product")
		slog.ErrorContext(ctx, "[ProductRepo.CreateProduct]", "no rows affected", commandTag.RowsAffected())
		return err
	}

	return
}
func (r *productRepo) FindAllProduct(ctx context.Context, limit, offset int) (products []entities.Product, total int, err error) {
	rows, err := r.db.Query(ctx, QueryFindAllProduct, limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.FindAllProduct]", "error when query", err)
		err = fmt.Errorf("cannot find product")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&total, &product.ID, &product.Name, &product.Description, &product.Price, &product.Variety, &product.Rating, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			err = fmt.Errorf("cannot find product")
			slog.ErrorContext(ctx, "[ProductRepo.FindAllProduct]", "error when scan", err)
			return
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		err = fmt.Errorf("cannot find product")
		slog.ErrorContext(ctx, "[ProductRepo.FindAllProduct]", "error when rows", err)
		return
	}
	return
}

func (r *productRepo) UpdateProduct(ctx context.Context, product entities.Product) (err error) {
	query := UpdateProduct
	switch {
	case product.Name != "" && product.Description != "" && product.Price != 0 && product.Variety != "" && product.Rating != 0 && product.Stock != 0 && product.Category != "":
		query += fmt.Sprintf(" name = '%s', description = '%s', price = %f, variety = '%s', rating = %d, stock = %d, category = '%s', updated_at = NOW()", product.Name, product.Description, product.Price, product.Variety, product.Rating, product.Stock, product.Category)
	case product.Name != "" && product.Description != "" && product.Price != 0 && product.Variety != "" && product.Rating != 0 && product.Stock != 0:
		query += fmt.Sprintf(" name = '%s', description = '%s', price = %f, variety = '%s', rating = %d, stock = %d, updated_at = NOW()", product.Name, product.Description, product.Price, product.Variety, product.Rating, product.Stock)
	case product.Name != "" && product.Description != "" && product.Price != 0 && product.Variety != "" && product.Rating != 0:
		query += fmt.Sprintf(" name = '%s', description = '%s', price = %f, variety = '%s', rating = %d, updated_at = NOW()", product.Name, product.Description, product.Price, product.Variety, product.Rating)
	case product.Name != "" && product.Description != "" && product.Price != 0 && product.Variety != "":
		query += fmt.Sprintf(" name = '%s', description = '%s', price = %f, variety = '%s', updated_at = NOW()", product.Name, product.Description, product.Price, product.Variety)
	case product.Name != "" && product.Description != "" && product.Price != 0:
		query += fmt.Sprintf(" name = '%s', description = '%s', price = %f, updated_at = NOW()", product.Name, product.Description, product.Price)
	case product.Name != "" && product.Description != "":
		query += fmt.Sprintf(" name = '%s', description = '%s', updated_at = NOW()", product.Name, product.Description)
	case product.Name != "":
		query += fmt.Sprintf(" name = '%s', updated_at = NOW()", product.Name)
	case product.Description != "":
		query += fmt.Sprintf(" description = '%s', updated_at = NOW()", product.Description)
	case product.Price != 0:
		query += fmt.Sprintf(" price = %f, updated_at = NOW()", product.Price)
	case product.Variety != "":
		query += fmt.Sprintf(" variety = '%s', updated_at = NOW()", product.Variety)
	case product.Rating != 0:
		query += fmt.Sprintf(" rating = %f, updated_at = NOW()", float64(product.Rating))
	case product.Stock != 0:
		query += fmt.Sprintf(" stock = %d, updated_at = NOW()", product.Stock)
	case product.Category != "":
		query += fmt.Sprintf(" category = '%s', updated_at = NOW()", product.Category)
	default:
		slog.ErrorContext(ctx, "[ProductRepo.UpdateProduct]", "no data to update", nil)
		return
	}

	query += fmt.Sprintf(" WHERE id = %d", product.ID)

	commandTag, err := r.db.Exec(ctx, query)
	if err != nil {
		err = fmt.Errorf("cannot update product")
		slog.ErrorContext(ctx, "[ProductRepo.UpdateProduct]", "error when exec query", err)
		return
	}

	if commandTag.RowsAffected() == 0 {
		err = fmt.Errorf("cannot update product")
		slog.ErrorContext(ctx, "[ProductRepo.UpdateProduct]", "no rows affected", commandTag.RowsAffected())
		return err
	}
	return
}

func (r *productRepo) SoftDeleteProduct(ctx context.Context, id int) (err error) {
	commandTag, err := r.db.Exec(ctx, QuerySoftDeleteProduct, id)
	if err != nil {
		err = fmt.Errorf("cannot delete product")
		slog.ErrorContext(ctx, "[ProductRepo.SoftDeleteProduct]", "error when exec query", err)
		return
	}

	if commandTag.RowsAffected() == 0 {
		err = fmt.Errorf("cannot delete product")
		slog.ErrorContext(ctx, "[ProductRepo.SoftDeleteProduct]", "no rows affected", commandTag.RowsAffected())
		return
	}
	return
}

func (r *productRepo) FindProduct(ctx context.Context, id int) (product entities.Product, err error) {
	err = r.db.QueryRow(ctx, QueryFindByIdProduct, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Variety, &product.Rating, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.FindProduct]", "error when query", err)
		err = fmt.Errorf("cannot find product")
		return
	}
	return
}
