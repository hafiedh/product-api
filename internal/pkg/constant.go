package pkg

type (
	DefaultResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Data    any    `json:"data"`
		Errors  any    `json:"errors"`
	}

	PaginationRequest struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}

	PaginationResponse struct {
		Results any `json:"result"`
		DefaultMetaData
	}

	DefaultMetaData struct {
		Page        uint `json:"page"`
		TotalPages  uint `json:"totalPages"`
		TotalItems  uint `json:"totalItems"`
		Limit       uint `json:"limit"`
		HasNext     bool `json:"hasNext"`
		HasPrevious bool `json:"hasPrevious"`
	}
)
