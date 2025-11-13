package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.WithContext(ctx).Order("start_date DESC").Find(&experiences).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all work experience: %w", err)
	}
	return experiences, nil
}

func (r *repository) GetWorkExperienceByID(ctx context.Context, id int64) (*models.WorkExperience, error) {
	var exp models.WorkExperience
	err := r.db.WithContext(ctx).First(&exp, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get work experience with id %d: %w", id, err)
	}
	return &exp, nil
}

func (r *repository) CreateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(exp).Error
	if err != nil {
		return fmt.Errorf("failed to create work experience: %w", err)
	}
	return nil
}

func (r *repository) UpdateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	return r.safeUpdate(ctx, exp, exp.ID)
}

func (r *repository) DeleteWorkExperience(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.WorkExperience{}, id))
}
