package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

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
