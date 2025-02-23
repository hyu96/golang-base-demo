package server

import (
	"github.com/google/wire"
	httpserver "github.com/huydq/gokits/libs/transport/http"
	"github.com/huydq/gokits/shared/core/biz/healthz"
	order_handler "github.com/huydq/order-service/transport/http/handler"
)

type OrderPublicServer struct {
	httpServer *httpserver.HttpServer

	healthz      *healthz.HealthZBiz
	orderHandler *order_handler.OrderHandler
}

func NewOrderPublicServer(
	httpServer *httpserver.HttpServer,
	healthzBiz *healthz.HealthZBiz,
	orderHandler *order_handler.OrderHandler,
) *OrderPublicServer {
	return &OrderPublicServer{
		httpServer:   httpServer,
		healthz:      healthzBiz,
		orderHandler: orderHandler,
	}
}

func (s *OrderPublicServer) Start() {
	s.setupHttpRoute()
	go s.httpServer.Serve()
}

func (s *OrderPublicServer) Stop() {
	s.httpServer.Stop()
}

var ProviderSet = wire.NewSet(NewOrderPublicServer)
