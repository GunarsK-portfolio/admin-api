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

// Profile methods
func (r *repository) GetProfile() (*models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile).Error
	return &profile, err
}

func (r *repository) UpdateProfile(profile *models.Profile) error {
	return r.db.Save(profile).Error
}

// Work Experience methods
func (r *repository) GetAllWorkExperience() ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.Order("display_order ASC").Find(&experiences).Error
	return experiences, err
}

func (r *repository) GetWorkExperienceByID(id int64) (*models.WorkExperience, error) {
	var exp models.WorkExperience
	err := r.db.First(&exp, id).Error
	return &exp, err
}

func (r *repository) CreateWorkExperience(exp *models.WorkExperience) error {
	return r.db.Create(exp).Error
}

func (r *repository) UpdateWorkExperience(exp *models.WorkExperience) error {
	return r.db.Save(exp).Error
}

func (r *repository) DeleteWorkExperience(id int64) error {
	return r.db.Delete(&models.WorkExperience{}, id).Error
}

// Certification methods
func (r *repository) GetAllCertifications() ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.Order("display_order ASC").Find(&certifications).Error
	return certifications, err
}

func (r *repository) GetCertificationByID(id int64) (*models.Certification, error) {
	var cert models.Certification
	err := r.db.First(&cert, id).Error
	return &cert, err
}

func (r *repository) CreateCertification(cert *models.Certification) error {
	return r.db.Create(cert).Error
}

func (r *repository) UpdateCertification(cert *models.Certification) error {
	return r.db.Save(cert).Error
}

func (r *repository) DeleteCertification(id int64) error {
	return r.db.Delete(&models.Certification{}, id).Error
}

// Miniature Project methods
func (r *repository) GetAllMiniatureProjects() ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	err := r.db.Preload("Images").Order("display_order ASC").Find(&projects).Error
	return projects, err
}

func (r *repository) GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	err := r.db.Preload("Images").First(&project, id).Error
	return &project, err
}

func (r *repository) CreateMiniatureProject(project *models.MiniatureProject) error {
	return r.db.Create(project).Error
}

func (r *repository) UpdateMiniatureProject(project *models.MiniatureProject) error {
	return r.db.Save(project).Error
}

func (r *repository) DeleteMiniatureProject(id int64) error {
	return r.db.Delete(&models.MiniatureProject{}, id).Error
}

// Image methods
func (r *repository) CreateImage(image *models.Image) error {
	return r.db.Create(image).Error
}

func (r *repository) DeleteImage(id int64) error {
	return r.db.Delete(&models.Image{}, id).Error
}
