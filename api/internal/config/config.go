package config

import (
	"github.com/spf13/viper"
)

// Config defines static configuration for app
type Config struct {
	// DB
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBPort int `mapstructure:"POSTGRES_PORT"`
	DBName string `mapstructure:"POSTGRES_DBNAME"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPass string `mapstructure:"POSTGRES_PASS"`

	// HTTP Server
	ServerPort string `mapstructure:"SERVER_PORT"`
}

// Load reads config from file or environment and return Config struct
func Load(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}