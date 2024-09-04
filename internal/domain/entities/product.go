package entities

import "time"

type (
	Product struct {
		ID          int       `db:"id" json:"id"`
		Name        string    `db:"name" json:"name" validate:"required"`
		Description string    `db:"description" json:"description" validate:"required"`
		Price       float64   `db:"price" json:"price" validate:"required"`
		Variety     string    `db:"variety" json:"variety" validate:"required"`
		Rating      int       `db:"rating" json:"rating" validate:"required"`
		Stock       int       `db:"stock" json:"stock" validate:"required"`
		Category    string    `db:"category" json:"category" validate:"required"`
		CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
		CreatedBy   string    `db:"created_by" json:"created_by,omitempty"`
		UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
		UpdatedBy   string    `db:"updated_by" json:"updated_by,omitempty"`
	}
)
