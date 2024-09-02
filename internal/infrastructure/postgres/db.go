package postgres

import (
	"context"
	"fmt"
	"log"

	"api-product/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(cfg config.PostgreSQLDB) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode, cfg.PoolMaxConns)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	log.Printf("Connected to PostgreSQL database %s", cfg.Name)
	return db, nil
}
