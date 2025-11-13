package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.WithContext(ctx).Order("issue_date DESC").Find(&certifications).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all certifications: %w", err)
	}
	return certifications, nil
}

func (r *repository) GetCertificationByID(ctx context.Context, id int64) (*models.Certification, error) {
	var cert models.Certification
	err := r.db.WithContext(ctx).First(&cert, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get certification by id %d: %w", id, err)
	}
	return &cert, nil
}

func (r *repository) CreateCertification(ctx context.Context, cert *models.Certification) error {
	if err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(cert).Error; err != nil {
		return fmt.Errorf("failed to create certification: %w", err)
	}
	return nil
}

func (r *repository) UpdateCertification(ctx context.Context, cert *models.Certification) error {
	return r.safeUpdate(ctx, cert, cert.ID)
}

func (r *repository) DeleteCertification(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.Certification{}, id))
}
