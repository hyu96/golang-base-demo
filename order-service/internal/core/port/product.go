package port

import (
	"context"
	"github.com/huydq/order-service/internal/core/domain/dto"
)

//go:generate mockgen -source=product.go -destination=./mocks/product.go

type IProductServiceClient interface {
	GetProducts(ctx context.Context, reqDto dto.GetProductsRequestDTO) (*dto.GetProductsResponseDTO, error)
}
