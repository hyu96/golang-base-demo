package service

import (
	"github.com/google/wire"
)

type ProductService struct {
}

func NewProductService() ProductService {
	return ProductService{}
}

var ProviderSet = wire.NewSet(NewProductService)
