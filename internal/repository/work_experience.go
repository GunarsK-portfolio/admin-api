package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

func (r *repository) GetAllWorkExperience() ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.Order("start_date DESC").Find(&experiences).Error
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
	return checkRowsAffected(r.db.Save(exp))
}

func (r *repository) DeleteWorkExperience(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.WorkExperience{}, id))
}
