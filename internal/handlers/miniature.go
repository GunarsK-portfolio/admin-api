package handlers

import (
	"net/http"
	"strconv"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"

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
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch miniature projects")
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
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	project, err := h.repo.GetMiniatureProjectByID(c.Request.Context(), id)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "miniature project not found", "failed to fetch miniature project")
		return
	}

	c.JSON(http.StatusOK, project)
}

// miniatureProjectRequest is the request body for create/update miniature project
type miniatureProjectRequest struct {
	models.MiniatureProject
	TechniqueIDs []int64 `json:"techniqueIds,omitempty"`
	PaintIDs     []int64 `json:"paintIds,omitempty"`
}

// CreateMiniatureProject godoc
// @Summary Create miniature project
// @Description Create a new miniature painting project with optional techniques and paints
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project body miniatureProjectRequest true "Miniature project data with optional techniqueIds and paintIds"
// @Success 201 {object} models.MiniatureProject
// @Header 201 {string} Location "URL of the created resource"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects [post]
func (h *Handler) CreateMiniatureProject(c *gin.Context) {
	var req miniatureProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	if err := h.repo.CreateMiniatureProject(ctx, &req.MiniatureProject); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "", "failed to create miniature project")
		return
	}

	// Set techniques and paints if provided
	if len(req.TechniqueIDs) > 0 {
		if err := h.repo.SetProjectTechniques(ctx, req.ID, req.TechniqueIDs); err != nil {
			commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to set techniques")
			return
		}
	}
	if len(req.PaintIDs) > 0 {
		if err := h.repo.SetProjectPaints(ctx, req.ID, req.PaintIDs); err != nil {
			commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to set paints")
			return
		}
	}

	// Reload project with all associations
	project, err := h.repo.GetMiniatureProjectByID(ctx, req.ID)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to fetch created project")
		return
	}

	setLocationHeader(c, project.ID)
	c.JSON(http.StatusCreated, project)
}

// UpdateMiniatureProject godoc
// @Summary Update miniature project
// @Description Update an existing miniature project with optional techniques and paints
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Param project body miniatureProjectRequest true "Miniature project data with optional techniqueIds and paintIds"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id} [put]
func (h *Handler) UpdateMiniatureProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req miniatureProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	req.ID = id

	if err := h.repo.UpdateMiniatureProject(ctx, &req.MiniatureProject); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "miniature project not found", "failed to update miniature project")
		return
	}

	// Update techniques and paints (always replace, even with empty arrays)
	if err := h.repo.SetProjectTechniques(ctx, id, req.TechniqueIDs); err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to set techniques")
		return
	}
	if err := h.repo.SetProjectPaints(ctx, id, req.PaintIDs); err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to set paints")
		return
	}

	// Reload project with all associations
	project, err := h.repo.GetMiniatureProjectByID(ctx, id)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to fetch updated project")
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
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.repo.DeleteMiniatureProject(c.Request.Context(), id); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "miniature project not found", "failed to delete miniature project")
		return
	}

	c.Status(http.StatusNoContent)
}

// AddImageToProject godoc
// @Summary Add image to miniature project
// @Description Link an uploaded image file to a miniature project with optional caption
// @Description Display order is automatically assigned based on upload order
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Param image body object{fileId=int64,caption=string} true "Image data (fileId required, caption optional)"
// @Success 201 {object} models.MiniatureFile
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /miniatures/projects/{id}/images [post]
func (h *Handler) AddImageToProject(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid project id")
		return
	}

	var req struct {
		FileID  int64  `json:"fileId" binding:"required"`
		Caption string `json:"caption"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	miniatureFile := &models.MiniatureFile{
		MiniatureProjectID: projectID,
		FileID:             req.FileID,
		Caption:            req.Caption,
	}

	if err := h.repo.AddImageToProject(c.Request.Context(), miniatureFile); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to add image to project")
		return
	}

	c.JSON(http.StatusCreated, miniatureFile)
}

// SetProjectTechniques godoc
// @Summary Set techniques for a miniature project
// @Description Replace all techniques for a project with the provided list
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Param techniques body object{techniqueIds=[]int64} true "List of technique IDs"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id}/techniques [put]
func (h *Handler) SetProjectTechniques(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid project id")
		return
	}

	var req struct {
		TechniqueIDs []int64 `json:"techniqueIds"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.SetProjectTechniques(c.Request.Context(), projectID, req.TechniqueIDs); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to set techniques")
		return
	}

	// Return updated project
	project, err := h.repo.GetMiniatureProjectByID(c.Request.Context(), projectID)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to fetch project")
		return
	}
	c.JSON(http.StatusOK, project)
}

// SetProjectPaints godoc
// @Summary Set paints for a miniature project
// @Description Replace all paints for a project with the provided list
// @Tags Miniatures - Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Project ID"
// @Param paints body object{paintIds=[]int64} true "List of paint IDs"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/projects/{id}/paints [put]
func (h *Handler) SetProjectPaints(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid project id")
		return
	}

	var req struct {
		PaintIDs []int64 `json:"paintIds"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.SetProjectPaints(c.Request.Context(), projectID, req.PaintIDs); err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to set paints")
		return
	}

	// Return updated project
	project, err := h.repo.GetMiniatureProjectByID(c.Request.Context(), projectID)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to fetch project")
		return
	}
	c.JSON(http.StatusOK, project)
}

// GetAllTechniques godoc
// @Summary Get all techniques
// @Description Get all painting techniques from the classifier table
// @Tags Miniatures - Techniques
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MiniatureTechnique
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /miniatures/techniques [get]
func (h *Handler) GetAllTechniques(c *gin.Context) {
	techniques, err := h.repo.GetAllTechniques(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch techniques")
		return
	}
	c.JSON(http.StatusOK, techniques)
}
