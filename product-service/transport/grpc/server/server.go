package server

import (
	"github.com/google/wire"
	grpcserver "github.com/huydq/gokits/libs/transport/grpc"
	"github.com/huydq/gokits/shared/core/biz/healthz"

	productGrpcController "github.com/huydq/product-service/transport/grpc/handler"
)

type ProductGrpcServer struct {
	grpcServer         *grpcserver.RPCServer
	healthz            *healthz.HealthZBiz
	productGrpcHandler *productGrpcController.ProductGrpcHandler
}

func NewProductGrpcServer(
	grpcServer *grpcserver.RPCServer,
	productGrpcHandler *productGrpcController.ProductGrpcHandler,
) *ProductGrpcServer {

	return &ProductGrpcServer{
		grpcServer:         grpcServer,
		productGrpcHandler: productGrpcHandler,
	}
}

func (s *ProductGrpcServer) Start() {
	go s.grpcServer.Serve(s.RegisterGrpcService)
}

func (s *ProductGrpcServer) Stop() {
	s.grpcServer.Stop()
}

var ProviderSet = wire.NewSet(NewProductGrpcServer)
