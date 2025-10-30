package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	// Profile
	GetProfile(ctx context.Context) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) error
	UpdateProfileAvatar(ctx context.Context, fileID int64) error
	DeleteProfileAvatar(ctx context.Context) error
	UpdateProfileResume(ctx context.Context, fileID int64) error
	DeleteProfileResume(ctx context.Context) error

	// Work Experience
	GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error)
	GetWorkExperienceByID(ctx context.Context, id int64) (*models.WorkExperience, error)
	CreateWorkExperience(ctx context.Context, exp *models.WorkExperience) error
	UpdateWorkExperience(ctx context.Context, exp *models.WorkExperience) error
	DeleteWorkExperience(ctx context.Context, id int64) error

	// Certifications
	GetAllCertifications(ctx context.Context) ([]models.Certification, error)
	GetCertificationByID(ctx context.Context, id int64) (*models.Certification, error)
	CreateCertification(ctx context.Context, cert *models.Certification) error
	UpdateCertification(ctx context.Context, cert *models.Certification) error
	DeleteCertification(ctx context.Context, id int64) error

	// Miniature Themes
	GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error)
	GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error)
	CreateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error
	UpdateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error
	DeleteMiniatureTheme(ctx context.Context, id int64) error

	// Miniature Projects
	GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error)
	CreateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error
	UpdateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error
	DeleteMiniatureProject(ctx context.Context, id int64) error

	// Miniature Paints
	GetAllMiniaturePaints(ctx context.Context) ([]models.MiniaturePaint, error)
	GetMiniaturePaintByID(ctx context.Context, id int64) (*models.MiniaturePaint, error)
	CreateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error
	UpdateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error
	DeleteMiniaturePaint(ctx context.Context, id int64) error

	// Skills
	GetAllSkills(ctx context.Context) ([]models.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*models.Skill, error)
	CreateSkill(ctx context.Context, skill *models.Skill) error
	UpdateSkill(ctx context.Context, skill *models.Skill) error
	DeleteSkill(ctx context.Context, id int64) error

	// Skill Types
	GetAllSkillTypes(ctx context.Context) ([]models.SkillType, error)
	GetSkillTypeByID(ctx context.Context, id int64) (*models.SkillType, error)
	CreateSkillType(ctx context.Context, skillType *models.SkillType) error
	UpdateSkillType(ctx context.Context, skillType *models.SkillType) error
	DeleteSkillType(ctx context.Context, id int64) error

	// Portfolio Projects
	GetAllPortfolioProjects(ctx context.Context) ([]models.PortfolioProject, error)
	GetPortfolioProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error)
	CreatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error
	UpdatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error
	DeletePortfolioProject(ctx context.Context, id int64) error

	// Images/Files (MinIO storage references)
	DeleteImage(ctx context.Context, id int64) error
}

type repository struct {
	db          *gorm.DB
	filesAPIURL string
}

func New(db *gorm.DB, filesAPIURL string) Repository {
	return &repository{
		db:          db,
		filesAPIURL: filesAPIURL,
	}
}

// checkRowsAffected returns gorm.ErrRecordNotFound if no rows were affected
func checkRowsAffected(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// safeUpdate performs an update excluding system fields (ID, CreatedAt, UpdatedAt)
// Uses Updates to avoid zero-value overwrites unlike Save
func (r *repository) safeUpdate(ctx context.Context, model interface{}, id int64) error {
	return checkRowsAffected(
		r.db.WithContext(ctx).Model(model).
			Where("id = ?", id).
			Omit("ID", "CreatedAt", "UpdatedAt").
			Updates(model),
	)
}
