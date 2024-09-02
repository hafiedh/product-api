CREATE TABLE IF NOT EXISTS products (
   id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    variety VARCHAR(255) NOT NULL,
    rating NUMERIC(3, 2) CHECK (rating >= 0 AND rating <= 5),
    stock INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(255) DEFAULT NULL
);
 CREATE INDEX idx_products_name ON products USING btree (name);