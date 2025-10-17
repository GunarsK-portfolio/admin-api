package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

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
