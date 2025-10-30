package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.WithContext(ctx).Order("issue_date DESC").Find(&certifications).Error
	return certifications, err
}

func (r *repository) GetCertificationByID(ctx context.Context, id int64) (*models.Certification, error) {
	var cert models.Certification
	err := r.db.WithContext(ctx).First(&cert, id).Error
	return &cert, err
}

func (r *repository) CreateCertification(ctx context.Context, cert *models.Certification) error {
	return r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(cert).Error
}

func (r *repository) UpdateCertification(ctx context.Context, cert *models.Certification) error {
	return r.safeUpdate(ctx, cert, cert.ID)
}

func (r *repository) DeleteCertification(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.Certification{}, id))
}
