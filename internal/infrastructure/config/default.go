package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("app.name", "delivery-manager-api")
	viper.SetDefault("app.environment", "local")
	viper.SetDefault("server.port", 5000)
}
