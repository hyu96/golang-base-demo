package elastic

import (
	"github.com/spf13/viper"

	"github.com/huydq/gokits/libs/ilog"
)

// ElasticConfig func
type ElasticConfig struct {
	ElasUser string
	ElasPass string
	ElasHost string
	ElasPort string
}

var Configs ElasticConfig

func LoadGrpcClientConfig() {
	if err := viper.UnmarshalKey("Client.Elastic", &Configs); err != nil {
		ilog.Errorf("getGDiscoveryServerConfigFromEnv - Error: %v", err)
		panic(err)
	}
	if Configs.ElasUser == "" {
		panic("load elastic client config empty")
	}
}
