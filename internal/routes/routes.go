package routes

import (
	"log"

	"github.com/GunarsK-portfolio/admin-api/docs"
	"github.com/GunarsK-portfolio/admin-api/internal/config"
	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/portfolio-common/health"
	"github.com/GunarsK-portfolio/portfolio-common/jwt"
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	common "github.com/GunarsK-portfolio/portfolio-common/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, handler *handlers.Handler, cfg *config.Config, metricsCollector *metrics.Metrics, healthAgg *health.Aggregator) {
	// Security middleware with CORS validation
	securityMiddleware := common.NewSecurityMiddleware(
		cfg.AllowedOrigins,
		"GET,POST,PUT,DELETE,OPTIONS",
		"Content-Type,Authorization",
		true,
	)
	router.Use(securityMiddleware.Apply())

	// Health check
	router.GET("/health", healthAgg.Handler())

	// Metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Protected API routes - JWT validation
	jwtService, err := jwt.NewValidatorOnly(cfg.JWTSecret)
	if err != nil {
		log.Fatalf("Failed to create JWT service: %v", err)
	}
	authMiddleware := common.NewAuthMiddleware(jwtService)
	v1 := router.Group("/api/v1")
	v1.Use(authMiddleware.ValidateToken())
	v1.Use(authMiddleware.AddTTLHeader()) // Add TTL header to all responses
	{
		// Portfolio domain
		portfolio := v1.Group("/portfolio")
		{
			// Profile
			portfolio.GET("/profile", common.RequirePermission(common.ResourceProfile, common.LevelRead), handler.GetProfile)
			portfolio.PUT("/profile", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfile)
			portfolio.PUT("/profile/avatar", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfileAvatar)
			portfolio.DELETE("/profile/avatar", common.RequirePermission(common.ResourceProfile, common.LevelDelete), handler.DeleteProfileAvatar)
			portfolio.PUT("/profile/resume", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfileResume)
			portfolio.DELETE("/profile/resume", common.RequirePermission(common.ResourceProfile, common.LevelDelete), handler.DeleteProfileResume)

			// Work Experience
			portfolio.GET("/experience", common.RequirePermission(common.ResourceExperience, common.LevelRead), handler.GetAllWorkExperience)
			portfolio.POST("/experience", common.RequirePermission(common.ResourceExperience, common.LevelEdit), handler.CreateWorkExperience)
			portfolio.GET("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelRead), handler.GetWorkExperienceByID)
			portfolio.PUT("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelEdit), handler.UpdateWorkExperience)
			portfolio.DELETE("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelDelete), handler.DeleteWorkExperience)

			// Certifications
			portfolio.GET("/certifications", common.RequirePermission(common.ResourceCertifications, common.LevelRead), handler.GetAllCertifications)
			portfolio.POST("/certifications", common.RequirePermission(common.ResourceCertifications, common.LevelEdit), handler.CreateCertification)
			portfolio.GET("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelRead), handler.GetCertificationByID)
			portfolio.PUT("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelEdit), handler.UpdateCertification)
			portfolio.DELETE("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelDelete), handler.DeleteCertification)

			// Skills
			portfolio.GET("/skills", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetAllSkills)
			portfolio.POST("/skills", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.CreateSkill)
			portfolio.GET("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetSkillByID)
			portfolio.PUT("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.UpdateSkill)
			portfolio.DELETE("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelDelete), handler.DeleteSkill)

			// Skill Types
			portfolio.GET("/skill-types", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetAllSkillTypes)
			portfolio.POST("/skill-types", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.CreateSkillType)
			portfolio.GET("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetSkillTypeByID)
			portfolio.PUT("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.UpdateSkillType)
			portfolio.DELETE("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelDelete), handler.DeleteSkillType)

			// Portfolio Projects
			portfolio.GET("/projects", common.RequirePermission(common.ResourceProjects, common.LevelRead), handler.GetAllPortfolioProjects)
			portfolio.POST("/projects", common.RequirePermission(common.ResourceProjects, common.LevelEdit), handler.CreatePortfolioProject)
			portfolio.GET("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelRead), handler.GetPortfolioProjectByID)
			portfolio.PUT("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelEdit), handler.UpdatePortfolioProject)
			portfolio.DELETE("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelDelete), handler.DeletePortfolioProject)
		}

		// Miniatures domain
		miniatures := v1.Group("/miniatures")
		{
			// Miniature Themes
			miniatures.GET("/themes", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniatureThemes)
			miniatures.POST("/themes", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniatureTheme)
			miniatures.GET("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniatureThemeByID)
			miniatures.PUT("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniatureTheme)
			miniatures.DELETE("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniatureTheme)

			// Miniature Projects
			miniatures.GET("/projects", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniatureProjects)
			miniatures.POST("/projects", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniatureProject)
			miniatures.GET("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniatureProjectByID)
			miniatures.PUT("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniatureProject)
			miniatures.DELETE("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniatureProject)
			miniatures.POST("/projects/:id/images", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.AddImageToProject)
			miniatures.PUT("/projects/:id/techniques", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.SetProjectTechniques)
			miniatures.PUT("/projects/:id/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.SetProjectPaints)

			// Miniature Techniques
			miniatures.GET("/techniques", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllTechniques)

			// Miniature Paints
			miniatures.GET("/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniaturePaints)
			miniatures.POST("/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniaturePaint)
			miniatures.GET("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniaturePaintByID)
			miniatures.PUT("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniaturePaint)
			miniatures.DELETE("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniaturePaint)
		}

		// Files (generic file deletion - requires delete permission on files resource)
		v1.DELETE("/files/:id", common.RequirePermission(common.ResourceFiles, common.LevelDelete), handler.DeleteImage)
	}

	// Swagger documentation (only if SWAGGER_HOST is configured)
	if cfg.SwaggerHost != "" {
		docs.SwaggerInfo.Host = cfg.SwaggerHost
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
