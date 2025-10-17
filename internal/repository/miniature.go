package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

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
