package product

type (
	UpdateProductRequest struct {
		Name        string  `json:"name`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Variety     string  `json:"variety"`
		Stock       int     `json:"stock"`
		Rating      int     `json:"rating"`
		Category    string  `json:"category"`
	}
)
