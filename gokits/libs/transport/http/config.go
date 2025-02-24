package httpserver

import (
	"os"

	"github.com/huydq/gokits/libs/env"
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Port string `json:"port,omitempty"`
}

var Config *HttpConfig

func setupHttpServerConfig() {
	Config = &HttpConfig{}
	if err := viper.UnmarshalKey("Server.Http", Config); err != nil {
		panic(err)
	}

	if Config.Port == "" {
		panic("http port empty!")
	}

	env.Config().Addr = Config.Port

	if viper.GetBool("APM.Active") {
		os.Setenv("ELASTIC_APM_SERVER_URL", viper.GetString("APM.ServerUrl"))
		os.Setenv("ELASTIC_APM_SERVICE_NAME", viper.GetString("APM.ServiceName"))
		os.Setenv("ELASTIC_APM_ENVIRONMENT", viper.GetString("APM.Environment"))
		os.Setenv("ELASTIC_APM_ACTIVE", viper.GetString("APM.Active"))
	}
}
