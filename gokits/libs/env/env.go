package env

import (
	"github.com/spf13/viper"
)

var config *EnvConfig

type EnvConfig struct {
	Environment     string `json:"environment,omitempty"`
	Addr            string `json:"addr,omitempty"`
	DCName          string `json:"dc_name,omitempty"`
	HostName        string `json:"host_name,omitempty"`
	PodName         string `json:"pod_name,omitempty"`
	PodID           string `json:"pod_id,omitempty"`
	ServiceName     string `json:"service_name,omitempty"`
	ShutdownTimeout int    `json:"shutdown_timeout,omitempty"`
}

func SetupEnvConfig() {
	// Service config
	config = &EnvConfig{}
	if err := viper.UnmarshalKey("Env", config); err != nil {
		panic(err)
	}
}

func Config() *EnvConfig {
	return config
}
