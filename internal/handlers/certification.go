package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllCertifications godoc
// @Summary Get all certifications
// @Description Get all certification entries
// @Tags Portfolio - Certifications
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Certification
// @Failure 401 {object} map[string]string
// @Router /portfolio/certifications [get]
func (h *Handler) GetAllCertifications(c *gin.Context) {
	certs, err := h.repo.GetAllCertifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch certifications"})
		return
	}

	c.JSON(http.StatusOK, certs)
}

// GetCertificationByID godoc
// @Summary Get certification by ID
// @Description Get a single certification entry by ID
// @Tags Portfolio - Certifications
// @Produce json
// @Security BearerAuth
// @Param id path int true "Certification ID"
// @Success 200 {object} models.Certification
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/certifications/{id} [get]
func (h *Handler) GetCertificationByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	cert, err := h.repo.GetCertificationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certification not found"})
		return
	}

	c.JSON(http.StatusOK, cert)
}

// CreateCertification godoc
// @Summary Create certification
// @Description Create a new certification entry
// @Tags Portfolio - Certifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param certification body models.Certification true "Certification data"
// @Success 201 {object} models.Certification
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/certifications [post]
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

// UpdateCertification godoc
// @Summary Update certification
// @Description Update an existing certification entry
// @Tags Portfolio - Certifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Certification ID"
// @Param certification body models.Certification true "Certification data"
// @Success 200 {object} models.Certification
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/certifications/{id} [put]
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

// DeleteCertification godoc
// @Summary Delete certification
// @Description Delete a certification entry
// @Tags Portfolio - Certifications
// @Security BearerAuth
// @Param id path int true "Certification ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/certifications/{id} [delete]
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
