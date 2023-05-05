package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Post PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func New() Config {
	return Config{
		Post: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			Username: viper.GetString("postgres.username"),
			Password: viper.GetString("postgres.password"),
			DBName:   viper.GetString("postgres.dbname"),
			SSLMode:  viper.GetString("postgres.sslmode"),
		},
	}
}
