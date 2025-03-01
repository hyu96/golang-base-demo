package repository

import (
	"context"
	"github.com/huydq/order-service/internal/core/domain/model"
)

//go:generate mockgen -source=index.go -destination=./mocks/index.go

type IOrderRepository interface {
	CreateOrder(ctx context.Context, orderAgg model.OrderAggregate) (int, error)
}
