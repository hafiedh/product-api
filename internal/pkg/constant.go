package pkg

type (
	DefaultResponse struct {
		Message string      `json:"message"`
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
		Errors  []string    `json:"errors"`
	}
)
