package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// Profile handlers

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

// Work Experience handlers

// CreateWorkExperience godoc
// @Summary Create work experience
// @Description Create a new work experience entry
// @Tags experience
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param experience body models.WorkExperience true "Work experience data"
// @Success 201 {object} models.WorkExperience
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /experience [post]
func (h *Handler) CreateWorkExperience(c *gin.Context) {
	var exp models.WorkExperience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateWorkExperience(&exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create work experience"})
		return
	}

	c.JSON(http.StatusCreated, exp)
}

// UpdateWorkExperience godoc
// @Summary Update work experience
// @Description Update an existing work experience entry
// @Tags experience
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Work Experience ID"
// @Param experience body models.WorkExperience true "Work experience data"
// @Success 200 {object} models.WorkExperience
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /experience/{id} [put]
func (h *Handler) UpdateWorkExperience(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var exp models.WorkExperience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exp.ID = id
	if err := h.repo.UpdateWorkExperience(&exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update work experience"})
		return
	}

	c.JSON(http.StatusOK, exp)
}

// DeleteWorkExperience godoc
// @Summary Delete work experience
// @Description Delete a work experience entry
// @Tags experience
// @Security BearerAuth
// @Param id path int true "Work Experience ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /experience/{id} [delete]
func (h *Handler) DeleteWorkExperience(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteWorkExperience(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete work experience"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Certification handlers (similar pattern)

func (h *Handler) CreateCertification(c *gin.Context) {
	var cert models.Certification
	if err := c.ShouldBindJSON(&cert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateCertification(&cert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create certification"})
		return
	}

	c.JSON(http.StatusCreated, cert)
}

func (h *Handler) UpdateCertification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var cert models.Certification
	if err := c.ShouldBindJSON(&cert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cert.ID = id
	if err := h.repo.UpdateCertification(&cert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update certification"})
		return
	}

	c.JSON(http.StatusOK, cert)
}

func (h *Handler) DeleteCertification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteCertification(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete certification"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Miniature Project handlers (similar pattern)

func (h *Handler) CreateMiniatureProject(c *gin.Context) {
	var project models.MiniatureProject
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateMiniatureProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create miniature project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *Handler) UpdateMiniatureProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var project models.MiniatureProject
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.ID = id
	if err := h.repo.UpdateMiniatureProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update miniature project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *Handler) DeleteMiniatureProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteMiniatureProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete miniature project"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Image handlers (simplified - just URL storage, actual upload handled separately)

func (h *Handler) DeleteImage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteImage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete image"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Health check

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
