package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
)

// SKILLS

// GetAllSkills godoc
// @Summary Get all skills
// @Description Get all skills
// @Tags Portfolio - Skills
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Skill
// @Failure 401 {object} map[string]string
// @Router /portfolio/skills [get]
func (h *Handler) GetAllSkills(c *gin.Context) {
	skills, err := h.repo.GetAllSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch skills"})
		return
	}

	c.JSON(http.StatusOK, skills)
}

// GetSkillByID godoc
// @Summary Get skill by ID
// @Description Get a single skill by ID
// @Tags Portfolio - Skills
// @Produce json
// @Security BearerAuth
// @Param id path int true "Skill ID"
// @Success 200 {object} models.Skill
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skills/{id} [get]
func (h *Handler) GetSkillByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	skill, err := h.repo.GetSkillByID(id)
	if err != nil {
		handleRepositoryError(c, err, "skill not found", "failed to fetch skill")
		return
	}

	c.JSON(http.StatusOK, skill)
}

// CreateSkill godoc
// @Summary Create skill
// @Description Create a new skill
// @Tags Portfolio - Skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param skill body models.Skill true "Skill data"
// @Success 201 {object} models.Skill
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skills [post]
func (h *Handler) CreateSkill(c *gin.Context) {
	var skill models.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateSkill(&skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create skill"})
		return
	}

	c.JSON(http.StatusCreated, skill)
}

// UpdateSkill godoc
// @Summary Update skill
// @Description Update an existing skill
// @Tags Portfolio - Skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Skill ID"
// @Param skill body models.Skill true "Skill data"
// @Success 200 {object} models.Skill
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skills/{id} [put]
func (h *Handler) UpdateSkill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var skill models.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill.ID = id
	if err := h.repo.UpdateSkill(&skill); err != nil {
		handleRepositoryError(c, err, "skill not found", "failed to update skill")
		return
	}

	c.JSON(http.StatusOK, skill)
}

// DeleteSkill godoc
// @Summary Delete skill
// @Description Delete a skill
// @Tags Portfolio - Skills
// @Security BearerAuth
// @Param id path int true "Skill ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skills/{id} [delete]
func (h *Handler) DeleteSkill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteSkill(id); err != nil {
		handleRepositoryError(c, err, "skill not found", "failed to delete skill")
		return
	}

	c.Status(http.StatusNoContent)
}

// SKILL TYPES

// GetAllSkillTypes godoc
// @Summary Get all skill types
// @Description Get all skill type categories
// @Tags Portfolio - Skills
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.SkillType
// @Failure 401 {object} map[string]string
// @Router /portfolio/skill-types [get]
func (h *Handler) GetAllSkillTypes(c *gin.Context) {
	skillTypes, err := h.repo.GetAllSkillTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch skill types"})
		return
	}

	c.JSON(http.StatusOK, skillTypes)
}

// GetSkillTypeByID godoc
// @Summary Get skill type by ID
// @Description Get a single skill type by ID
// @Tags Portfolio - Skills
// @Produce json
// @Security BearerAuth
// @Param id path int true "Skill Type ID"
// @Success 200 {object} models.SkillType
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skill-types/{id} [get]
func (h *Handler) GetSkillTypeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	skillType, err := h.repo.GetSkillTypeByID(id)
	if err != nil {
		handleRepositoryError(c, err, "skill type not found", "failed to fetch skill type")
		return
	}

	c.JSON(http.StatusOK, skillType)
}

// CreateSkillType godoc
// @Summary Create skill type
// @Description Create a new skill type category
// @Tags Portfolio - Skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param skillType body models.SkillType true "Skill type data"
// @Success 201 {object} models.SkillType
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skill-types [post]
func (h *Handler) CreateSkillType(c *gin.Context) {
	var skillType models.SkillType
	if err := c.ShouldBindJSON(&skillType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateSkillType(&skillType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create skill type"})
		return
	}

	c.JSON(http.StatusCreated, skillType)
}

// UpdateSkillType godoc
// @Summary Update skill type
// @Description Update an existing skill type
// @Tags Portfolio - Skills
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Skill Type ID"
// @Param skillType body models.SkillType true "Skill type data"
// @Success 200 {object} models.SkillType
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skill-types/{id} [put]
func (h *Handler) UpdateSkillType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var skillType models.SkillType
	if err := c.ShouldBindJSON(&skillType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skillType.ID = id
	if err := h.repo.UpdateSkillType(&skillType); err != nil {
		handleRepositoryError(c, err, "skill type not found", "failed to update skill type")
		return
	}

	c.JSON(http.StatusOK, skillType)
}

// DeleteSkillType godoc
// @Summary Delete skill type
// @Description Delete a skill type
// @Tags Portfolio - Skills
// @Security BearerAuth
// @Param id path int true "Skill Type ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /portfolio/skill-types/{id} [delete]
func (h *Handler) DeleteSkillType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteSkillType(id); err != nil {
		handleRepositoryError(c, err, "skill type not found", "failed to delete skill type")
		return
	}

	c.Status(http.StatusNoContent)
}
