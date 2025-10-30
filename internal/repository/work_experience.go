package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.WithContext(ctx).Order("start_date DESC").Find(&experiences).Error
	return experiences, err
}

func (r *repository) GetWorkExperienceByID(ctx context.Context, id int64) (*models.WorkExperience, error) {
	var exp models.WorkExperience
	err := r.db.WithContext(ctx).First(&exp, id).Error
	return &exp, err
}

func (r *repository) CreateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	return r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(exp).Error
}

func (r *repository) UpdateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	return r.safeUpdate(ctx, exp, exp.ID)
}

func (r *repository) DeleteWorkExperience(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.WorkExperience{}, id))
}
