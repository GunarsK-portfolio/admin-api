package handlers

import (
	"errors"
	"net/http"
	"path"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/GunarsK-portfolio/portfolio-common/logger"
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
	// Log internal errors with structured logging
	logger.GetLogger(c).Error("Repository error",
		"error", err,
		"message", internalMsg,
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
	)
	c.JSON(http.StatusInternalServerError, gin.H{"error": internalMsg})
}

// setLocationHeader sets the Location header for created resources
func setLocationHeader(c *gin.Context, id int64) {
	location := path.Join(c.Request.URL.Path, strconv.FormatInt(id, 10))
	c.Header("Location", location)
}
