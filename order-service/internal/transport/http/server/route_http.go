package server

func (s *OrderPublicServer) setupHttpRoute() {
	apiGroup := s.httpServer.Group("/")
	orderGroup := apiGroup.Group("orders")

	orderGroup.Post("", s.orderHandler.CreateOrder)
}
