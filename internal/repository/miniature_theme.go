package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	var themes []models.MiniatureTheme
	err := r.db.WithContext(ctx).Order("display_order ASC, name ASC").Find(&themes).Error
	return themes, err
}

func (r *repository) GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error) {
	var theme models.MiniatureTheme
	err := r.db.WithContext(ctx).Preload("Miniatures").First(&theme, id).Error
	return &theme, err
}

func (r *repository) CreateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	return r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(theme).Error
}

func (r *repository) UpdateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	return r.safeUpdate(ctx, theme, theme.ID)
}

func (r *repository) DeleteMiniatureTheme(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.MiniatureTheme{}, id))
}
