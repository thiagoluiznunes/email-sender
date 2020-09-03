package config

import (
	"strings"

	"github.com/spf13/viper"
)

func setup() {
	viper.SetEnvPrefix("api")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}

// Read returns the configuration values,
// based on the configuration files and environment variables.
func Read() (*Config, error) {
	setup()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
