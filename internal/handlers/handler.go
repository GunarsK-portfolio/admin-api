package handlers

import (
	"github.com/GunarsK-portfolio/admin-api/internal/repository"
	commonhandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// setLocationHeader wraps the common helper for backward compatibility
var setLocationHeader = commonhandlers.SetLocationHeader
