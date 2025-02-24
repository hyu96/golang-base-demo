package grpc

import (
	"context"
	"fmt"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/balancer/weightedroundrobin"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/huydq/gokits/libs/ilog"
)

const (
	rls                   = "rls"
	grpcLB                = "grpclb"
	pickFirstBalancerName = "pick_first"
	roundRobin            = "round_robin"
	weightedRoundRobin    = "weighted_round_robin"
	// consistentHash        = "consistent_hash"
)

// NewRPCClientByServiceDiscovery func
func NewRPCClientByServiceDiscovery(config *GrpcConfig) (c *grpc.ClientConn, err error) {
	balancer := roundrobin.Name

	switch config.Balancer {
	case rls:
		balancer = rls
	case grpcLB:
		balancer = grpcLB
	case pickFirstBalancerName:
		balancer = grpc.PickFirstBalancerName
	case weightedRoundRobin:
		balancer = weightedroundrobin.Name
		// case consistentHash:
		// b = g_loadbalancer.NewKBalancer(r, g_loadbalancer.NewKetamaSelector(g_loadbalancer.DefaultKetamaKey))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	c, err = grpc.DialContext(
		ctx,
		config.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, balancer)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(5*1024*1024)),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id"))),
		grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id"))),
	)
	if err != nil {
		ilog.Errorf("NewRPCClientByServiceDiscovery - Error: %+v", err)
		panic(err)
	}
	return
}

func NewRPCClientExtraByServiceDiscovery(discovery *GrpcConfig) (c *grpc.ClientConn, err error) {
	balancer := roundrobin.Name

	switch discovery.Balancer {
	case rls:
		balancer = rls
	case grpcLB:
		balancer = grpcLB
	case pickFirstBalancerName:
		balancer = grpc.PickFirstBalancerName
	case weightedRoundRobin:
		balancer = weightedroundrobin.Name
		// case consistentHash:
		// b = g_loadbalancer.NewKBalancer(r, g_loadbalancer.NewKetamaSelector(g_loadbalancer.DefaultKetamaKey))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	c, err = grpc.DialContext(
		ctx,
		discovery.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, balancer)),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id"))),
		grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTraceHeaderName("icom-tracer-id"))),
	)

	if err != nil {
		ilog.Errorf("NewRPCClientExtraByServiceDiscovery - Error: %+v", err)
		panic(err)
	}

	return
}

// RPCClient type

var mapRPCClient map[string]*grpc.ClientConn

// NewRPCClient func
func InstallRPCClient() {
	LoadGrpcClientConfig()
	mapRPCClient = map[string]*grpc.ClientConn{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for _, config := range Configs {
		conn, err := grpc.DialContext(ctx, config.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("cannot connect grpc err: %s", err))
		}
		mapRPCClient[config.ServiceName] = conn
		ilog.Infof("[=] gprc client conn done: %s, state: %s", config.ServiceName, conn.GetState().String())
	}
}

func GetClientConn(name string) *grpc.ClientConn {
	if conn, ok := mapRPCClient[name]; ok {
		return conn
	}

	ilog.Panicf("get grpc Client Connection fail, name: %s", name)
	return nil
}
