package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	common "github.com/GunarsK-portfolio/portfolio-common/config"
)

type Config struct {
	DBHost         string `validate:"required"`
	DBPort         string `validate:"required,number,min=1,max=65535"`
	DBUser         string `validate:"required"`
	DBPassword     string `validate:"required"`
	DBName         string `validate:"required"`
	Port           string `validate:"required,number,min=1,max=65535"`
	AuthServiceURL string `validate:"required,url"`
	FilesAPIURL    string `validate:"required,url"`
}

func Load() *Config {
	cfg := &Config{
		DBHost:         common.GetEnvRequired("DB_HOST"),
		DBPort:         common.GetEnvRequired("DB_PORT"),
		DBUser:         common.GetEnvRequired("DB_USER"),
		DBPassword:     common.GetEnvRequired("DB_PASSWORD"),
		DBName:         common.GetEnvRequired("DB_NAME"),
		Port:           common.GetEnv("PORT", "8083"),
		AuthServiceURL: common.GetEnvRequired("AUTH_SERVICE_URL"),
		FilesAPIURL:    common.GetEnvRequired("FILES_API_URL"),
	}

	// Validate configuration
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	return cfg
}
