package product_handler

import (
	"context"
	"errors"
	"github.com/huydq/product-service/util"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
)

func (h *ProductGrpcHandler) GetProducts(ctx context.Context, request *pbProductMgmt.GetProductRequest) (*pbProductMgmt.GetProductResponse, error) {
	if len(request.GetId()) == 0 {
		return nil,
			errors.New(util.ERR_INVALID_PRODUCT_ID)
	}

	productIdList := make([]int, len(request.GetId()))
	for index, productId := range request.GetId() {
		productIdList[index] = int(productId)
	}

	products, err := h.productService.GetProducts(ctx, productIdList)
	if err != nil {
		return nil, err
	}
	return &pbProductMgmt.GetProductResponse{
		Products: products,
	}, nil
}
