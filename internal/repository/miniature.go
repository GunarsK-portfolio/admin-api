package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"gorm.io/gorm"
)

func (r *repository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	// Preload Theme and MiniatureFiles (with nested File from storage.files)
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC")
		}).
		Preload("MiniatureFiles.File").
		Order("display_order ASC").
		Find(&projects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature projects: %w", err)
	}

	// Convert MiniatureFiles to Images for frontend
	for i := range projects {
		projects[i].Images = r.convertMiniatureFilesToImages(projects[i].MiniatureFiles)
	}

	return projects, nil
}

func (r *repository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	// Preload Theme and MiniatureFiles (with nested File from storage.files)
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC")
		}).
		Preload("MiniatureFiles.File").
		First(&project, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature project with id %d: %w", id, err)
	}

	// Convert MiniatureFiles to Images for frontend
	project.Images = r.convertMiniatureFilesToImages(project.MiniatureFiles)

	return &project, nil
}

// Helper function to convert MiniatureFiles to simplified Images for frontend
func (r *repository) convertMiniatureFilesToImages(files []models.MiniatureFile) []models.Image {
	images := make([]models.Image, 0, len(files))
	for _, file := range files {
		if file.File != nil {
			images = append(images, models.Image{
				ID:      file.ID,
				URL:     utils.BuildFileURL(r.filesAPIURL, file.File.FileType, file.File.S3Key),
				Caption: file.Caption,
			})
		}
	}
	return images
}

func (r *repository) CreateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(project).Error
	if err != nil {
		return fmt.Errorf("failed to create miniature project: %w", err)
	}
	return nil
}

func (r *repository) UpdateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	return r.safeUpdate(ctx, project, project.ID)
}

// DeleteMiniatureProject deletes a miniature project and automatically cascades to:
// - miniatures.miniature_files (links to images)
// - miniatures.miniature_techniques (links to techniques)
// - miniatures.miniature_paints (links to paints)
// Note: Actual files in storage.files are NOT deleted (cleanup job handles orphaned files)
func (r *repository) DeleteMiniatureProject(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.MiniatureProject{}, id))
}
