package main

import (
	"fmt"
	"log"

	_ "github.com/GunarsK-portfolio/admin-api/docs"
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/GunarsK-portfolio/admin-api/internal/routes"
	commondb "github.com/GunarsK-portfolio/portfolio-common/database"
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
	// Load configuration
	cfg := config.Load()

	// Connect to database
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
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository
	repo := repository.New(db, cfg.FilesAPIURL)

	// Initialize handlers
	handler := handlers.New(repo)

	// Setup router
	router := gin.Default()

	// Setup routes
	routes.Setup(router, handler, cfg)

	// Start server
	log.Printf("Starting admin API on port %s", cfg.Port)
	if err := router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
