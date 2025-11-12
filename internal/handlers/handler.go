package handlers

import (
	"path"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// setLocationHeader sets the Location header for created resources
func setLocationHeader(c *gin.Context, id int64) {
	location := path.Join(c.Request.URL.Path, strconv.FormatInt(id, 10))
	c.Header("Location", location)
}
