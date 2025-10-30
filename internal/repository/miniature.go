package repository

import (
	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"gorm.io/gorm"
)

func (r *repository) GetAllMiniatureProjects() ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	// Preload Theme and MiniatureFiles (with nested File from storage.files)
	err := r.db.
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC")
		}).
		Preload("MiniatureFiles.File").
		Order("display_order ASC").
		Find(&projects).Error

	// Convert MiniatureFiles to Images for frontend
	for i := range projects {
		projects[i].Images = r.convertMiniatureFilesToImages(projects[i].MiniatureFiles)
	}

	return projects, err
}

func (r *repository) GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	// Preload Theme and MiniatureFiles (with nested File from storage.files)
	err := r.db.
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC")
		}).
		Preload("MiniatureFiles.File").
		First(&project, id).Error

	// Convert MiniatureFiles to Images for frontend
	project.Images = r.convertMiniatureFilesToImages(project.MiniatureFiles)

	return &project, err
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

func (r *repository) CreateMiniatureProject(project *models.MiniatureProject) error {
	return r.db.Omit("ID", "CreatedAt", "UpdatedAt").Create(project).Error
}

func (r *repository) UpdateMiniatureProject(project *models.MiniatureProject) error {
	return r.safeUpdate(project, project.ID)
}

// DeleteMiniatureProject deletes a miniature project and automatically cascades to:
// - miniatures.miniature_files (links to images)
// - miniatures.miniature_techniques (links to techniques)
// - miniatures.miniature_paints (links to paints)
// Note: Actual files in storage.files are NOT deleted (cleanup job handles orphaned files)
func (r *repository) DeleteMiniatureProject(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.MiniatureProject{}, id))
}
