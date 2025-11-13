package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
)

// Skills

func (r *repository) GetAllSkills(ctx context.Context) ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.WithContext(ctx).Preload("SkillType").Order("display_order ASC, skill ASC").Find(&skills).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all skills: %w", err)
	}
	return skills, nil
}

func (r *repository) GetSkillByID(ctx context.Context, id int64) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.WithContext(ctx).Preload("SkillType").First(&skill, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get skill with id %d: %w", id, err)
	}
	return &skill, nil
}

func (r *repository) CreateSkill(ctx context.Context, skill *models.Skill) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(skill).Error
	if err != nil {
		return fmt.Errorf("failed to create skill: %w", err)
	}
	return nil
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
	if err != nil {
		return nil, fmt.Errorf("failed to get all skill types: %w", err)
	}
	return skillTypes, nil
}

func (r *repository) GetSkillTypeByID(ctx context.Context, id int64) (*models.SkillType, error) {
	var skillType models.SkillType
	err := r.db.WithContext(ctx).First(&skillType, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get skill type with id %d: %w", id, err)
	}
	return &skillType, nil
}

func (r *repository) CreateSkillType(ctx context.Context, skillType *models.SkillType) error {
	err := r.db.WithContext(ctx).Omit("ID", "CreatedAt", "UpdatedAt").Create(skillType).Error
	if err != nil {
		return fmt.Errorf("failed to create skill type: %w", err)
	}
	return nil
}

func (r *repository) UpdateSkillType(ctx context.Context, skillType *models.SkillType) error {
	return r.safeUpdate(ctx, skillType, skillType.ID)
}

func (r *repository) DeleteSkillType(ctx context.Context, id int64) error {
	return checkRowsAffected(r.db.WithContext(ctx).Delete(&models.SkillType{}, id))
}
