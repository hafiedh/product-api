package healthcheck

type HealthCheckResponse struct {
	Message    string `json:"message"`
	ServerTime string `json:"server_time"`
	Version    string `json:"version"`
}
