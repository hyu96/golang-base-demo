package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/util"
	pbOrderMgmt "github.com/huydq/proto/gen-go/order"
)

// CreateOrder handles order creation
func (uc *OrderService) CreateOrder(ctx context.Context, orderDto pbOrderMgmt.CreateOrderRequest) (*pbOrderMgmt.CreateOrderResponse, error) {
	order := models.Order{
		CustomerID: int(orderDto.CustomerId),
	}

	orderItems := make([]models.OrderItem, len(orderDto.Items))
	for i, item := range orderDto.Items {
		orderItems[i] = models.OrderItem{
			ProductID: int(item.ProductId),
			Quantity:  int(item.Quantity),
		}
	}

	orderAgg := models.OrderAggregate{
		Order: order,
		Items: orderItems,
	}
	orderId, err := uc.orderRepo.CreateOrder(ctx, orderAgg)
	if err != nil {
		log.Errorf("Create Order failed", err.Error())
		return nil, errors.New(util.ERR_INTERNAL_SERVER_ERROR)
	}

	return &pbOrderMgmt.CreateOrderResponse{
		OrderId: int32(orderId),
	}, nil
}
