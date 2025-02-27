package product_handler

import (
	"github.com/google/wire"
	"github.com/huydq/product-service/app/core/service"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
)

type ProductGrpcHandler struct {
	pbProductMgmt.UnimplementedProductManagementServiceServer
	productService service.ProductService
}

func NewProductGrpcHandler(productService service.ProductService) *ProductGrpcHandler {
	return &ProductGrpcHandler{
		productService: productService,
	}
}

var ProviderSet = wire.NewSet(NewProductGrpcHandler)
