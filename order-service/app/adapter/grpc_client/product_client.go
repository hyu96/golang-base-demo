package grpc_client

import (
	"context"
	"github.com/google/wire"
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

func (g *GrpcProductServiceClient) GetProducts(
	ctx context.Context,
	in *pbProductMgmt.GetProductRequest,
) ([]*pbProductMgmt.Product, error) {

	ctx, cancel := context.WithTimeout(ctx, KDefaultTimeout)
	defer cancel()

	res, err := g.client.GetProducts(ctx, in)
	if err != nil {
		return nil, err
	}

	return res.Products, nil
}

type IGrpcProductServiceClient interface {
	GetProducts(ctx context.Context, in *pbProductMgmt.GetProductRequest) ([]*pbProductMgmt.Product, error)
}

var ProviderSet = wire.NewSet(NewGrpcProductMgmtServiceClient)
