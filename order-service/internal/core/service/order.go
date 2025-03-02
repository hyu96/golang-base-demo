package service

import (
	"context"
	"errors"
	"github.com/huydq/order-service/internal/core/domain/dto"
	"go.elastic.co/apm/v2"

	"github.com/gofiber/fiber/v2/log"
	"github.com/huydq/order-service/internal/core/domain/model"
	"github.com/huydq/order-service/util"
)

// CreateOrder handles order creation
func (uc *OrderService) CreateOrder(ctx context.Context, orderDto dto.CreateOrderRequestDTO) (*dto.CreateOrderResponseDTO, error) {
	// Start APM transaction
	tx := apm.DefaultTracer.StartTransaction("CreateOrder", "request")
	defer tx.End()
	ctx = apm.ContextWithTransaction(ctx, tx)

	// Create order span
	span, ctx := apm.StartSpan(ctx, "CreateOrder", "service")
	defer span.End()

	order := model.Order{
		CustomerID: orderDto.CustomerID,
	}

	// Get products span
	span, ctx = apm.StartSpan(ctx, "GetProducts", "service")
	productReqDto := dto.GetProductsRequestDTO{
		ProductIDs: make([]int, len(orderDto.Items)),
	}
	for index, item := range orderDto.Items {
		productReqDto.ProductIDs[index] = item.ProductID
	}

	productsResDto, err := uc.productServiceClient.GetProducts(ctx, productReqDto)
	span.End()

	if err != nil {
		log.Errorf("GetProducts failed", err.Error())
		return nil, errors.New(util.ERR_INTERNAL_SERVER_ERROR)
	}

	productsMap := make(map[int]*dto.ProductDTO, len(productsResDto.Products))
	for _, product := range productsResDto.Products {
		productsMap[product.ID] = product
	}

	orderItems := make([]model.OrderItem, len(orderDto.Items))
	for i, item := range orderDto.Items {
		orderItems[i] = model.OrderItem{
			ProductID: item.ProductID,
			Price:     productsMap[item.ProductID].Price,
			Quantity:  item.Quantity,
		}
	}

	orderAgg := model.OrderAggregate{
		Order: order,
		Items: orderItems,
	}
	// Create order in DB span
	span, ctx = apm.StartSpan(ctx, "CreateOrderDB", "database")
	orderId, err := uc.orderRepo.CreateOrder(ctx, orderAgg)
	span.End()
	
	if err != nil {
		log.Errorf("Create Order failed", err.Error())
		apm.CaptureError(ctx, err).Send()
		return nil, errors.New(util.ERR_INTERNAL_SERVER_ERROR)
	}

	return &dto.CreateOrderResponseDTO{
		OrderID: orderId,
	}, nil
}
