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

func (r *repository) UpdateProfileAvatar(fileID int64) error {
	return r.db.Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("avatar_file_id", fileID).Error
}

func (r *repository) DeleteProfileAvatar() error {
	return r.db.Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("avatar_file_id", nil).Error
}

func (r *repository) UpdateProfileResume(fileID int64) error {
	return r.db.Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("resume_file_id", fileID).Error
}

func (r *repository) DeleteProfileResume() error {
	return r.db.Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("resume_file_id", nil).Error
}
