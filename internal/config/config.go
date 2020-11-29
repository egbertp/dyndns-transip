package config

import (
	"github.com/egbertp/dyndns-transip/internal/logger"
	"github.com/spf13/viper"
)

var _config *viper.Viper

// Init the config setup
func Init() {
	_config = viper.New()
	_config.SetConfigType("yaml")
	_config.AddConfigPath("/etc")
	_config.AddConfigPath(".")
	_config.SetConfigName("dyndns-transip")

	// ToDo: make sure that the version command stil works if the config file is
	// not successfully parsed
	if err := _config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Get().Fatalf("%s\n", err.Error())
		} else {
			logger.Get().Fatalf("Error parsing config file. (%s)\n", err.Error())
		}
	}
}

// Get the config instance
func Get() *viper.Viper {
	return _config
}
