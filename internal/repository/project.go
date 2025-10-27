package repository

import (
	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"gorm.io/gorm"
)

func (r *repository) GetAllPortfolioProjects() ([]models.PortfolioProject, error) {
	var projects []models.PortfolioProject
	err := r.db.Preload("Technologies").Preload("ImageFile").Order("display_order ASC, created_at DESC").Find(&projects).Error
	if err != nil {
		return nil, err
	}

	// Populate image URLs for all projects
	for i := range projects {
		utils.PopulateFileURL(projects[i].ImageFile, r.filesAPIURL)
	}

	return projects, nil
}

func (r *repository) GetPortfolioProjectByID(id int64) (*models.PortfolioProject, error) {
	var project models.PortfolioProject
	err := r.db.Preload("Technologies").Preload("ImageFile").First(&project, id).Error
	if err != nil {
		return nil, err
	}

	// Populate image URL
	utils.PopulateFileURL(project.ImageFile, r.filesAPIURL)

	return &project, err
}

func (r *repository) CreatePortfolioProject(project *models.PortfolioProject) error {
	return r.db.Create(project).Error
}

func (r *repository) UpdatePortfolioProject(project *models.PortfolioProject) error {
	// Save will update all fields including associations
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(project).Error
}

// DeletePortfolioProject deletes a portfolio project and automatically cascades to:
// - portfolio.project_technologies (links to skills/technologies)
// Note: Image file in storage.files is NOT deleted (cleanup job handles orphaned files)
func (r *repository) DeletePortfolioProject(id int64) error {
	return r.db.Delete(&models.PortfolioProject{}, id).Error
}
