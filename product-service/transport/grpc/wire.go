//go:build wireinject
// +build wireinject

package grpc_server

import (
	"github.com/google/wire"
	"github.com/huydq/product-service/app/core/service"
	"github.com/huydq/product-service/transport/grpc/server"

	"github.com/huydq/gokits/libs/transport/grpc"
	product_handler "github.com/huydq/product-service/transport/grpc/handler"
)

func wireApp(rpcServer *grpc.RPCServer) (*server.ProductGrpcServer, error) {
	panic(wire.Build(
		server.ProviderSet,
		product_handler.ProviderSet,
		service.ProviderSet,
	))
}
