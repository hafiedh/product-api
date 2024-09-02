package healthcheck

import (
	"context"
	"time"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Validate() Service {
	return s
}

func (s *service) HealthCheck(ctx context.Context) (res HealthCheckResponse, err error) {
	res = HealthCheckResponse{
		Message:    "Server up and running",
		ServerTime: time.Now().Format(time.RFC1123),
		Version:    "v1.0.0",
	}

	return
}
