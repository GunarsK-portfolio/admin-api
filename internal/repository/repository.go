package repository

import (
	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	// Profile
	GetProfile() (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error

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

	// Miniature Projects
	GetAllMiniatureProjects() ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error)
	CreateMiniatureProject(project *models.MiniatureProject) error
	UpdateMiniatureProject(project *models.MiniatureProject) error
	DeleteMiniatureProject(id int64) error

	// Images
	CreateImage(image *models.Image) error
	DeleteImage(id int64) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}
