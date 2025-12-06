package routes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GunarsK-portfolio/admin-api/internal/handlers"
	"github.com/GunarsK-portfolio/admin-api/internal/models"
	common "github.com/GunarsK-portfolio/portfolio-common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// =============================================================================
// Mock Repository
// =============================================================================

// mockRepository implements handlers.Repository for route-level RBAC testing.
// Uses function fields to allow per-test behavior customization.
type mockRepository struct {
	// Profile
	getProfileFunc          func(ctx context.Context) (*models.Profile, error)
	updateProfileFunc       func(ctx context.Context, profile *models.Profile) error
	updateProfileAvatarFunc func(ctx context.Context, fileID int64) error
	deleteProfileAvatarFunc func(ctx context.Context) error
	updateProfileResumeFunc func(ctx context.Context, fileID int64) error
	deleteProfileResumeFunc func(ctx context.Context) error

	// Work Experience
	getAllWorkExperienceFunc  func(ctx context.Context) ([]models.WorkExperience, error)
	getWorkExperienceByIDFunc func(ctx context.Context, id int64) (*models.WorkExperience, error)
	createWorkExperienceFunc  func(ctx context.Context, exp *models.WorkExperience) error
	updateWorkExperienceFunc  func(ctx context.Context, exp *models.WorkExperience) error
	deleteWorkExperienceFunc  func(ctx context.Context, id int64) error

	// Certifications
	getAllCertificationsFunc func(ctx context.Context) ([]models.Certification, error)
	getCertificationByIDFunc func(ctx context.Context, id int64) (*models.Certification, error)
	createCertificationFunc  func(ctx context.Context, cert *models.Certification) error
	updateCertificationFunc  func(ctx context.Context, cert *models.Certification) error
	deleteCertificationFunc  func(ctx context.Context, id int64) error

	// Skills
	getAllSkillsFunc func(ctx context.Context) ([]models.Skill, error)
	getSkillByIDFunc func(ctx context.Context, id int64) (*models.Skill, error)
	createSkillFunc  func(ctx context.Context, skill *models.Skill) error
	updateSkillFunc  func(ctx context.Context, skill *models.Skill) error
	deleteSkillFunc  func(ctx context.Context, id int64) error

	// Skill Types
	getAllSkillTypesFunc func(ctx context.Context) ([]models.SkillType, error)
	getSkillTypeByIDFunc func(ctx context.Context, id int64) (*models.SkillType, error)
	createSkillTypeFunc  func(ctx context.Context, skillType *models.SkillType) error
	updateSkillTypeFunc  func(ctx context.Context, skillType *models.SkillType) error
	deleteSkillTypeFunc  func(ctx context.Context, id int64) error

	// Portfolio Projects
	getAllPortfolioProjectsFunc func(ctx context.Context) ([]models.PortfolioProject, error)
	getPortfolioProjectByIDFunc func(ctx context.Context, id int64) (*models.PortfolioProject, error)
	createPortfolioProjectFunc  func(ctx context.Context, project *models.PortfolioProject) error
	updatePortfolioProjectFunc  func(ctx context.Context, project *models.PortfolioProject) error
	deletePortfolioProjectFunc  func(ctx context.Context, id int64) error

	// Miniature Themes
	getAllMiniatureThemesFunc func(ctx context.Context) ([]models.MiniatureTheme, error)
	getMiniatureThemeByIDFunc func(ctx context.Context, id int64) (*models.MiniatureTheme, error)
	createMiniatureThemeFunc  func(ctx context.Context, theme *models.MiniatureTheme) error
	updateMiniatureThemeFunc  func(ctx context.Context, theme *models.MiniatureTheme) error
	deleteMiniatureThemeFunc  func(ctx context.Context, id int64) error

	// Miniature Projects
	getAllMiniatureProjectsFunc func(ctx context.Context) ([]models.MiniatureProject, error)
	getMiniatureProjectByIDFunc func(ctx context.Context, id int64) (*models.MiniatureProject, error)
	createMiniatureProjectFunc  func(ctx context.Context, project *models.MiniatureProject) error
	updateMiniatureProjectFunc  func(ctx context.Context, project *models.MiniatureProject) error
	deleteMiniatureProjectFunc  func(ctx context.Context, id int64) error
	addImageToProjectFunc       func(ctx context.Context, miniatureFile *models.MiniatureFile) error
	setProjectTechniquesFunc    func(ctx context.Context, projectID int64, techniqueIDs []int64) error
	setProjectPaintsFunc        func(ctx context.Context, projectID int64, paintIDs []int64) error

	// Miniature Techniques
	getAllTechniquesFunc func(ctx context.Context) ([]models.MiniatureTechnique, error)

	// Miniature Paints
	getAllMiniaturePaintsFunc func(ctx context.Context) ([]models.MiniaturePaint, error)
	getMiniaturePaintByIDFunc func(ctx context.Context, id int64) (*models.MiniaturePaint, error)
	createMiniaturePaintFunc  func(ctx context.Context, paint *models.MiniaturePaint) error
	updateMiniaturePaintFunc  func(ctx context.Context, paint *models.MiniaturePaint) error
	deleteMiniaturePaintFunc  func(ctx context.Context, id int64) error

	// Images/Files
	deleteImageFunc func(ctx context.Context, id int64) error
}

// Profile
func (m *mockRepository) GetProfile(ctx context.Context) (*models.Profile, error) {
	if m.getProfileFunc != nil {
		return m.getProfileFunc(ctx)
	}
	return &models.Profile{ID: 1}, nil
}

func (m *mockRepository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	if m.updateProfileFunc != nil {
		return m.updateProfileFunc(ctx, profile)
	}
	return nil
}

func (m *mockRepository) UpdateProfileAvatar(ctx context.Context, fileID int64) error {
	if m.updateProfileAvatarFunc != nil {
		return m.updateProfileAvatarFunc(ctx, fileID)
	}
	return nil
}

func (m *mockRepository) DeleteProfileAvatar(ctx context.Context) error {
	if m.deleteProfileAvatarFunc != nil {
		return m.deleteProfileAvatarFunc(ctx)
	}
	return nil
}

func (m *mockRepository) UpdateProfileResume(ctx context.Context, fileID int64) error {
	if m.updateProfileResumeFunc != nil {
		return m.updateProfileResumeFunc(ctx, fileID)
	}
	return nil
}

func (m *mockRepository) DeleteProfileResume(ctx context.Context) error {
	if m.deleteProfileResumeFunc != nil {
		return m.deleteProfileResumeFunc(ctx)
	}
	return nil
}

// Work Experience
func (m *mockRepository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	if m.getAllWorkExperienceFunc != nil {
		return m.getAllWorkExperienceFunc(ctx)
	}
	return []models.WorkExperience{}, nil
}

func (m *mockRepository) GetWorkExperienceByID(ctx context.Context, id int64) (*models.WorkExperience, error) {
	if m.getWorkExperienceByIDFunc != nil {
		return m.getWorkExperienceByIDFunc(ctx, id)
	}
	return &models.WorkExperience{ID: id}, nil
}

func (m *mockRepository) CreateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	if m.createWorkExperienceFunc != nil {
		return m.createWorkExperienceFunc(ctx, exp)
	}
	return nil
}

func (m *mockRepository) UpdateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	if m.updateWorkExperienceFunc != nil {
		return m.updateWorkExperienceFunc(ctx, exp)
	}
	return nil
}

func (m *mockRepository) DeleteWorkExperience(ctx context.Context, id int64) error {
	if m.deleteWorkExperienceFunc != nil {
		return m.deleteWorkExperienceFunc(ctx, id)
	}
	return nil
}

// Certifications
func (m *mockRepository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	if m.getAllCertificationsFunc != nil {
		return m.getAllCertificationsFunc(ctx)
	}
	return []models.Certification{}, nil
}

func (m *mockRepository) GetCertificationByID(ctx context.Context, id int64) (*models.Certification, error) {
	if m.getCertificationByIDFunc != nil {
		return m.getCertificationByIDFunc(ctx, id)
	}
	return &models.Certification{ID: id}, nil
}

func (m *mockRepository) CreateCertification(ctx context.Context, cert *models.Certification) error {
	if m.createCertificationFunc != nil {
		return m.createCertificationFunc(ctx, cert)
	}
	return nil
}

func (m *mockRepository) UpdateCertification(ctx context.Context, cert *models.Certification) error {
	if m.updateCertificationFunc != nil {
		return m.updateCertificationFunc(ctx, cert)
	}
	return nil
}

func (m *mockRepository) DeleteCertification(ctx context.Context, id int64) error {
	if m.deleteCertificationFunc != nil {
		return m.deleteCertificationFunc(ctx, id)
	}
	return nil
}

// Skills
func (m *mockRepository) GetAllSkills(ctx context.Context) ([]models.Skill, error) {
	if m.getAllSkillsFunc != nil {
		return m.getAllSkillsFunc(ctx)
	}
	return []models.Skill{}, nil
}

func (m *mockRepository) GetSkillByID(ctx context.Context, id int64) (*models.Skill, error) {
	if m.getSkillByIDFunc != nil {
		return m.getSkillByIDFunc(ctx, id)
	}
	return &models.Skill{ID: id}, nil
}

func (m *mockRepository) CreateSkill(ctx context.Context, skill *models.Skill) error {
	if m.createSkillFunc != nil {
		return m.createSkillFunc(ctx, skill)
	}
	return nil
}

func (m *mockRepository) UpdateSkill(ctx context.Context, skill *models.Skill) error {
	if m.updateSkillFunc != nil {
		return m.updateSkillFunc(ctx, skill)
	}
	return nil
}

func (m *mockRepository) DeleteSkill(ctx context.Context, id int64) error {
	if m.deleteSkillFunc != nil {
		return m.deleteSkillFunc(ctx, id)
	}
	return nil
}

// Skill Types
func (m *mockRepository) GetAllSkillTypes(ctx context.Context) ([]models.SkillType, error) {
	if m.getAllSkillTypesFunc != nil {
		return m.getAllSkillTypesFunc(ctx)
	}
	return []models.SkillType{}, nil
}

func (m *mockRepository) GetSkillTypeByID(ctx context.Context, id int64) (*models.SkillType, error) {
	if m.getSkillTypeByIDFunc != nil {
		return m.getSkillTypeByIDFunc(ctx, id)
	}
	return &models.SkillType{ID: id}, nil
}

func (m *mockRepository) CreateSkillType(ctx context.Context, skillType *models.SkillType) error {
	if m.createSkillTypeFunc != nil {
		return m.createSkillTypeFunc(ctx, skillType)
	}
	return nil
}

func (m *mockRepository) UpdateSkillType(ctx context.Context, skillType *models.SkillType) error {
	if m.updateSkillTypeFunc != nil {
		return m.updateSkillTypeFunc(ctx, skillType)
	}
	return nil
}

func (m *mockRepository) DeleteSkillType(ctx context.Context, id int64) error {
	if m.deleteSkillTypeFunc != nil {
		return m.deleteSkillTypeFunc(ctx, id)
	}
	return nil
}

// Portfolio Projects
func (m *mockRepository) GetAllPortfolioProjects(ctx context.Context) ([]models.PortfolioProject, error) {
	if m.getAllPortfolioProjectsFunc != nil {
		return m.getAllPortfolioProjectsFunc(ctx)
	}
	return []models.PortfolioProject{}, nil
}

func (m *mockRepository) GetPortfolioProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error) {
	if m.getPortfolioProjectByIDFunc != nil {
		return m.getPortfolioProjectByIDFunc(ctx, id)
	}
	return &models.PortfolioProject{ID: id}, nil
}

func (m *mockRepository) CreatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	if m.createPortfolioProjectFunc != nil {
		return m.createPortfolioProjectFunc(ctx, project)
	}
	return nil
}

func (m *mockRepository) UpdatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	if m.updatePortfolioProjectFunc != nil {
		return m.updatePortfolioProjectFunc(ctx, project)
	}
	return nil
}

func (m *mockRepository) DeletePortfolioProject(ctx context.Context, id int64) error {
	if m.deletePortfolioProjectFunc != nil {
		return m.deletePortfolioProjectFunc(ctx, id)
	}
	return nil
}

// Miniature Themes
func (m *mockRepository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	if m.getAllMiniatureThemesFunc != nil {
		return m.getAllMiniatureThemesFunc(ctx)
	}
	return []models.MiniatureTheme{}, nil
}

func (m *mockRepository) GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error) {
	if m.getMiniatureThemeByIDFunc != nil {
		return m.getMiniatureThemeByIDFunc(ctx, id)
	}
	return &models.MiniatureTheme{ID: id}, nil
}

func (m *mockRepository) CreateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	if m.createMiniatureThemeFunc != nil {
		return m.createMiniatureThemeFunc(ctx, theme)
	}
	return nil
}

func (m *mockRepository) UpdateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	if m.updateMiniatureThemeFunc != nil {
		return m.updateMiniatureThemeFunc(ctx, theme)
	}
	return nil
}

func (m *mockRepository) DeleteMiniatureTheme(ctx context.Context, id int64) error {
	if m.deleteMiniatureThemeFunc != nil {
		return m.deleteMiniatureThemeFunc(ctx, id)
	}
	return nil
}

// Miniature Projects
func (m *mockRepository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	if m.getAllMiniatureProjectsFunc != nil {
		return m.getAllMiniatureProjectsFunc(ctx)
	}
	return []models.MiniatureProject{}, nil
}

func (m *mockRepository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	if m.getMiniatureProjectByIDFunc != nil {
		return m.getMiniatureProjectByIDFunc(ctx, id)
	}
	return &models.MiniatureProject{ID: id}, nil
}

func (m *mockRepository) CreateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	if m.createMiniatureProjectFunc != nil {
		return m.createMiniatureProjectFunc(ctx, project)
	}
	return nil
}

func (m *mockRepository) UpdateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	if m.updateMiniatureProjectFunc != nil {
		return m.updateMiniatureProjectFunc(ctx, project)
	}
	return nil
}

func (m *mockRepository) DeleteMiniatureProject(ctx context.Context, id int64) error {
	if m.deleteMiniatureProjectFunc != nil {
		return m.deleteMiniatureProjectFunc(ctx, id)
	}
	return nil
}

func (m *mockRepository) AddImageToProject(ctx context.Context, miniatureFile *models.MiniatureFile) error {
	if m.addImageToProjectFunc != nil {
		return m.addImageToProjectFunc(ctx, miniatureFile)
	}
	return nil
}

func (m *mockRepository) SetProjectTechniques(ctx context.Context, projectID int64, techniqueIDs []int64) error {
	if m.setProjectTechniquesFunc != nil {
		return m.setProjectTechniquesFunc(ctx, projectID, techniqueIDs)
	}
	return nil
}

func (m *mockRepository) SetProjectPaints(ctx context.Context, projectID int64, paintIDs []int64) error {
	if m.setProjectPaintsFunc != nil {
		return m.setProjectPaintsFunc(ctx, projectID, paintIDs)
	}
	return nil
}

// Miniature Techniques
func (m *mockRepository) GetAllTechniques(ctx context.Context) ([]models.MiniatureTechnique, error) {
	if m.getAllTechniquesFunc != nil {
		return m.getAllTechniquesFunc(ctx)
	}
	return []models.MiniatureTechnique{}, nil
}

// Miniature Paints
func (m *mockRepository) GetAllMiniaturePaints(ctx context.Context) ([]models.MiniaturePaint, error) {
	if m.getAllMiniaturePaintsFunc != nil {
		return m.getAllMiniaturePaintsFunc(ctx)
	}
	return []models.MiniaturePaint{}, nil
}

func (m *mockRepository) GetMiniaturePaintByID(ctx context.Context, id int64) (*models.MiniaturePaint, error) {
	if m.getMiniaturePaintByIDFunc != nil {
		return m.getMiniaturePaintByIDFunc(ctx, id)
	}
	return &models.MiniaturePaint{ID: id}, nil
}

func (m *mockRepository) CreateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	if m.createMiniaturePaintFunc != nil {
		return m.createMiniaturePaintFunc(ctx, paint)
	}
	return nil
}

func (m *mockRepository) UpdateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	if m.updateMiniaturePaintFunc != nil {
		return m.updateMiniaturePaintFunc(ctx, paint)
	}
	return nil
}

func (m *mockRepository) DeleteMiniaturePaint(ctx context.Context, id int64) error {
	if m.deleteMiniaturePaintFunc != nil {
		return m.deleteMiniaturePaintFunc(ctx, id)
	}
	return nil
}

// Images/Files
func (m *mockRepository) DeleteImage(ctx context.Context, id int64) error {
	if m.deleteImageFunc != nil {
		return m.deleteImageFunc(ctx, id)
	}
	return nil
}

// =============================================================================
// Test Helpers
// =============================================================================

// injectScopes creates middleware that injects scopes into the Gin context,
// simulating what ValidateToken middleware does with JWT claims.
func injectScopes(scopes map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("scopes", scopes)
		c.Next()
	}
}

// setupRouterWithScopes creates a test router with custom scope injection
// for testing permission scenarios without real JWT validation.
func setupRouterWithScopes(t *testing.T, scopes map[string]string) *gin.Engine {
	t.Helper()

	router := gin.New()
	handler := handlers.New(&mockRepository{})

	v1 := router.Group("/api/v1")
	v1.Use(injectScopes(scopes))
	{
		// Portfolio domain
		portfolio := v1.Group("/portfolio")
		{
			portfolio.GET("/profile", common.RequirePermission(common.ResourceProfile, common.LevelRead), handler.GetProfile)
			portfolio.PUT("/profile", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfile)
			portfolio.PUT("/profile/avatar", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfileAvatar)
			portfolio.DELETE("/profile/avatar", common.RequirePermission(common.ResourceProfile, common.LevelDelete), handler.DeleteProfileAvatar)
			portfolio.PUT("/profile/resume", common.RequirePermission(common.ResourceProfile, common.LevelEdit), handler.UpdateProfileResume)
			portfolio.DELETE("/profile/resume", common.RequirePermission(common.ResourceProfile, common.LevelDelete), handler.DeleteProfileResume)

			portfolio.GET("/experience", common.RequirePermission(common.ResourceExperience, common.LevelRead), handler.GetAllWorkExperience)
			portfolio.POST("/experience", common.RequirePermission(common.ResourceExperience, common.LevelEdit), handler.CreateWorkExperience)
			portfolio.GET("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelRead), handler.GetWorkExperienceByID)
			portfolio.PUT("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelEdit), handler.UpdateWorkExperience)
			portfolio.DELETE("/experience/:id", common.RequirePermission(common.ResourceExperience, common.LevelDelete), handler.DeleteWorkExperience)

			portfolio.GET("/certifications", common.RequirePermission(common.ResourceCertifications, common.LevelRead), handler.GetAllCertifications)
			portfolio.POST("/certifications", common.RequirePermission(common.ResourceCertifications, common.LevelEdit), handler.CreateCertification)
			portfolio.GET("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelRead), handler.GetCertificationByID)
			portfolio.PUT("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelEdit), handler.UpdateCertification)
			portfolio.DELETE("/certifications/:id", common.RequirePermission(common.ResourceCertifications, common.LevelDelete), handler.DeleteCertification)

			portfolio.GET("/skills", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetAllSkills)
			portfolio.POST("/skills", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.CreateSkill)
			portfolio.GET("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetSkillByID)
			portfolio.PUT("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.UpdateSkill)
			portfolio.DELETE("/skills/:id", common.RequirePermission(common.ResourceSkills, common.LevelDelete), handler.DeleteSkill)

			portfolio.GET("/skill-types", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetAllSkillTypes)
			portfolio.POST("/skill-types", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.CreateSkillType)
			portfolio.GET("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelRead), handler.GetSkillTypeByID)
			portfolio.PUT("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelEdit), handler.UpdateSkillType)
			portfolio.DELETE("/skill-types/:id", common.RequirePermission(common.ResourceSkills, common.LevelDelete), handler.DeleteSkillType)

			portfolio.GET("/projects", common.RequirePermission(common.ResourceProjects, common.LevelRead), handler.GetAllPortfolioProjects)
			portfolio.POST("/projects", common.RequirePermission(common.ResourceProjects, common.LevelEdit), handler.CreatePortfolioProject)
			portfolio.GET("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelRead), handler.GetPortfolioProjectByID)
			portfolio.PUT("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelEdit), handler.UpdatePortfolioProject)
			portfolio.DELETE("/projects/:id", common.RequirePermission(common.ResourceProjects, common.LevelDelete), handler.DeletePortfolioProject)
		}

		// Miniatures domain
		miniatures := v1.Group("/miniatures")
		{
			miniatures.GET("/themes", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniatureThemes)
			miniatures.POST("/themes", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniatureTheme)
			miniatures.GET("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniatureThemeByID)
			miniatures.PUT("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniatureTheme)
			miniatures.DELETE("/themes/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniatureTheme)

			miniatures.GET("/projects", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniatureProjects)
			miniatures.POST("/projects", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniatureProject)
			miniatures.GET("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniatureProjectByID)
			miniatures.PUT("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniatureProject)
			miniatures.DELETE("/projects/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniatureProject)
			miniatures.POST("/projects/:id/images", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.AddImageToProject)
			miniatures.PUT("/projects/:id/techniques", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.SetProjectTechniques)
			miniatures.PUT("/projects/:id/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.SetProjectPaints)

			miniatures.GET("/techniques", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllTechniques)

			miniatures.GET("/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetAllMiniaturePaints)
			miniatures.POST("/paints", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.CreateMiniaturePaint)
			miniatures.GET("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelRead), handler.GetMiniaturePaintByID)
			miniatures.PUT("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelEdit), handler.UpdateMiniaturePaint)
			miniatures.DELETE("/paints/:id", common.RequirePermission(common.ResourceMiniatures, common.LevelDelete), handler.DeleteMiniaturePaint)
		}

		// Files
		v1.DELETE("/files/:id", common.RequirePermission(common.ResourceFiles, common.LevelDelete), handler.DeleteImage)
	}

	return router
}

func performRequest(t *testing.T, router *gin.Engine, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// =============================================================================
// Route Permission Definitions
// =============================================================================

type routePermission struct {
	method   string
	path     string
	resource string
	level    string
}

var portfolioRoutes = []routePermission{
	// Profile
	{"GET", "/api/v1/portfolio/profile", common.ResourceProfile, common.LevelRead},
	{"PUT", "/api/v1/portfolio/profile", common.ResourceProfile, common.LevelEdit},
	{"PUT", "/api/v1/portfolio/profile/avatar", common.ResourceProfile, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/profile/avatar", common.ResourceProfile, common.LevelDelete},
	{"PUT", "/api/v1/portfolio/profile/resume", common.ResourceProfile, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/profile/resume", common.ResourceProfile, common.LevelDelete},
	// Work Experience
	{"GET", "/api/v1/portfolio/experience", common.ResourceExperience, common.LevelRead},
	{"POST", "/api/v1/portfolio/experience", common.ResourceExperience, common.LevelEdit},
	{"GET", "/api/v1/portfolio/experience/1", common.ResourceExperience, common.LevelRead},
	{"PUT", "/api/v1/portfolio/experience/1", common.ResourceExperience, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/experience/1", common.ResourceExperience, common.LevelDelete},
	// Certifications
	{"GET", "/api/v1/portfolio/certifications", common.ResourceCertifications, common.LevelRead},
	{"POST", "/api/v1/portfolio/certifications", common.ResourceCertifications, common.LevelEdit},
	{"GET", "/api/v1/portfolio/certifications/1", common.ResourceCertifications, common.LevelRead},
	{"PUT", "/api/v1/portfolio/certifications/1", common.ResourceCertifications, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/certifications/1", common.ResourceCertifications, common.LevelDelete},
	// Skills
	{"GET", "/api/v1/portfolio/skills", common.ResourceSkills, common.LevelRead},
	{"POST", "/api/v1/portfolio/skills", common.ResourceSkills, common.LevelEdit},
	{"GET", "/api/v1/portfolio/skills/1", common.ResourceSkills, common.LevelRead},
	{"PUT", "/api/v1/portfolio/skills/1", common.ResourceSkills, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/skills/1", common.ResourceSkills, common.LevelDelete},
	// Skill Types
	{"GET", "/api/v1/portfolio/skill-types", common.ResourceSkills, common.LevelRead},
	{"POST", "/api/v1/portfolio/skill-types", common.ResourceSkills, common.LevelEdit},
	{"GET", "/api/v1/portfolio/skill-types/1", common.ResourceSkills, common.LevelRead},
	{"PUT", "/api/v1/portfolio/skill-types/1", common.ResourceSkills, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/skill-types/1", common.ResourceSkills, common.LevelDelete},
	// Portfolio Projects
	{"GET", "/api/v1/portfolio/projects", common.ResourceProjects, common.LevelRead},
	{"POST", "/api/v1/portfolio/projects", common.ResourceProjects, common.LevelEdit},
	{"GET", "/api/v1/portfolio/projects/1", common.ResourceProjects, common.LevelRead},
	{"PUT", "/api/v1/portfolio/projects/1", common.ResourceProjects, common.LevelEdit},
	{"DELETE", "/api/v1/portfolio/projects/1", common.ResourceProjects, common.LevelDelete},
}

var miniaturesRoutes = []routePermission{
	// Themes
	{"GET", "/api/v1/miniatures/themes", common.ResourceMiniatures, common.LevelRead},
	{"POST", "/api/v1/miniatures/themes", common.ResourceMiniatures, common.LevelEdit},
	{"GET", "/api/v1/miniatures/themes/1", common.ResourceMiniatures, common.LevelRead},
	{"PUT", "/api/v1/miniatures/themes/1", common.ResourceMiniatures, common.LevelEdit},
	{"DELETE", "/api/v1/miniatures/themes/1", common.ResourceMiniatures, common.LevelDelete},
	// Projects
	{"GET", "/api/v1/miniatures/projects", common.ResourceMiniatures, common.LevelRead},
	{"POST", "/api/v1/miniatures/projects", common.ResourceMiniatures, common.LevelEdit},
	{"GET", "/api/v1/miniatures/projects/1", common.ResourceMiniatures, common.LevelRead},
	{"PUT", "/api/v1/miniatures/projects/1", common.ResourceMiniatures, common.LevelEdit},
	{"DELETE", "/api/v1/miniatures/projects/1", common.ResourceMiniatures, common.LevelDelete},
	{"POST", "/api/v1/miniatures/projects/1/images", common.ResourceMiniatures, common.LevelEdit},
	{"PUT", "/api/v1/miniatures/projects/1/techniques", common.ResourceMiniatures, common.LevelEdit},
	{"PUT", "/api/v1/miniatures/projects/1/paints", common.ResourceMiniatures, common.LevelEdit},
	// Techniques
	{"GET", "/api/v1/miniatures/techniques", common.ResourceMiniatures, common.LevelRead},
	// Paints
	{"GET", "/api/v1/miniatures/paints", common.ResourceMiniatures, common.LevelRead},
	{"POST", "/api/v1/miniatures/paints", common.ResourceMiniatures, common.LevelEdit},
	{"GET", "/api/v1/miniatures/paints/1", common.ResourceMiniatures, common.LevelRead},
	{"PUT", "/api/v1/miniatures/paints/1", common.ResourceMiniatures, common.LevelEdit},
	{"DELETE", "/api/v1/miniatures/paints/1", common.ResourceMiniatures, common.LevelDelete},
}

var filesRoutes = []routePermission{
	{"DELETE", "/api/v1/files/1", common.ResourceFiles, common.LevelDelete},
}

// =============================================================================
// Portfolio Route Permission Tests
// =============================================================================

func TestPortfolioRoutes_Forbidden_WithoutPermission(t *testing.T) {
	for _, route := range portfolioRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			router := setupRouterWithScopes(t, map[string]string{})
			w := performRequest(t, router, route.method, route.path)

			if w.Code != http.StatusForbidden {
				t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if response["error"] != "insufficient permissions" {
				t.Errorf("error = %v, want 'insufficient permissions'", response["error"])
			}
			if response["resource"] != route.resource {
				t.Errorf("resource = %v, want %q", response["resource"], route.resource)
			}
			if response["required"] != route.level {
				t.Errorf("required = %v, want %q", response["required"], route.level)
			}
		})
	}
}

func TestPortfolioRoutes_Allowed_WithPermission(t *testing.T) {
	for _, route := range portfolioRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			scopes := map[string]string{route.resource: route.level}
			router := setupRouterWithScopes(t, scopes)
			w := performRequest(t, router, route.method, route.path)

			if w.Code == http.StatusForbidden {
				t.Errorf("got 403 Forbidden with permission %s:%s", route.resource, route.level)
			}
			if w.Code == http.StatusUnauthorized {
				t.Errorf("got 401 Unauthorized - scopes not injected")
			}
		})
	}
}

// =============================================================================
// Miniatures Route Permission Tests
// =============================================================================

func TestMiniaturesRoutes_Forbidden_WithoutPermission(t *testing.T) {
	for _, route := range miniaturesRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			router := setupRouterWithScopes(t, map[string]string{})
			w := performRequest(t, router, route.method, route.path)

			if w.Code != http.StatusForbidden {
				t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
			}
		})
	}
}

func TestMiniaturesRoutes_Allowed_WithPermission(t *testing.T) {
	for _, route := range miniaturesRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			scopes := map[string]string{route.resource: route.level}
			router := setupRouterWithScopes(t, scopes)
			w := performRequest(t, router, route.method, route.path)

			if w.Code == http.StatusForbidden {
				t.Errorf("got 403 Forbidden with permission %s:%s", route.resource, route.level)
			}
		})
	}
}

// =============================================================================
// Files Route Permission Tests
// =============================================================================

func TestFilesRoutes_Forbidden_WithoutPermission(t *testing.T) {
	for _, route := range filesRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			router := setupRouterWithScopes(t, map[string]string{})
			w := performRequest(t, router, route.method, route.path)

			if w.Code != http.StatusForbidden {
				t.Errorf("status = %d, want %d", w.Code, http.StatusForbidden)
			}
		})
	}
}

func TestFilesRoutes_Allowed_WithPermission(t *testing.T) {
	for _, route := range filesRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			scopes := map[string]string{route.resource: route.level}
			router := setupRouterWithScopes(t, scopes)
			w := performRequest(t, router, route.method, route.path)

			if w.Code == http.StatusForbidden {
				t.Errorf("got 403 Forbidden with permission %s:%s", route.resource, route.level)
			}
		})
	}
}

// =============================================================================
// Permission Hierarchy Tests
// =============================================================================

func TestPermissionHierarchy(t *testing.T) {
	tests := []struct {
		name       string
		granted    string
		required   string
		wantAccess bool
	}{
		{"delete grants delete", common.LevelDelete, common.LevelDelete, true},
		{"delete grants edit", common.LevelDelete, common.LevelEdit, true},
		{"delete grants read", common.LevelDelete, common.LevelRead, true},
		{"edit grants edit", common.LevelEdit, common.LevelEdit, true},
		{"edit grants read", common.LevelEdit, common.LevelRead, true},
		{"edit denies delete", common.LevelEdit, common.LevelDelete, false},
		{"read grants read", common.LevelRead, common.LevelRead, true},
		{"read denies edit", common.LevelRead, common.LevelEdit, false},
		{"read denies delete", common.LevelRead, common.LevelDelete, false},
		{"none denies read", common.LevelNone, common.LevelRead, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var method, path string
			switch tt.required {
			case common.LevelRead:
				method, path = "GET", "/api/v1/portfolio/profile"
			case common.LevelEdit:
				method, path = "PUT", "/api/v1/portfolio/profile"
			case common.LevelDelete:
				method, path = "DELETE", "/api/v1/portfolio/profile/avatar"
			}

			scopes := map[string]string{common.ResourceProfile: tt.granted}
			router := setupRouterWithScopes(t, scopes)
			w := performRequest(t, router, method, path)

			gotAccess := w.Code != http.StatusForbidden
			if gotAccess != tt.wantAccess {
				t.Errorf("granted=%s required=%s: gotAccess=%v wantAccess=%v (status=%d)",
					tt.granted, tt.required, gotAccess, tt.wantAccess, w.Code)
			}
		})
	}
}

// =============================================================================
// Cross-Resource Permission Tests
// =============================================================================

func TestCrossResourcePermissions_Denied(t *testing.T) {
	// Profile delete permission should NOT grant experience read
	scopes := map[string]string{common.ResourceProfile: common.LevelDelete}
	router := setupRouterWithScopes(t, scopes)

	w := performRequest(t, router, "GET", "/api/v1/portfolio/experience")

	if w.Code != http.StatusForbidden {
		t.Errorf("profile:delete should not grant experience:read, got status %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if response["resource"] != common.ResourceExperience {
		t.Errorf("error resource = %v, want %q", response["resource"], common.ResourceExperience)
	}
}

func TestMultipleResourcePermissions(t *testing.T) {
	scopes := map[string]string{
		common.ResourceProfile:        common.LevelRead,
		common.ResourceExperience:     common.LevelEdit,
		common.ResourceCertifications: common.LevelDelete,
	}
	router := setupRouterWithScopes(t, scopes)

	tests := []struct {
		method     string
		path       string
		wantAccess bool
	}{
		// Profile: read only
		{"GET", "/api/v1/portfolio/profile", true},
		{"PUT", "/api/v1/portfolio/profile", false},
		// Experience: edit (includes read)
		{"GET", "/api/v1/portfolio/experience", true},
		{"POST", "/api/v1/portfolio/experience", true},
		{"DELETE", "/api/v1/portfolio/experience/1", false},
		// Certifications: delete (includes all)
		{"GET", "/api/v1/portfolio/certifications", true},
		{"POST", "/api/v1/portfolio/certifications", true},
		{"DELETE", "/api/v1/portfolio/certifications/1", true},
		// Skills: no permission
		{"GET", "/api/v1/portfolio/skills", false},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			w := performRequest(t, router, tt.method, tt.path)
			gotAccess := w.Code != http.StatusForbidden

			if gotAccess != tt.wantAccess {
				t.Errorf("gotAccess=%v wantAccess=%v (status=%d)", gotAccess, tt.wantAccess, w.Code)
			}
		})
	}
}

// =============================================================================
// Middleware Error Handling Tests
// =============================================================================

func TestRoutes_NoScopes_Unauthorized(t *testing.T) {
	router := gin.New()
	handler := handlers.New(&mockRepository{})

	// Route without scope injection middleware
	router.GET("/api/v1/portfolio/profile",
		common.RequirePermission(common.ResourceProfile, common.LevelRead),
		handler.GetProfile,
	)

	req, _ := http.NewRequest("GET", "/api/v1/portfolio/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d (no scopes = unauthorized)", w.Code, http.StatusUnauthorized)
	}
}

func TestRoutes_InvalidScopesFormat_InternalError(t *testing.T) {
	router := gin.New()
	handler := handlers.New(&mockRepository{})

	// Inject invalid scopes format (string instead of map)
	router.Use(func(c *gin.Context) {
		c.Set("scopes", "invalid-format")
		c.Next()
	})

	router.GET("/api/v1/portfolio/profile",
		common.RequirePermission(common.ResourceProfile, common.LevelRead),
		handler.GetProfile,
	)

	req, _ := http.NewRequest("GET", "/api/v1/portfolio/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d (invalid scopes = internal error)", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Repository Error Propagation Test
// =============================================================================

func TestRoutes_RepositoryError_Propagated(t *testing.T) {
	router := gin.New()
	mockRepo := &mockRepository{
		getProfileFunc: func(ctx context.Context) (*models.Profile, error) {
			return nil, errors.New("database connection failed")
		},
	}
	handler := handlers.New(mockRepo)

	router.Use(injectScopes(map[string]string{common.ResourceProfile: common.LevelRead}))
	router.GET("/api/v1/portfolio/profile", handler.GetProfile)

	req, _ := http.NewRequest("GET", "/api/v1/portfolio/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d (repository error)", w.Code, http.StatusInternalServerError)
	}
}
