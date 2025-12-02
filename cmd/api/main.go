package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/GunarsK-portfolio/admin-api/docs"
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/GunarsK-portfolio/admin-api/internal/routes"
	commondb "github.com/GunarsK-portfolio/portfolio-common/database"
	"github.com/GunarsK-portfolio/portfolio-common/health"
	"github.com/GunarsK-portfolio/portfolio-common/logger"
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	"github.com/GunarsK-portfolio/portfolio-common/server"
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

	//nolint:staticcheck // Embedded field name required due to ambiguous fields
	db, err := commondb.Connect(commondb.PostgresConfig{
		Host:     cfg.DatabaseConfig.Host,
		Port:     strconv.Itoa(cfg.DatabaseConfig.Port),
		User:     cfg.DatabaseConfig.User,
		Password: cfg.DatabaseConfig.Password,
		DBName:   cfg.DatabaseConfig.Name,
		SSLMode:  cfg.DatabaseConfig.SSLMode,
	})
	if err != nil {
		appLogger.Error("Failed to connect to database", "error", err)
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if closeErr := commondb.CloseDB(db); closeErr != nil {
			appLogger.Error("Failed to close database", "error", closeErr)
		}
	}()
	appLogger.Info("Database connection established")

	// Health checks
	healthAgg := health.NewAggregator(3 * time.Second)
	healthAgg.Register(health.NewPostgresChecker(db))

	repo := repository.New(db, cfg.FilesAPIURL)
	handler := handlers.New(repo)

	router := gin.New()
	router.Use(logger.Recovery(appLogger))
	router.Use(logger.RequestLogger(appLogger))
	router.Use(metricsCollector.Middleware())

	routes.Setup(router, handler, cfg, metricsCollector, healthAgg)

	appLogger.Info("Admin API ready", "port", cfg.ServiceConfig.Port, "environment", os.Getenv("ENVIRONMENT"))

	serverCfg := server.DefaultConfig(strconv.Itoa(cfg.ServiceConfig.Port))
	if err := server.Run(router, serverCfg, appLogger); err != nil {
		appLogger.Error("Server error", "error", err)
		log.Fatal("Server error:", err)
	}
}
