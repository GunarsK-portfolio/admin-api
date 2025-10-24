package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GunarsK-portfolio/admin-api/docs"
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/GunarsK-portfolio/admin-api/internal/routes"
	commondb "github.com/GunarsK-portfolio/portfolio-common/database"
	"github.com/GunarsK-portfolio/portfolio-common/logger"
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	"github.com/gin-gonic/gin"
)

// @title Portfolio Admin API
// @version 1.0
// @description Admin API for managing portfolio content
// @host localhost:8083
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	appLogger := logger.New(logger.Config{
		Level:       os.Getenv("LOG_LEVEL"),
		Format:      os.Getenv("LOG_FORMAT"),
		ServiceName: "admin-api",
		AddSource:   os.Getenv("LOG_SOURCE") == "true",
	})

	appLogger.Info("Starting admin API", "version", "1.0")

	metricsCollector := metrics.New(metrics.Config{
		ServiceName: "admin",
		Namespace:   "portfolio",
	})

	db, err := commondb.Connect(commondb.PostgresConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  "disable",
		TimeZone: "UTC",
	})
	if err != nil {
		appLogger.Error("Failed to connect to database", "error", err)
		log.Fatal("Failed to connect to database:", err)
	}
	appLogger.Info("Database connection established")

	repo := repository.New(db, cfg.FilesAPIURL)
	handler := handlers.New(repo)

	router := gin.New()
	router.Use(logger.Recovery(appLogger))
	router.Use(logger.RequestLogger(appLogger))
	router.Use(metricsCollector.Middleware())

	routes.Setup(router, handler, cfg, metricsCollector)

	appLogger.Info("Admin API ready", "port", cfg.Port, "environment", os.Getenv("ENVIRONMENT"))
	if err := router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		appLogger.Error("Failed to start server", "error", err)
		log.Fatal("Failed to start server:", err)
	}
}
