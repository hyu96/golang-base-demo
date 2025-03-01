package dto

type GetProductsRequestDTO struct {
	ProductIDs []int `json:"product_ids"`
}

type GetProductsResponseDTO struct {
	Products []*ProductDTO `json:"product_ids"`
}

type ProductDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
