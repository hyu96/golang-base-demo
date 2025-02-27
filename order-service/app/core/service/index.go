package service

import (
	"github.com/google/wire"
	"github.com/huydq/order-service/app/adapter/grpc_client"
	"github.com/huydq/order-service/app/core/repository"
)

type OrderService struct {
	productGrpcClient grpc_client.IGrpcProductServiceClient
	orderRepo         repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository, productGrpcClient grpc_client.IGrpcProductServiceClient) OrderService {
	return OrderService{
		orderRepo:         orderRepo,
		productGrpcClient: productGrpcClient,
	}
}

var ProviderSet = wire.NewSet(NewOrderService)
