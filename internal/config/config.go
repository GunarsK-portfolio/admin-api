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
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "portfolio_user"),
		DBPassword:     getEnv("DB_PASSWORD", "portfolio_pass"),
		DBName:         getEnv("DB_NAME", "portfolio"),
		S3Endpoint:     getEnv("S3_ENDPOINT", "http://localhost:9000"),
		S3AccessKey:    getEnv("S3_ACCESS_KEY", "minioadmin"),
		S3SecretKey:    getEnv("S3_SECRET_KEY", "minioadmin"),
		S3Bucket:       getEnv("S3_BUCKET", "portfolio-images"),
		S3UseSSL:       getEnv("S3_USE_SSL", "false"),
		Port:           getEnv("PORT", "8083"),
		AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8084"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
