package grpc_client

import (
	"context"
	"github.com/google/wire"
	"github.com/huydq/order-service/internal/core/domain/dto"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
	"time"

	"github.com/huydq/gokits/libs/client/grpc"
)

const (
	KDefaultTimeout = 30 * time.Second
)

type GrpcProductServiceClient struct {
	client pbProductMgmt.ProductManagementServiceClient
}

func NewGrpcProductMgmtServiceClient() *GrpcProductServiceClient {
	return &GrpcProductServiceClient{
		pbProductMgmt.NewProductManagementServiceClient(grpc.GetClientConn("product-service")),
	}
}

func (g GrpcProductServiceClient) GetProducts(
	ctx context.Context,
	reqDto dto.GetProductsRequestDTO,
) (*dto.GetProductsResponseDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, KDefaultTimeout)
	defer cancel()

	idList := make([]int32, len(reqDto.ProductIDs))
	for index, id := range reqDto.ProductIDs {
		idList[index] = int32(id)
	}

	input := &pbProductMgmt.GetProductRequest{
		Id: idList,
	}
	res, err := g.client.GetProducts(ctx, input)
	if err != nil {
		return nil, err
	}

	resDto := &dto.GetProductsResponseDTO{
		Products: make([]*dto.ProductDTO, len(res.Products)),
	}

	for index, product := range res.Products {
		resDto.Products[index] = &dto.ProductDTO{
			ID:    int(product.Id),
			Name:  product.Name,
			Price: float64(product.Price),
		}
	}

	return resDto, nil
}

var ProviderSet = wire.NewSet(NewGrpcProductMgmtServiceClient)
