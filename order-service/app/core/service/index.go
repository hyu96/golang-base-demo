package service

import (
	"github.com/google/wire"
	"github.com/huydq/order-service/app/core/repository"
)

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderService(orderRepo repository.IOrderRepository) OrderService {
	return OrderService{
		orderRepo: orderRepo,
	}
}

var ProviderSet = wire.NewSet(NewOrderService)
