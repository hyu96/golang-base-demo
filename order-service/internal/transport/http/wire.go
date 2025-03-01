//go:build wireinject
// +build wireinject

package http_server

import (
	"github.com/google/wire"
	httpserver "github.com/huydq/gokits/libs/transport/http"
	"github.com/huydq/order-service/internal/adapter/grpc_client"
	"github.com/huydq/order-service/internal/adapter/persistence/postgres"
	"github.com/huydq/order-service/internal/core/port"
	"github.com/huydq/order-service/internal/core/repository"
	"github.com/huydq/order-service/internal/core/service"
	"github.com/huydq/order-service/internal/transport/http/handler"
	"github.com/huydq/order-service/internal/transport/http/server"

	"github.com/huydq/gokits/shared/core/biz/healthz"
)

func wireApp(httpServer *httpserver.HttpServer, orderPgClient postgres.OrderPostgresClient) (*server.OrderPublicServer, error) {
	panic(wire.Build(
		healthz.ProviderSet,
		server.ProviderSet,
		order_handler.ProviderSet,
		service.ProviderSet,
		postgres.NewOrderRepository,
		wire.Bind(new(repository.IOrderRepository), new(*postgres.OrderRepository)),
		grpc_client.NewGrpcProductMgmtServiceClient,
		wire.Bind(new(port.IProductServiceClient), new(*grpc_client.GrpcProductServiceClient)),
	))
}
