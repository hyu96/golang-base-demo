package grpc

import (
	"github.com/spf13/viper"

	"github.com/huydq/gokits/libs/ilog"
)

// GrpcConfig func
type GrpcConfig struct {
	RegionDC    string
	ServiceName string
	Addr        string
	Username    string
	Password    string
	Balancer    string
}

var Configs []GrpcConfig

func LoadGrpcClientConfig() {
	Configs = make([]GrpcConfig, 0)

	if err := viper.UnmarshalKey("Client.Grpc", &Configs); err != nil {
		ilog.Errorf("getGDiscoveryServerConfigFromEnv - Error: %v", err)
		panic(err)
	}

	if len(Configs) == 0 {
		panic("load grpc client config empty")
	}
}
