package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

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
