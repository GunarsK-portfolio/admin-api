package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

func (r *repository) GetAllMiniaturePaints() ([]models.MiniaturePaint, error) {
	var paints []models.MiniaturePaint
	err := r.db.Order("manufacturer ASC, name ASC").Find(&paints).Error
	return paints, err
}

func (r *repository) GetMiniaturePaintByID(id int64) (*models.MiniaturePaint, error) {
	var paint models.MiniaturePaint
	err := r.db.First(&paint, id).Error
	return &paint, err
}

func (r *repository) CreateMiniaturePaint(paint *models.MiniaturePaint) error {
	return r.db.Create(paint).Error
}

func (r *repository) UpdateMiniaturePaint(paint *models.MiniaturePaint) error {
	return checkRowsAffected(r.db.Save(paint))
}

func (r *repository) DeleteMiniaturePaint(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.MiniaturePaint{}, id))
}
