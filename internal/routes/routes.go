package routes

import (
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	common "github.com/GunarsK-portfolio/portfolio-common/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, handler *handlers.Handler, cfg *config.Config, metricsCollector *metrics.Metrics) {
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

	// Health check
	router.GET("/health", handler.HealthCheck)

	// Metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Protected API routes
	authMiddleware := common.NewAuthMiddleware(cfg.AuthServiceURL)
	v1 := router.Group("/api/v1")
	v1.Use(authMiddleware.ValidateToken())
	v1.Use(authMiddleware.AddTTLHeader()) // Add TTL header to all responses
	{
		// Portfolio domain
		portfolio := v1.Group("/portfolio")
		{
			// Profile
			portfolio.GET("/profile", handler.GetProfile)
			portfolio.PUT("/profile", handler.UpdateProfile)
			portfolio.PUT("/profile/avatar", handler.UpdateProfileAvatar)
			portfolio.DELETE("/profile/avatar", handler.DeleteProfileAvatar)
			portfolio.PUT("/profile/resume", handler.UpdateProfileResume)
			portfolio.DELETE("/profile/resume", handler.DeleteProfileResume)

			// Work Experience
			portfolio.GET("/experience", handler.GetAllWorkExperience)
			portfolio.POST("/experience", handler.CreateWorkExperience)
			portfolio.GET("/experience/:id", handler.GetWorkExperienceByID)
			portfolio.PUT("/experience/:id", handler.UpdateWorkExperience)
			portfolio.DELETE("/experience/:id", handler.DeleteWorkExperience)

			// Certifications
			portfolio.GET("/certifications", handler.GetAllCertifications)
			portfolio.POST("/certifications", handler.CreateCertification)
			portfolio.GET("/certifications/:id", handler.GetCertificationByID)
			portfolio.PUT("/certifications/:id", handler.UpdateCertification)
			portfolio.DELETE("/certifications/:id", handler.DeleteCertification)

			// Skills
			portfolio.GET("/skills", handler.GetAllSkills)
			portfolio.POST("/skills", handler.CreateSkill)
			portfolio.GET("/skills/:id", handler.GetSkillByID)
			portfolio.PUT("/skills/:id", handler.UpdateSkill)
			portfolio.DELETE("/skills/:id", handler.DeleteSkill)

			// Skill Types
			portfolio.GET("/skill-types", handler.GetAllSkillTypes)
			portfolio.POST("/skill-types", handler.CreateSkillType)
			portfolio.GET("/skill-types/:id", handler.GetSkillTypeByID)
			portfolio.PUT("/skill-types/:id", handler.UpdateSkillType)
			portfolio.DELETE("/skill-types/:id", handler.DeleteSkillType)

			// Portfolio Projects
			portfolio.GET("/projects", handler.GetAllPortfolioProjects)
			portfolio.POST("/projects", handler.CreatePortfolioProject)
			portfolio.GET("/projects/:id", handler.GetPortfolioProjectByID)
			portfolio.PUT("/projects/:id", handler.UpdatePortfolioProject)
			portfolio.DELETE("/projects/:id", handler.DeletePortfolioProject)
			// TODO: Add portfolio.PUT("/projects/:id/image", handler.UpdatePortfolioProjectImage)
			// TODO: Add portfolio.DELETE("/projects/:id/image", handler.DeletePortfolioProjectImage)
		}

		// Miniatures domain
		miniatures := v1.Group("/miniatures")
		{
			// Miniature Themes
			miniatures.GET("/themes", handler.GetAllMiniatureThemes)
			miniatures.POST("/themes", handler.CreateMiniatureTheme)
			miniatures.GET("/themes/:id", handler.GetMiniatureThemeByID)
			miniatures.PUT("/themes/:id", handler.UpdateMiniatureTheme)
			miniatures.DELETE("/themes/:id", handler.DeleteMiniatureTheme)

			// Miniature Projects
			miniatures.GET("/projects", handler.GetAllMiniatureProjects)
			miniatures.POST("/projects", handler.CreateMiniatureProject)
			miniatures.GET("/projects/:id", handler.GetMiniatureProjectByID)
			miniatures.PUT("/projects/:id", handler.UpdateMiniatureProject)
			miniatures.DELETE("/projects/:id", handler.DeleteMiniatureProject)
			// TODO: Add miniatures.GET("/projects/:id/images", handler.GetMiniatureProjectImages)
			// TODO: Add miniatures.POST("/projects/:id/images", handler.AddMiniatureProjectImage)

			// Miniature Paints
			miniatures.GET("/paints", handler.GetAllMiniaturePaints)
			miniatures.POST("/paints", handler.CreateMiniaturePaint)
			miniatures.GET("/paints/:id", handler.GetMiniaturePaintByID)
			miniatures.PUT("/paints/:id", handler.UpdateMiniaturePaint)
			miniatures.DELETE("/paints/:id", handler.DeleteMiniaturePaint)
		}

		// Files (generic file deletion for all file types)
		v1.DELETE("/files/:id", handler.DeleteImage)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
