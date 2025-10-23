package repository

import (
	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	// Profile
	GetProfile() (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error
	UpdateProfileAvatar(fileID int64) error
	DeleteProfileAvatar() error
	UpdateProfileResume(fileID int64) error
	DeleteProfileResume() error

	// Work Experience
	GetAllWorkExperience() ([]models.WorkExperience, error)
	GetWorkExperienceByID(id int64) (*models.WorkExperience, error)
	CreateWorkExperience(exp *models.WorkExperience) error
	UpdateWorkExperience(exp *models.WorkExperience) error
	DeleteWorkExperience(id int64) error

	// Certifications
	GetAllCertifications() ([]models.Certification, error)
	GetCertificationByID(id int64) (*models.Certification, error)
	CreateCertification(cert *models.Certification) error
	UpdateCertification(cert *models.Certification) error
	DeleteCertification(id int64) error

	// Miniature Themes
	GetAllMiniatureThemes() ([]models.MiniatureTheme, error)
	GetMiniatureThemeByID(id int64) (*models.MiniatureTheme, error)
	CreateMiniatureTheme(theme *models.MiniatureTheme) error
	UpdateMiniatureTheme(theme *models.MiniatureTheme) error
	DeleteMiniatureTheme(id int64) error

	// Miniature Projects
	GetAllMiniatureProjects() ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error)
	CreateMiniatureProject(project *models.MiniatureProject) error
	UpdateMiniatureProject(project *models.MiniatureProject) error
	DeleteMiniatureProject(id int64) error

	// Skills
	GetAllSkills() ([]models.Skill, error)
	GetSkillByID(id int64) (*models.Skill, error)
	CreateSkill(skill *models.Skill) error
	UpdateSkill(skill *models.Skill) error
	DeleteSkill(id int64) error

	// Skill Types
	GetAllSkillTypes() ([]models.SkillType, error)
	GetSkillTypeByID(id int64) (*models.SkillType, error)
	CreateSkillType(skillType *models.SkillType) error
	UpdateSkillType(skillType *models.SkillType) error
	DeleteSkillType(id int64) error

	// Portfolio Projects
	GetAllPortfolioProjects() ([]models.PortfolioProject, error)
	GetPortfolioProjectByID(id int64) (*models.PortfolioProject, error)
	CreatePortfolioProject(project *models.PortfolioProject) error
	UpdatePortfolioProject(project *models.PortfolioProject) error
	DeletePortfolioProject(id int64) error

	// Images/Files (MinIO storage references)
	DeleteImage(id int64) error
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
