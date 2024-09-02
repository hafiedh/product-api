package cmd

import (
	"api-product/internal/infrastructure/container"
	"api-product/internal/server"
)

func Run() {
	server.StartService(container.New())
}
