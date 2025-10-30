package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllMiniatureThemes godoc
// @Summary Get all miniature themes
// @Description Get all miniature painting themes
// @Tags Miniatures - Themes
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MiniatureTheme
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /miniatures/themes [get]
func (h *Handler) GetAllMiniatureThemes(c *gin.Context) {
	themes, err := h.repo.GetAllMiniatureThemes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch miniature themes"})
		return
	}

	c.JSON(http.StatusOK, themes)
}

// GetMiniatureThemeByID godoc
// @Summary Get miniature theme by ID
// @Description Get a single miniature theme by ID
// @Tags Miniatures - Themes
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Theme ID"
// @Success 200 {object} models.MiniatureTheme
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/themes/{id} [get]
func (h *Handler) GetMiniatureThemeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	theme, err := h.repo.GetMiniatureThemeByID(c.Request.Context(), id)
	if err != nil {
		handleRepositoryError(c, err, "miniature theme not found", "failed to fetch miniature theme")
		return
	}

	c.JSON(http.StatusOK, theme)
}

// CreateMiniatureTheme godoc
// @Summary Create miniature theme
// @Description Create a new miniature painting theme
// @Tags Miniatures - Themes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param theme body models.MiniatureTheme true "Miniature theme data"
// @Success 201 {object} models.MiniatureTheme
// @Header 201 {string} Location "URL of the created resource"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/themes [post]
func (h *Handler) CreateMiniatureTheme(c *gin.Context) {
	var theme models.MiniatureTheme
	if err := c.ShouldBindJSON(&theme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateMiniatureTheme(c.Request.Context(), &theme); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create miniature theme"})
		return
	}

	setLocationHeader(c, theme.ID)
	c.JSON(http.StatusCreated, theme)
}

// UpdateMiniatureTheme godoc
// @Summary Update miniature theme
// @Description Update an existing miniature theme
// @Tags Miniatures - Themes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Miniature Theme ID"
// @Param theme body models.MiniatureTheme true "Miniature theme data"
// @Success 200 {object} models.MiniatureTheme
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/themes/{id} [put]
func (h *Handler) UpdateMiniatureTheme(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var theme models.MiniatureTheme
	if err := c.ShouldBindJSON(&theme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	theme.ID = id
	if err := h.repo.UpdateMiniatureTheme(c.Request.Context(), &theme); err != nil {
		handleRepositoryError(c, err, "miniature theme not found", "failed to update miniature theme")
		return
	}

	c.JSON(http.StatusOK, theme)
}

// DeleteMiniatureTheme godoc
// @Summary Delete miniature theme
// @Description Delete a miniature theme
// @Tags Miniatures - Themes
// @Security BearerAuth
// @Param id path int true "Miniature Theme ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /miniatures/themes/{id} [delete]
func (h *Handler) DeleteMiniatureTheme(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteMiniatureTheme(c.Request.Context(), id); err != nil {
		handleRepositoryError(c, err, "miniature theme not found", "failed to delete miniature theme")
		return
	}

	c.Status(http.StatusNoContent)
}
