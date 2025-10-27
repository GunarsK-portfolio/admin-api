package repository

import (
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetProfile() (*models.Profile, error) {
	var profile models.Profile
	err := r.db.
		Preload("AvatarFile").
		Preload("ResumeFile").
		First(&profile).Error
	if err != nil {
		return nil, err
	}

	// Construct avatar URL
	if profile.AvatarFile != nil {
		profile.AvatarURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, profile.AvatarFile.FileType, profile.AvatarFile.S3Key)
	}

	// Construct resume URL and filename
	if profile.ResumeFile != nil {
		profile.ResumeURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, profile.ResumeFile.FileType, profile.ResumeFile.S3Key)
		profile.ResumeFileName = profile.ResumeFile.FileName
	}

	return &profile, nil
}

func (r *repository) UpdateProfile(profile *models.Profile) error {
	// Upsert: update if exists, insert if doesn't (singleton pattern)
	var existing models.Profile
	err := r.db.First(&existing).Error

	if err != nil {
		// No profile exists, create the first one
		return r.db.Create(profile).Error
	}

	// Profile exists, update it
	return r.db.Model(&existing).Updates(map[string]interface{}{
		"full_name":      profile.FullName,
		"title":          profile.Title,
		"bio":            profile.Bio,
		"email":          profile.Email,
		"phone":          profile.Phone,
		"location":       profile.Location,
		"avatar_file_id": profile.AvatarFileID,
		"resume_file_id": profile.ResumeFileID,
	}).Error
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
