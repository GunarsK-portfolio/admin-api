package repository

import (
	"context"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

// Skills

func (r *repository) GetAllSkills(ctx context.Context) ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.WithContext(ctx).Preload("SkillType").Order("display_order ASC, skill ASC").Find(&skills).Error
	return skills, err
}

func (r *repository) GetSkillByID(ctx context.Context, id int64) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.WithContext(ctx).Preload("SkillType").First(&skill, id).Error
	return &skill, err
}

func (r *repository) CreateSkill(ctx context.Context, skill *models.Skill) error {
	return r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(skill).Error
}

func (r *repository) UpdateSkill(ctx context.Context, skill *models.Skill) error {
	return r.safeUpdate(ctx, skill, skill.ID)
}

func (r *repository) DeleteSkill(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.Skill{}, id))
}

// Skill Types

func (r *repository) GetAllSkillTypes(ctx context.Context) ([]models.SkillType, error) {
	var skillTypes []models.SkillType
	err := r.db.WithContext(ctx).Order("display_order ASC, name ASC").Find(&skillTypes).Error
	return skillTypes, err
}

func (r *repository) GetSkillTypeByID(ctx context.Context, id int64) (*models.SkillType, error) {
	var skillType models.SkillType
	err := r.db.WithContext(ctx).First(&skillType, id).Error
	return &skillType, err
}

func (r *repository) CreateSkillType(ctx context.Context, skillType *models.SkillType) error {
	return r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(skillType).Error
}

func (r *repository) UpdateSkillType(ctx context.Context, skillType *models.SkillType) error {
	return r.safeUpdate(ctx, skillType, skillType.ID)
}

func (r *repository) DeleteSkillType(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.SkillType{}, id))
}
