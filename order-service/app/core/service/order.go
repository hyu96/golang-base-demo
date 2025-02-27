package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/util"
	pbOrderMgmt "github.com/huydq/proto/gen-go/order"
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
)

// CreateOrder handles order creation
func (uc *OrderService) CreateOrder(ctx context.Context, orderDto pbOrderMgmt.CreateOrderRequest) (*pbOrderMgmt.CreateOrderResponse, error) {
	order := models.Order{
		CustomerID: int(orderDto.CustomerId),
	}

	productIdList := make([]int32, len(orderDto.Items))
	for index, item := range orderDto.Items {
		productIdList[index] = item.ProductId
	}

	products, err := uc.productGrpcClient.GetProducts(ctx, &pbProductMgmt.GetProductRequest{
		Id: productIdList,
	})

	productsMap := make(map[int32]*pbProductMgmt.Product, len(products))
	for _, product := range products {
		productsMap[product.Id] = product
	}

	orderItems := make([]models.OrderItem, len(orderDto.Items))
	for i, item := range orderDto.Items {
		orderItems[i] = models.OrderItem{
			ProductID: int(item.ProductId),
			Price:     float64(productsMap[item.ProductId].Price),
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
