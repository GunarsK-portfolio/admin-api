package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

func (r *repository) GetAllMiniaturePaints(ctx context.Context) ([]models.MiniaturePaint, error) {
	var paints []models.MiniaturePaint
	err := r.db.WithContext(ctx).Order("manufacturer ASC, name ASC").Find(&paints).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature paints: %w", err)
	}
	return paints, nil
}

func (r *repository) GetMiniaturePaintByID(ctx context.Context, id int64) (*models.MiniaturePaint, error) {
	var paint models.MiniaturePaint
	err := r.db.WithContext(ctx).First(&paint, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature paint with id %d: %w", id, err)
	}
	return &paint, nil
}

func (r *repository) CreateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(paint).Error
	if err != nil {
		return fmt.Errorf("failed to create miniature paint: %w", err)
	}
	return nil
}

func (r *repository) UpdateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	return r.safeUpdate(ctx, paint, paint.ID)
}

func (r *repository) DeleteMiniaturePaint(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.MiniaturePaint{}, id))
}
