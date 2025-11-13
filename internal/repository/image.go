package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

// DeleteImage deletes a miniature file record (junction table entry)
// NOTE: This deletes the link between a miniature and a file, not the actual file in S3
// The actual file in storage.files remains (can be cleaned up separately if orphaned)
func (r *repository) DeleteImage(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&models.MiniatureFile{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete image with id %d: %w", id, err)
	}
	return nil
}
