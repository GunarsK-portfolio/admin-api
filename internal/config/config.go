package config

import "os"

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	S3Endpoint   string
	S3AccessKey  string
	S3SecretKey  string
	S3Bucket     string
	S3UseSSL        string
	Port            string
	AuthServiceURL  string
}

func Load() *Config {
	return &Config{
		DBHost:         getEnvRequired("DB_HOST"),
		DBPort:         getEnvRequired("DB_PORT"),
		DBUser:         getEnvRequired("DB_USER"),
		DBPassword:     getEnvRequired("DB_PASSWORD"),
		DBName:         getEnvRequired("DB_NAME"),
		S3Endpoint:     getEnvRequired("S3_ENDPOINT"),
		S3AccessKey:    getEnvRequired("S3_ACCESS_KEY"),
		S3SecretKey:    getEnvRequired("S3_SECRET_KEY"),
		S3Bucket:       getEnvRequired("S3_BUCKET"),
		S3UseSSL:       getEnv("S3_USE_SSL", "false"),
		Port:           getEnv("PORT", "8083"),
		AuthServiceURL: getEnvRequired("AUTH_SERVICE_URL"),
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
