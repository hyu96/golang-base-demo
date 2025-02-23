package itoken

import (
	"fmt"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	SecretKey   string `json:"SecretKey,omitempty"`
	RedisPrefix string `json:"RedisPrefix,omitempty"`
}

var conf *JWTConfig

func GetJWTConfig() *JWTConfig {
	if conf != nil {
		return conf
	}

	conf = &JWTConfig{}
	if err := viper.UnmarshalKey("JWT", conf); err != nil {
		panic(fmt.Errorf("not found config name with env %q for OpenIDJwt with error: %+v", "OpenIDJwt", err))
	}

	return conf
}
