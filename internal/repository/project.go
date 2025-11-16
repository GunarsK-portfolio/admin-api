package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/portfolio-common/utils"
)

func (r *repository) GetAllPortfolioProjects(ctx context.Context) ([]models.PortfolioProject, error) {
	var projects []models.PortfolioProject
	err := r.db.WithContext(ctx).Preload("Technologies").Preload("ImageFile").Order("display_order ASC, created_at DESC").Find(&projects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all portfolio projects: %w", err)
	}

	// Populate file URLs using helper
	for i := range projects {
		utils.PopulateFileURL(projects[i].ImageFile, r.filesAPIURL)
	}

	return projects, nil
}

func (r *repository) GetPortfolioProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error) {
	var project models.PortfolioProject
	err := r.db.WithContext(ctx).Preload("Technologies").Preload("ImageFile").First(&project, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get portfolio project with id %d: %w", id, err)
	}

	// Populate file URL using helper
	utils.PopulateFileURL(project.ImageFile, r.filesAPIURL)

	return &project, nil
}

func (r *repository) CreatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(project).Error
	if err != nil {
		return fmt.Errorf("failed to create portfolio project: %w", err)
	}
	return nil
}

func (r *repository) UpdatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	// Use safeUpdateWithAssociations for Technologies many-to-many relationship
	return r.safeUpdateWithAssociations(ctx, project, project.ID)
}

// DeletePortfolioProject deletes a portfolio project and automatically cascades to:
// - portfolio.project_technologies (links to skills/technologies)
// Note: Image file in storage.files is NOT deleted (cleanup job handles orphaned files)
func (r *repository) DeletePortfolioProject(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.PortfolioProject{}, id))
}
