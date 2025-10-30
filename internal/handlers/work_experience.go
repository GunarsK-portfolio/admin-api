package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllWorkExperience godoc
// @Summary Get all work experience
// @Description Get all work experience entries
// @Tags Portfolio - Experience
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.WorkExperience
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /portfolio/experience [get]
func (h *Handler) GetAllWorkExperience(c *gin.Context) {
	experiences, err := h.repo.GetAllWorkExperience(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch work experience"})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

// GetWorkExperienceByID godoc
// @Summary Get work experience by ID
// @Description Get a single work experience entry by ID
// @Tags Portfolio - Experience
// @Produce json
// @Security BearerAuth
// @Param id path int true "Work Experience ID"
// @Success 200 {object} models.WorkExperience
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/experience/{id} [get]
func (h *Handler) GetWorkExperienceByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	exp, err := h.repo.GetWorkExperienceByID(c.Request.Context(), id)
	if err != nil {
		handleRepositoryError(c, err, "work experience not found", "failed to fetch work experience")
		return
	}

	c.JSON(http.StatusOK, exp)
}

// CreateWorkExperience godoc
// @Summary Create work experience
// @Description Create a new work experience entry
// @Tags Portfolio - Experience
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param experience body models.WorkExperience true "Work experience data"
// @Success 201 {object} models.WorkExperience
// @Header 201 {string} Location "URL of the created resource"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/experience [post]
func (h *Handler) CreateWorkExperience(c *gin.Context) {
	var exp models.WorkExperience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateWorkExperience(c.Request.Context(), &exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create work experience"})
		return
	}

	setLocationHeader(c, exp.ID)
	c.JSON(http.StatusCreated, exp)
}

// UpdateWorkExperience godoc
// @Summary Update work experience
// @Description Update an existing work experience entry
// @Tags Portfolio - Experience
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Work Experience ID"
// @Param experience body models.WorkExperience true "Work experience data"
// @Success 200 {object} models.WorkExperience
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/experience/{id} [put]
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
	if err := h.repo.UpdateWorkExperience(c.Request.Context(), &exp); err != nil {
		handleRepositoryError(c, err, "work experience not found", "failed to update work experience")
		return
	}

	c.JSON(http.StatusOK, exp)
}

// DeleteWorkExperience godoc
// @Summary Delete work experience
// @Description Delete a work experience entry
// @Tags Portfolio - Experience
// @Security BearerAuth
// @Param id path int true "Work Experience ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/experience/{id} [delete]
func (h *Handler) DeleteWorkExperience(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteWorkExperience(c.Request.Context(), id); err != nil {
		handleRepositoryError(c, err, "work experience not found", "failed to delete work experience")
		return
	}

	c.Status(http.StatusNoContent)
}
