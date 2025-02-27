package grpc_client

import (
	"context"
	"github.com/huydq/gokits/libs/client/grpc"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
	"time"
)

const (
	KDefaultTimeout = 30 * time.Second
)

type GrpcProductServiceClient struct {
	client pbProductMgmt.ProductManagementServiceClient
}

func NewGrpcProductServiceClient() *GrpcProductServiceClient {
	return &GrpcProductServiceClient{
		pbProductMgmt.NewProductManagementServiceClient(grpc.GetClientConn(pbProductMgmt.ProductManagementService_ServiceDesc.ServiceName)),
	}
}

func (g *GrpcProductServiceClient) GetProduct(
	ctx context.Context,
	in *pbProductMgmt.GetProductRequest,
) (*pbProductMgmt.Product, error) {

	ctx, cancel := context.WithTimeout(ctx, KDefaultTimeout)
	defer cancel()

	return g.client.GetProduct(ctx, in)
}
