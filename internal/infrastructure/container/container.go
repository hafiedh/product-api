package container

import (
	"fmt"
	"os"

	"api-product/internal/config"
	"api-product/internal/infrastructure/postgres"
	"api-product/internal/usecase/healthcheck"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Container struct {
	Config             *config.DefaultConfig
	PostgresDB         *pgxpool.Pool
	HealthCheckService healthcheck.Service
}

func (c *Container) Validate() *Container {
	if c.Config == nil {
		panic("Config is nil")
	}
	if c.HealthCheckService == nil {
		panic("HealthCheckService is nil")
	}
	if c.PostgresDB == nil {
		panic("PostgresDB is nil")
	}
	return c
}

func New() *Container {

	config.Load(os.Getenv("env"), ".env")

	defConfig := &config.DefaultConfig{
		Apps: config.Apps{
			Name:     config.GetString("app.name"),
			Address:  config.GetString("address"),
			HttpPort: config.GetString("port"),
		},
	}

	postgresConfig := &config.PostgreSQLDB{
		Host:         config.GetString("postgresql.product_db.host"),
		User:         config.GetString("postgresql.product_db.user"),
		Password:     config.GetString("postgresql.product_db.password"),
		Name:         config.GetString("postgresql.product_db.db"),
		Port:         config.GetInt("postgresql.product_db.port"),
		SSLMode:      config.GetString("postgresql.product_db.ssl"),
		Schema:       config.GetString("postgresql.product_db.schema"),
		Debug:        config.GetBool("postgresql.product_db.debug"),
		PoolMaxConns: config.GetInt("postgresql.product_db.poolMaxConns"),
	}

	postgresDB, err := postgres.NewDB(*postgresConfig)
	if err != nil {
		fmt.Printf("Error connecting to PostgreSQL database: %v", err)
	}
	healthCheckService := healthcheck.NewService().Validate()

	container := &Container{
		Config:             defConfig,
		HealthCheckService: healthCheckService,
		PostgresDB:         postgresDB,
	}
	container.Validate()
	return container

}
