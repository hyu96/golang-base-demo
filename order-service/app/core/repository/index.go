package repository

import (
	"context"
	"github.com/google/wire"
	csql "github.com/huydq/gokits/libs/storage/pg-client"
	"github.com/huydq/order-service/app/core/models"
)

//go:generate mockgen -source=index.go -destination=../mocks/index.go

type IOrderRepository interface {
	CreateOrder(ctx context.Context, orderAgg models.OrderAggregate) (int, error)
}

type OrderRepository struct {
	orderClient csql.BasePostgresSqlxDB
}

type OrderPostgresClient struct {
	db csql.BasePostgresSqlxDB
}

func NewOrderPostgresClient(db csql.BasePostgresSqlxDB) *OrderPostgresClient {
	return &OrderPostgresClient{
		db: db,
	}
}

func NewOrderRepository(client OrderPostgresClient) *OrderRepository {
	return &OrderRepository{
		orderClient: client.db,
	}
}

var ProviderSet = wire.NewSet(NewOrderRepository, NewOrderPostgresClient)
