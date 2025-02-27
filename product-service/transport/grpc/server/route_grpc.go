package server

import (
	pbProductMgmt "github.com/huydq/proto/gen-go/product"
	"google.golang.org/grpc"
)

func (s *ProductGrpcServer) RegisterGrpcService(server *grpc.Server) {
	pbProductMgmt.RegisterProductManagementServiceServer(server, s.productGrpcHandler)
}
