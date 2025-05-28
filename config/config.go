// // Файл: digital-market/config/config.go
// package config

// import (
// 	"log"

// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	Server ServerConfig
// 	DB     DatabaseConfig
// 	JWT    JWTConfig
// 	Redis  RedisConfig
// }

// type ServerConfig struct {
// 	Port  string
// 	Debug bool
// }

// type DatabaseConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	Name     string
// 	SSLMode  string
// }

// type JWTConfig struct {
// 	Secret    string
// 	ExpiresIn string
// }

// type RedisConfig struct {
// 	Host     string
// 	Port     string
// 	Password string
// 	DB       int
// }

// func LoadConfig() *Config {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./config")
// 	viper.AddConfigPath(".")

// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalf("Error reading config file: %s", err)
// 	}

// 	var config Config
// 	if err := viper.Unmarshal(&config); err != nil {
// 		log.Fatalf("Unable to decode config into struct: %v", err)
// 	}

// 	return &config
// }


// Файл: config/config.go
package config

import (
	"log"
	"os"

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
	// Загрузка из файла конфигурации
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Загрузка конфигурации из файла (если есть)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %s", err)
	}

	// Настройка загрузки из переменных окружения
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	
	// Приоритет переменных окружения над файлом конфигурации
	var config Config
	
	// Сервер
	config.Server.Port = getEnvOrDefault("SERVER_PORT", viper.GetString("server.port"), "8080")
	config.Server.Debug = getEnvAsBoolOrDefault("SERVER_DEBUG", viper.GetBool("server.debug"), false)
	
	// База данных
	config.DB.Host = getEnvOrDefault("DB_HOST", viper.GetString("db.host"), "localhost")
	config.DB.Port = getEnvOrDefault("DB_PORT", viper.GetString("db.port"), "5432")
	config.DB.User = getEnvOrDefault("DB_USER", viper.GetString("db.user"), "postgres")
	config.DB.Password = getEnvOrDefault("DB_PASSWORD", viper.GetString("db.password"), "postgres")
	config.DB.Name = getEnvOrDefault("DB_NAME", viper.GetString("db.name"), "digital_market")
	config.DB.SSLMode = getEnvOrDefault("DB_SSLMODE", viper.GetString("db.sslmode"), "disable")
	
	// JWT
	config.JWT.Secret = getEnvOrDefault("JWT_SECRET", viper.GetString("jwt.secret"), "super_secret_key_change_in_production")
	config.JWT.ExpiresIn = getEnvOrDefault("JWT_EXPIRES_IN", viper.GetString("jwt.expires_in"), "24h")
	
	// Redis
	config.Redis.Host = getEnvOrDefault("REDIS_HOST", viper.GetString("redis.host"), "localhost")
	config.Redis.Port = getEnvOrDefault("REDIS_PORT", viper.GetString("redis.port"), "6379")
	config.Redis.Password = getEnvOrDefault("REDIS_PASSWORD", viper.GetString("redis.password"), "")
	config.Redis.DB = viper.GetInt("redis.db")
	
	return &config
}

func getEnvOrDefault(envKey, configValue, defaultValue string) string {
	if val, exists := os.LookupEnv(envKey); exists {
		return val
	}
	if configValue != "" {
		return configValue
	}
	return defaultValue
}

func getEnvAsBoolOrDefault(envKey string, configValue, defaultValue bool) bool {
	if val, exists := os.LookupEnv(envKey); exists {
		return val == "true" || val == "1" || val == "yes"
	}
	return configValue
}
