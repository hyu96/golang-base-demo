package ilog

import (
	"github.com/spf13/viper"
)

var config *LogConfig

type LogConfig struct {
	Development     bool   `json:"development,omitempty"`
	LogDir          string `json:"log_dir,omitempty"`
	LogFileLevel    string `json:"log_file_level,omitempty"`
	LogConsoleLevel string `json:"log_console_level,omitempty"`
	Description     string `json:"description,omitempty"`
}

func setupLogConfig() {
	config = &LogConfig{}
	if err := viper.UnmarshalKey("Log", config); err != nil {
		panic(err)
	}

	if config.LogDir == "" {
		panic("Log.log_dirs should not be empty!")
	}
}

func Config() *LogConfig {
	return config
}
