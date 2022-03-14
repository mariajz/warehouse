
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ShabbyInternetBaseUrl string `mapstructure:"SHABBY_INTERNET_BASE_URL"`

	LogLevel string `mapstructure:"LOG_LEVEL"`

	DBConfig `mapstructure:",squash"`
}

func LoadConfig() *AppConfig {
	var appConfig AppConfig
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Errorf("Failed to load config")
	}
	return &appConfig
}
