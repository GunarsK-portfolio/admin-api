package handlers

import (
	"net/http"
	"strconv"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllPortfolioProjects godoc
// @Summary Get all portfolio projects
// @Description Get all portfolio projects
// @Tags Portfolio - Projects
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.PortfolioProject
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /portfolio/projects [get]
func (h *Handler) GetAllPortfolioProjects(c *gin.Context) {
	projects, err := h.repo.GetAllPortfolioProjects(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch portfolio projects")
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetPortfolioProjectByID godoc
// @Summary Get portfolio project by ID
// @Description Get a single portfolio project by ID
// @Tags Portfolio - Projects
// @Produce json
// @Security BearerAuth
// @Param id path int true "Portfolio Project ID"
// @Success 200 {object} models.PortfolioProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/projects/{id} [get]
func (h *Handler) GetPortfolioProjectByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	project, err := h.repo.GetPortfolioProjectByID(c.Request.Context(), id)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "portfolio project not found", "failed to fetch portfolio project")
		return
	}

	c.JSON(http.StatusOK, project)
}

// CreatePortfolioProject godoc
// @Summary Create portfolio project
// @Description Create a new portfolio project
// @Tags Portfolio - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project body models.PortfolioProject true "Portfolio project data"
// @Success 201 {object} models.PortfolioProject
// @Header 201 {string} Location "URL of the created resource"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/projects [post]
func (h *Handler) CreatePortfolioProject(c *gin.Context) {
	var project models.PortfolioProject
	if err := c.ShouldBindJSON(&project); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.CreatePortfolioProject(c.Request.Context(), &project); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "", "failed to create portfolio project")
		return
	}

	setLocationHeader(c, project.ID)
	c.JSON(http.StatusCreated, project)
}

// UpdatePortfolioProject godoc
// @Summary Update portfolio project
// @Description Update an existing portfolio project
// @Tags Portfolio - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Portfolio Project ID"
// @Param project body models.PortfolioProject true "Portfolio project data"
// @Success 200 {object} models.PortfolioProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/projects/{id} [put]
func (h *Handler) UpdatePortfolioProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var project models.PortfolioProject
	if err := c.ShouldBindJSON(&project); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	project.ID = id
	if err := h.repo.UpdatePortfolioProject(c.Request.Context(), &project); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "portfolio project not found", "failed to update portfolio project")
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeletePortfolioProject godoc
// @Summary Delete portfolio project
// @Description Delete a portfolio project and all associated technology links
// @Description Note: This is a cascade delete - one API call deletes the project and all technology associations
// @Description The project image file in S3 is preserved and cleaned up by background job
// @Tags Portfolio - Projects
// @Security BearerAuth
// @Param id path int true "Portfolio Project ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/projects/{id} [delete]
func (h *Handler) DeletePortfolioProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.repo.DeletePortfolioProject(c.Request.Context(), id); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "portfolio project not found", "failed to delete portfolio project")
		return
	}

	c.Status(http.StatusNoContent)
}
