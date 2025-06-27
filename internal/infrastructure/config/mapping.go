package config

type Config struct {
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
}

type App struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
}

type Server struct {
	Port int `mapstructure:"port"`
}
