package service

import (
	"github.com/google/wire"
	"github.com/huydq/order-service/internal/core/port"
	"github.com/huydq/order-service/internal/core/repository"
)

type OrderService struct {
	orderRepo            repository.IOrderRepository
	productServiceClient port.IProductServiceClient
}

func NewOrderService(orderRepo repository.IOrderRepository, productClient port.IProductServiceClient) OrderService {
	return OrderService{
		orderRepo:            orderRepo,
		productServiceClient: productClient,
	}
}

var ProviderSet = wire.NewSet(NewOrderService)
