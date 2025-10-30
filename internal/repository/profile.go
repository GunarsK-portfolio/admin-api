package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/portfolio-common/utils"
)

func (r *repository) GetProfile(ctx context.Context) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.WithContext(ctx).
		Preload("AvatarFile").
		Preload("ResumeFile").
		First(&profile).Error
	if err != nil {
		return nil, err
	}

	// Populate file URLs using helper
	utils.PopulateFileURL(profile.AvatarFile, r.filesAPIURL)
	utils.PopulateFileURL(profile.ResumeFile, r.filesAPIURL)

	return &profile, nil
}

func (r *repository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	// Upsert: update if exists, insert if doesn't (singleton pattern)
	var existing models.Profile
	err := r.db.WithContext(ctx).First(&existing).Error

	if err != nil {
		// No profile exists, create the first one
		return r.db.WithContext(ctx).Create(profile).Error
	}

	// Profile exists, update it
	return r.db.WithContext(ctx).Model(&existing).Updates(map[string]interface{}{
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

func (r *repository) UpdateProfileAvatar(ctx context.Context, fileID int64) error {
	return r.db.WithContext(ctx).Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("avatar_file_id", fileID).Error
}

func (r *repository) DeleteProfileAvatar(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("avatar_file_id", nil).Error
}

func (r *repository) UpdateProfileResume(ctx context.Context, fileID int64) error {
	return r.db.WithContext(ctx).Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("resume_file_id", fileID).Error
}

func (r *repository) DeleteProfileResume(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&models.Profile{}).
		Where("id = (SELECT MIN(id) FROM portfolio.profile)").
		Update("resume_file_id", nil).Error
}
