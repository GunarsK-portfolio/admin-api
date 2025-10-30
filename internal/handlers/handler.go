package handlers

import (
	"errors"
	"net/http"

	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// handleRepositoryError checks if the error is a record not found error
// and responds with appropriate HTTP status code (404 or 500)
func handleRepositoryError(c *gin.Context, err error, notFoundMsg, internalMsg string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMsg})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": internalMsg})
}
