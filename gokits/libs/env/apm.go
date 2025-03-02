package env

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"os"
)

type APMConfig struct {
	ServerURL   string
	ServiceName string
	Environment string
	LogLevel    string
}

func LoadAPMConfig() {
	// Service config
	apmConfig := &APMConfig{}
	if err := viper.UnmarshalKey("APM", apmConfig); err != nil {
		log.Warn("APM Config is missing")
	}

	os.Setenv("ELASTIC_APM_SERVER_URL", apmConfig.ServerURL)
	os.Setenv("ELASTIC_APM_SERVICE_NAME", apmConfig.ServiceName)
	os.Setenv("ELASTIC_APM_ENVIRONMENT", apmConfig.Environment)
	os.Setenv("ELASTIC_APM_LOG_LEVEL", apmConfig.LogLevel)
}
