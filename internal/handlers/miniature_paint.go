package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllMiniaturePaints godoc
// @Summary Get all miniature paints
// @Description Get all miniature paint entries
// @Tags Miniatures - Paints
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MiniaturePaint
// @Failure 401 {object} map[string]string
// @Router /miniatures/paints [get]
func (h *Handler) GetAllMiniaturePaints(c *gin.Context) {
	paints, err := h.repo.GetAllMiniaturePaints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch miniature paints"})
		return
	}

	c.JSON(http.StatusOK, paints)
}

// GetMiniaturePaintByID godoc
// @Summary Get miniature paint by ID
// @Description Get a single miniature paint entry by ID
// @Tags Miniatures - Paints
// @Produce json
// @Security BearerAuth
// @Param id path int true "Paint ID"
// @Success 200 {object} models.MiniaturePaint
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/paints/{id} [get]
func (h *Handler) GetMiniaturePaintByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	paint, err := h.repo.GetMiniaturePaintByID(id)
	if err != nil {
		handleRepositoryError(c, err, "miniature paint not found", "failed to fetch miniature paint")
		return
	}

	c.JSON(http.StatusOK, paint)
}

// CreateMiniaturePaint godoc
// @Summary Create miniature paint
// @Description Create a new miniature paint entry
// @Tags Miniatures - Paints
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param paint body models.MiniaturePaint true "Paint data"
// @Success 201 {object} models.MiniaturePaint
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/paints [post]
func (h *Handler) CreateMiniaturePaint(c *gin.Context) {
	var paint models.MiniaturePaint
	if err := c.ShouldBindJSON(&paint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateMiniaturePaint(&paint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create miniature paint"})
		return
	}

	setLocationHeader(c, paint.ID)
	c.JSON(http.StatusCreated, paint)
}

// UpdateMiniaturePaint godoc
// @Summary Update miniature paint
// @Description Update an existing miniature paint entry
// @Tags Miniatures - Paints
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Paint ID"
// @Param paint body models.MiniaturePaint true "Paint data"
// @Success 200 {object} models.MiniaturePaint
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/paints/{id} [put]
func (h *Handler) UpdateMiniaturePaint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var paint models.MiniaturePaint
	if err := c.ShouldBindJSON(&paint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paint.ID = id
	if err := h.repo.UpdateMiniaturePaint(&paint); err != nil {
		handleRepositoryError(c, err, "miniature paint not found", "failed to update miniature paint")
		return
	}

	c.JSON(http.StatusOK, paint)
}

// DeleteMiniaturePaint godoc
// @Summary Delete miniature paint
// @Description Delete a miniature paint entry
// @Tags Miniatures - Paints
// @Security BearerAuth
// @Param id path int true "Paint ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/paints/{id} [delete]
func (h *Handler) DeleteMiniaturePaint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteMiniaturePaint(id); err != nil {
		handleRepositoryError(c, err, "miniature paint not found", "failed to delete miniature paint")
		return
	}

	c.Status(http.StatusNoContent)
}
