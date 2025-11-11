package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	common "github.com/GunarsK-portfolio/portfolio-common/config"
)

type Config struct {
	common.DatabaseConfig
	common.ServiceConfig
	AuthServiceURL string `validate:"required,url"`
	FilesAPIURL    string `validate:"required,url"`
}

func Load() *Config {
	cfg := &Config{
		DatabaseConfig: common.NewDatabaseConfig(),
		ServiceConfig:  common.NewServiceConfig("8083"),
		AuthServiceURL: common.GetEnvRequired("AUTH_SERVICE_URL"),
		FilesAPIURL:    common.GetEnvRequired("FILES_API_URL"),
	}

	// Validate service-specific fields
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	return cfg
}
