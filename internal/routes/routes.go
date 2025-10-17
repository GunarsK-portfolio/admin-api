package routes

import (
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	common "github.com/GunarsK-portfolio/portfolio-common/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, handler *handlers.Handler, cfg *config.Config) {
	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check (no auth)
	router.GET("/api/v1/health", handler.HealthCheck)

	// Protected API routes
	authMiddleware := common.NewAuthMiddleware(cfg.AuthServiceURL)
	v1 := router.Group("/api/v1")
	v1.Use(authMiddleware.ValidateToken())
	{
		// Profile
		v1.POST("/profile", handler.UpdateProfile)

		// Work Experience
		v1.POST("/experience", handler.CreateWorkExperience)
		v1.PUT("/experience/:id", handler.UpdateWorkExperience)
		v1.DELETE("/experience/:id", handler.DeleteWorkExperience)

		// Certifications
		v1.POST("/certifications", handler.CreateCertification)
		v1.PUT("/certifications/:id", handler.UpdateCertification)
		v1.DELETE("/certifications/:id", handler.DeleteCertification)

		// Miniature Projects
		v1.POST("/miniatures", handler.CreateMiniatureProject)
		v1.PUT("/miniatures/:id", handler.UpdateMiniatureProject)
		v1.DELETE("/miniatures/:id", handler.DeleteMiniatureProject)

		// Images
		v1.DELETE("/images/:id", handler.DeleteImage)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
