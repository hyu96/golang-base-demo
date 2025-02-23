package csql

import (
	"fmt"
	"strings"
	"time"

	"github.com/huydq/gokits/libs/env"
	"github.com/spf13/viper"
)

var (
	Name        string
	Environment string
	DSN         string
	Active      int
	Idle        int
	LifeTime    int // In seconds
)

const (
	KDefaultTimeout = 30 * time.Second

	DB_ORDER_SERVICE = "order_service"
)

type SQLConfig struct {
	Name     string `json:"name,omitempty"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname""`
	SSLMode  string `json:"sslmode"`
	DSN      string `json:"dsn,omitempty"`

	Driver          string `json:"driver,omitempty"` // can be postgres but default is mysql
	Environment     string `json:"environment,omitempty"`
	MaxOpenConns    int    `json:"active,omitempty"`
	MaxIdleConns    int    `json:"idle,omitempty"`
	ConnMaxLifetime int    `json:"lifetime,omitempty"` // Connection's lifetime in seconds
	ConnMaxIdleTime int    `json:"maxIdleTime,omitempty"`
}

// default value env key is "MySQL";
// if configKeys was set, key env will be first value (not empty) of this;
func getConfigFromEnv(configKeys ...string) []*SQLConfig {

	configKey := "Postgres"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	raw := make([]*SQLConfig, 0)

	if err := viper.UnmarshalKey(configKey, &raw); err != nil {
		err := fmt.Errorf("not found config name with env %q for SQL with error: %+v", configKey, err)
		panic(err)
	}

	sqlConfigs := make([]*SQLConfig, len(configKeys))
	for _, config := range raw {
		if config.Name == "" {
			config.Name = "master"
		}

		if config.Environment == "" {
			config.Environment = env.Config().Environment
		}

		if config.MaxOpenConns == 0 {
			config.MaxOpenConns = 50
		}

		if config.MaxIdleConns == 0 {
			config.MaxIdleConns = 50
		}

		if config.ConnMaxLifetime == 0 {
			config.ConnMaxLifetime = 5 * 60
		}

		if config.ConnMaxLifetime == 0 {
			config.ConnMaxLifetime = 5 * 60
		}

		if config.Driver == "" {
			config.Driver = "postgres"
		}

		dsn := fmt.Sprintf("host=%s port=%s user='%s' password='%s' dbname='%s' sslmode=%s",
			config.Host,
			config.Port,
			config.Username,
			config.Password,
			config.DBName,
			config.SSLMode) // Convert seconds to milliseconds
		config.DSN = dsn

		sqlConfigs = append(sqlConfigs, config)
	}

	if len(sqlConfigs) == 0 {
		err := fmt.Errorf("not found valid config with env %q for SQL", configKey)
		panic(err)
	}

	return sqlConfigs
}
