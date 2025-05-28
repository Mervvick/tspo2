// Файл: digital-market/config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
	JWT    JWTConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port  string
	Debug bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret    string
	ExpiresIn string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &config
}
