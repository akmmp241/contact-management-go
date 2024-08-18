package app

import "github.com/spf13/viper"

type Config struct {
	env *viper.Viper
}

func NewConfig() *Config {
	config := viper.New()
	config.SetConfigFile(".env")
	return &Config{env: config}
}
