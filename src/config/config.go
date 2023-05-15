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
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func New() Config {
	return Config{
		Post: PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: "2002",
			//Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:  viper.GetString("db.dbname"),
			SSLMode: viper.GetString("db.sslmode"),
		},
	}
}
