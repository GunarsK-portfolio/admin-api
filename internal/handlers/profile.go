package handlers

import (
	"net/http"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// UpdateProfile godoc
// @Summary Update profile
// @Description Update profile information
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body models.Profile true "Profile data"
// @Success 200 {object} models.Profile
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /profile [post]
func (h *Handler) UpdateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateProfile(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
