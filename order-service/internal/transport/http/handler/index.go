package order_handler

import (
	"github.com/google/wire"
	"github.com/huydq/order-service/internal/core/service"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

var ProviderSet = wire.NewSet(NewOrderHandler)
