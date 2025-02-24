//go:build wireinject
// +build wireinject

package http_server

import (
	"github.com/google/wire"
	"github.com/huydq/order-service/app/core/repository"
	"github.com/huydq/order-service/app/core/service"
	order_handler "github.com/huydq/order-service/transport/http/handler"
	"github.com/huydq/order-service/transport/http/server"

	httpserver "github.com/huydq/gokits/libs/transport/http"

	"github.com/huydq/gokits/shared/core/biz/healthz"
)

func wireApp(httpServer *httpserver.HttpServer, orderClient repository.OrderPostgresClient) (*server.OrderPublicServer, error) {
	panic(wire.Build(
		healthz.ProviderSet,
		server.ProviderSet,
		order_handler.ProviderSet,
		service.ProviderSet,
		repository.NewOrderRepository,
		//wire.Struct(new(repository.OrderRepository)),
		wire.Bind(new(repository.IOrderRepository), new(*repository.OrderRepository)),
	))
}
