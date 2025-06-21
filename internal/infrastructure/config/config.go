package config

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func New() (cfg *Config, err error) {
	setup()

	viper.ReadInConfig()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func setup() {
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	setDefaults()

	viper.SetConfigName("config")
	viper.AddConfigPath(path.Dir(os.Args[0]))
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
}
