package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllMiniatureProjects godoc
// @Summary Get all miniature projects
// @Description Get all miniature painting projects
// @Tags Miniatures - Projects
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MiniatureProject
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /miniatures/projects [get]
func (h *Handler) GetAllMiniatureProjects(c *gin.Context) {
	projects, err := h.repo.GetAllMiniatureProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch miniature projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetMiniatureProjectByID godoc
// @Summary Get miniature project by ID
// @Description Get a single miniature project by ID
// @Tags Miniatures - Projects
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id} [get]
func (h *Handler) GetMiniatureProjectByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	project, err := h.repo.GetMiniatureProjectByID(c.Request.Context(), id)
	if err != nil {
		handleRepositoryError(c, err, "miniature project not found", "failed to fetch miniature project")
		return
	}

	c.JSON(http.StatusOK, project)
}

// CreateMiniatureProject godoc
// @Summary Create miniature project
// @Description Create a new miniature painting project
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project body models.MiniatureProject true "Miniature project data"
// @Success 201 {object} models.MiniatureProject
// @Header 201 {string} Location "URL of the created resource"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects [post]
func (h *Handler) CreateMiniatureProject(c *gin.Context) {
	var project models.MiniatureProject
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateMiniatureProject(c.Request.Context(), &project); err != nil {
		handleRepositoryError(c, err, "", "failed to create miniature project")
		return
	}

	setLocationHeader(c, project.ID)
	c.JSON(http.StatusCreated, project)
}

// UpdateMiniatureProject godoc
// @Summary Update miniature project
// @Description Update an existing miniature project
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Param project body models.MiniatureProject true "Miniature project data"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id} [put]
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
	if err := h.repo.UpdateMiniatureProject(c.Request.Context(), &project); err != nil {
		handleRepositoryError(c, err, "miniature project not found", "failed to update miniature project")
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteMiniatureProject godoc
// @Summary Delete miniature project
// @Description Delete a miniature project and all associated data (images, techniques, paints)
// @Description Note: This is a cascade delete - one API call deletes everything
// @Description Actual image files in S3 are preserved and cleaned up by background job
// @Tags Miniatures - Projects
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id} [delete]
func (h *Handler) DeleteMiniatureProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteMiniatureProject(c.Request.Context(), id); err != nil {
		handleRepositoryError(c, err, "miniature project not found", "failed to delete miniature project")
		return
	}

	c.Status(http.StatusNoContent)
}
