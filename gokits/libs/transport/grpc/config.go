package grpc

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config func;
type Config struct {
	Addr string `mapstructure:"port" json:"addr,omitempty"`
}

var rpcConfig *Config

func loadGRPCConfig() {
	if rpcConfig != nil {
		return
	}

	rpcConfig = &Config{}

	if err := viper.UnmarshalKey("Server.Grpc", rpcConfig); err != nil || rpcConfig.Addr == "" {
		err = fmt.Errorf("getGRPCConfig - Error: %v", err)
		panic(err)
	}
}
