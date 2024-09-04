package repositories

const (
	QueryCreateProduct = `INSERT INTO products (name, description, price, variety, rating, stock, category) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	QueryFindAllProduct = `SELECT COUNT(*) OVER() AS total_count,
       id,
       name,
       description,
       price,
       variety,
       rating,
       stock,
       category,
       created_at,
       updated_at
		FROM products
		WHERE deleted_at IS NULL
		ORDER BY id DESC
		LIMIT $1 OFFSET $2`

	QueryFindByIdProduct = `SELECT id, name, description, price, variety, rating, stock, category, created_at, updated_at
	FROM products WHERE id = $1 AND deleted_at IS NULL`

	UpdateProduct = ` UPDATE products SET`

	QuerySoftDeleteProduct = `
	UPDATE products SET deleted_at = NOW() WHERE id = $1
	`
)
