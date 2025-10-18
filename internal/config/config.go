package config

import "os"

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	Port           string
	AuthServiceURL string
	FilesAPIURL    string
}

func Load() *Config {
	return &Config{
		DBHost:         getEnvRequired("DB_HOST"),
		DBPort:         getEnvRequired("DB_PORT"),
		DBUser:         getEnvRequired("DB_USER"),
		DBPassword:     getEnvRequired("DB_PASSWORD"),
		DBName:         getEnvRequired("DB_NAME"),
		Port:           getEnv("PORT", "8083"),
		AuthServiceURL: getEnvRequired("AUTH_SERVICE_URL"),
		FilesAPIURL:    getEnvRequired("FILES_API_URL"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Required environment variable " + key + " is not set")
	}
	return value
}
