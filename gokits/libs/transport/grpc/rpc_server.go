package grpc

import (
	"net"
	"os"

	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"go.elastic.co/apm/v2"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/huydq/gokits/libs/auth/jwt/itoken"
	"github.com/huydq/gokits/libs/env"
	"github.com/huydq/gokits/libs/ilog"
)

type RPCServer struct {
	addr string
	s    *grpc.Server
}

func NewRPCServer() *RPCServer {
	loadGRPCConfig()
	itoken.GetJWTConfig()

	s := &RPCServer{
		addr: rpcConfig.Addr,
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(BizUnaryRecoveryHandler2),
	}

	os.Setenv("ELASTIC_APM_SERVER_URL", viper.GetString("APM.ServerUrl"))
	apmTracer, err := apm.NewTracer(env.Config().ServiceName, "")
	if err != nil {
		ilog.Errorf("RPCServer::Serve - Error failed to init apm tracer: %v", err)
		panic(err)
	}

	s.s = grpc.NewServer(
		// unary
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id")),
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithTracer(apmTracer)),
		),

		// stream
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id")),
			apmgrpc.NewStreamServerInterceptor(apmgrpc.WithTracer(apmTracer)),
		),
	)

	reflection.Register(s.s)
	return s
}

type RegisterRPCServerFunc func(s *grpc.Server)

func (s *RPCServer) Serve(regFunc RegisterRPCServerFunc) {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		ilog.Errorf("RPCServer::Serve - Error failed to listen: %v", err)
		return
	}

	ilog.Infof("rpc listening on:%s", s.addr)

	if regFunc != nil {
		regFunc(s.s)
	}

	defer s.s.GracefulStop()

	if err := s.s.Serve(listener); err != nil {
		ilog.Errorf("failed to serve: %s", err)
	}
}

func (s *RPCServer) Stop() {
	s.s.GracefulStop()
}

func (s *RPCServer) GetGRPCOriginServer() *grpc.Server {
	return s.s
}
