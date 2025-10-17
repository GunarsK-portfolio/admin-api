package handlers

import "github.com/GunarsK-portfolio/admin-api/internal/repository"

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}
