package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	var themes []models.MiniatureTheme
	err := r.db.WithContext(ctx).Order("display_order ASC, name ASC").Find(&themes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature themes: %w", err)
	}
	return themes, nil
}

func (r *repository) GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error) {
	var theme models.MiniatureTheme
	err := r.db.WithContext(ctx).Preload("Miniatures").First(&theme, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature theme with id %d: %w", id, err)
	}
	return &theme, nil
}

func (r *repository) CreateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(theme).Error
	if err != nil {
		return fmt.Errorf("failed to create miniature theme: %w", err)
	}
	return nil
}

func (r *repository) UpdateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	return r.safeUpdate(ctx, theme, theme.ID)
}

func (r *repository) DeleteMiniatureTheme(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.MiniatureTheme{}, id))
}
