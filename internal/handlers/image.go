package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteImage godoc
// @Summary Delete miniature image
// @Description Delete a miniature file record (removes link between miniature and file)
// @Description Note: This does not delete the actual file from S3/storage
// @Tags Files
// @Security BearerAuth
// @Param id path int true "Miniature File ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /files/{id} [delete]
func (h *Handler) DeleteImage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteImage(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete image"})
		return
	}

	c.Status(http.StatusNoContent)
}
