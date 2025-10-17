package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

func (r *repository) CreateImage(image *models.Image) error {
	return r.db.Create(image).Error
}

func (r *repository) DeleteImage(id int64) error {
	return r.db.Delete(&models.Image{}, id).Error
}
