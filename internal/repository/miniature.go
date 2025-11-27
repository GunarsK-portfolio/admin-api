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
	// Preload Theme, MiniatureFiles, Techniques, and Paints
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC, miniatures.miniature_files.id ASC")
		}).
		Preload("MiniatureFiles.File").
		Preload("Techniques.Technique").
		Preload("Paints.Paint").
		Order("display_order ASC, id ASC").
		Find(&projects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature projects: %w", err)
	}

	// Build image URLs (requires filesAPIURL)
	for i := range projects {
		projects[i].Images = utils.ConvertMiniatureFilesToImages(projects[i].MiniatureFiles, r.filesAPIURL)
	}

	return projects, nil
}

func (r *repository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	// Preload Theme, MiniatureFiles, Techniques, and Paints
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC, miniatures.miniature_files.id ASC")
		}).
		Preload("MiniatureFiles.File").
		Preload("Techniques.Technique").
		Preload("Paints.Paint").
		First(&project, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature project with id %d: %w", id, err)
	}

	// Build image URLs (requires filesAPIURL)
	project.Images = utils.ConvertMiniatureFilesToImages(project.MiniatureFiles, r.filesAPIURL)

	return &project, nil
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

// AddImageToProject links an uploaded file to a miniature project
// Display order is auto-assigned based on the current maximum order + 1
func (r *repository) AddImageToProject(ctx context.Context, miniatureFile *models.MiniatureFile) error {
	// First verify the project exists
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.MiniatureProject{}).
		Where("id = ?", miniatureFile.MiniatureProjectID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to verify project: %w", err)
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	// Auto-assign display order (get max + 1)
	var maxOrder int
	r.db.WithContext(ctx).Model(&models.MiniatureFile{}).
		Where("miniature_project_id = ?", miniatureFile.MiniatureProjectID).
		Select("COALESCE(MAX(display_order), -1)").
		Scan(&maxOrder)
	miniatureFile.DisplayOrder = maxOrder + 1

	// Create the miniature_files record
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt").Create(miniatureFile).Error
	if err != nil {
		return fmt.Errorf("failed to add image to project: %w", err)
	}

	// Reload with file data
	err = r.db.WithContext(ctx).Preload("File").First(miniatureFile, miniatureFile.ID).Error
	if err != nil {
		return fmt.Errorf("failed to reload image data: %w", err)
	}

	return nil
}

// SetProjectTechniques replaces all techniques for a project
func (r *repository) SetProjectTechniques(ctx context.Context, projectID int64, techniqueIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete existing techniques
		if err := tx.Where("miniature_project_id = ?", projectID).
			Delete(&models.MiniatureProjectTechnique{}).Error; err != nil {
			return fmt.Errorf("failed to clear techniques: %w", err)
		}

		// Insert new techniques
		for _, techniqueID := range techniqueIDs {
			link := models.MiniatureProjectTechnique{
				MiniatureProjectID: projectID,
				TechniqueID:        techniqueID,
			}
			if err := tx.Omit("ID", "CreatedAt").Create(&link).Error; err != nil {
				return fmt.Errorf("failed to add technique %d: %w", techniqueID, err)
			}
		}
		return nil
	})
}

// SetProjectPaints replaces all paints for a project
func (r *repository) SetProjectPaints(ctx context.Context, projectID int64, paintIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete existing paints
		if err := tx.Where("miniature_project_id = ?", projectID).
			Delete(&models.MiniatureProjectPaint{}).Error; err != nil {
			return fmt.Errorf("failed to clear paints: %w", err)
		}

		// Insert new paints
		for _, paintID := range paintIDs {
			link := models.MiniatureProjectPaint{
				MiniatureProjectID: projectID,
				PaintID:            paintID,
			}
			if err := tx.Omit("ID", "CreatedAt").Create(&link).Error; err != nil {
				return fmt.Errorf("failed to add paint %d: %w", paintID, err)
			}
		}
		return nil
	})
}

// GetAllTechniques returns all techniques from the classifier table
func (r *repository) GetAllTechniques(ctx context.Context) ([]models.MiniatureTechnique, error) {
	var techniques []models.MiniatureTechnique
	err := r.db.WithContext(ctx).Order("display_order ASC, name ASC").Find(&techniques).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get techniques: %w", err)
	}
	return techniques, nil
}
