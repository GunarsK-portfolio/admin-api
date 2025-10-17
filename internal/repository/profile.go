package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

func (r *repository) GetProfile() (*models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile).Error
	return &profile, err
}

func (r *repository) UpdateProfile(profile *models.Profile) error {
	return r.db.Save(profile).Error
}
