package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

func (r *repository) GetAllMiniatureThemes() ([]models.MiniatureTheme, error) {
	var themes []models.MiniatureTheme
	err := r.db.Order("display_order ASC, name ASC").Find(&themes).Error
	return themes, err
}

func (r *repository) GetMiniatureThemeByID(id int64) (*models.MiniatureTheme, error) {
	var theme models.MiniatureTheme
	err := r.db.Preload("Miniatures").First(&theme, id).Error
	return &theme, err
}

func (r *repository) CreateMiniatureTheme(theme *models.MiniatureTheme) error {
	return r.db.Omit("ID", "CreatedAt", "UpdatedAt").Create(theme).Error
}

func (r *repository) UpdateMiniatureTheme(theme *models.MiniatureTheme) error {
	return r.safeUpdate(theme, theme.ID)
}

func (r *repository) DeleteMiniatureTheme(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.MiniatureTheme{}, id))
}
