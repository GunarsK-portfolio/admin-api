package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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
	// Log internal errors for debugging
	log.Printf("Internal error in %s %s: %v", c.Request.Method, c.Request.URL.Path, err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": internalMsg})
}

// setLocationHeader sets the Location header for created resources
func setLocationHeader(c *gin.Context, id int64) {
	c.Header("Location", c.Request.URL.Path+"/"+strconv.FormatInt(id, 10))
}
