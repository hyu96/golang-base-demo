package gorm

import (
	"fmt"

	"github.com/huydq/gokits/libs/env"
	"github.com/spf13/viper"
)

// GormConfig func
type GormConfig struct {
	Name        string `json:"name,omitempty"`
	Driver      string `json:"driver,omitempty"` // can be postgres but default is mysql
	Environment string `json:"environment,omitempty"`
	DSN         string `json:"dsn,omitempty"`
	Active      int    `json:"active,omitempty"`
	Idle        int    `json:"idle,omitempty"`
	Lifetime    int    `json:"lifetime,omitempty"` // Connection's lifetime in seconds
}

var configs []GormConfig

func LoadGormClientConfig() {
	configKey := "MySQL"

	raw := make([]GormConfig, 0)

	if err := viper.UnmarshalKey(configKey, &raw); err != nil {
		err := fmt.Errorf("not found config name with env %q for SQL with error: %+v", configKey, err)
		panic(err)
	}

	configs = make([]GormConfig, 0)
	for _, config := range raw {
		if config.DSN == "" {
			continue
		}

		if config.Name == "" {
			config.Name = "immaster"
		}

		if config.Environment == "" {
			config.Environment = env.Config().Environment
		}

		if config.Active == 0 {
			config.Active = 50
		}

		if config.Idle == 0 {
			config.Idle = 50
		}

		if config.Lifetime == 0 {
			config.Lifetime = 5 * 60
		}

		if config.Driver == "" {
			config.Driver = "mysql"
		}

		configs = append(configs, config)
	}

	if len(configs) == 0 {
		err := fmt.Errorf("not found valid config with env %q for SQL", configKey)
		panic(err)
	}
}
