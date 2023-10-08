package config

import "github.com/spf13/viper"

// Init config files.
func Init() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
