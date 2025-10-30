package handlers

import (
	"net/http"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get profile
// @Description Get profile information
// @Tags Portfolio - Profile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Profile
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	profile, err := h.repo.GetProfile()
	if err != nil {
		handleRepositoryError(c, err, "profile not found", "failed to fetch profile")
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile godoc
// @Summary Update profile
// @Description Update profile information
// @Tags Portfolio - Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body models.Profile true "Profile data"
// @Success 200 {object} models.Profile
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile [put]
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

// UpdateProfileAvatar godoc
// @Summary Update profile avatar
// @Description Update profile avatar by file ID
// @Tags Portfolio - Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fileId body object{fileId=int64} true "File ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile/avatar [put]
func (h *Handler) UpdateProfileAvatar(c *gin.Context) {
	var request struct {
		FileID int64 `json:"fileId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateProfileAvatar(request.FileID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update avatar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "avatar updated successfully"})
}

// DeleteProfileAvatar godoc
// @Summary Delete profile avatar
// @Description Remove profile avatar (sets avatar_file_id to NULL)
// @Tags Portfolio - Profile
// @Security BearerAuth
// @Success 204
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile/avatar [delete]
func (h *Handler) DeleteProfileAvatar(c *gin.Context) {
	if err := h.repo.DeleteProfileAvatar(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete avatar"})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateProfileResume godoc
// @Summary Update profile resume
// @Description Update profile resume by file ID
// @Tags Portfolio - Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fileId body object{fileId=int64} true "File ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile/resume [put]
func (h *Handler) UpdateProfileResume(c *gin.Context) {
	var request struct {
		FileID int64 `json:"fileId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateProfileResume(request.FileID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update resume"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resume updated successfully"})
}

// DeleteProfileResume godoc
// @Summary Delete profile resume
// @Description Remove profile resume (sets resume_file_id to NULL)
// @Tags Portfolio - Profile
// @Security BearerAuth
// @Success 204
// @Failure 401 {object} map[string]string
// @Router /portfolio/profile/resume [delete]
func (h *Handler) DeleteProfileResume(c *gin.Context) {
	if err := h.repo.DeleteProfileResume(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete resume"})
		return
	}

	c.Status(http.StatusNoContent)
}
