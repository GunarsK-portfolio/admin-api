package repository

import "github.com/GunarsK-portfolio/admin-api/internal/models"

// Skills

func (r *repository) GetAllSkills() ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.Preload("SkillType").Order("display_order ASC, skill ASC").Find(&skills).Error
	return skills, err
}

func (r *repository) GetSkillByID(id int64) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.Preload("SkillType").First(&skill, id).Error
	return &skill, err
}

func (r *repository) CreateSkill(skill *models.Skill) error {
	return r.db.Omit("ID", "CreatedAt", "UpdatedAt").Create(skill).Error
}

func (r *repository) UpdateSkill(skill *models.Skill) error {
	return r.safeUpdate(skill, skill.ID)
}

func (r *repository) DeleteSkill(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.Skill{}, id))
}

// Skill Types

func (r *repository) GetAllSkillTypes() ([]models.SkillType, error) {
	var skillTypes []models.SkillType
	err := r.db.Order("display_order ASC, name ASC").Find(&skillTypes).Error
	return skillTypes, err
}

func (r *repository) GetSkillTypeByID(id int64) (*models.SkillType, error) {
	var skillType models.SkillType
	err := r.db.First(&skillType, id).Error
	return &skillType, err
}

func (r *repository) CreateSkillType(skillType *models.SkillType) error {
	return r.db.Omit("ID", "CreatedAt", "UpdatedAt").Create(skillType).Error
}

func (r *repository) UpdateSkillType(skillType *models.SkillType) error {
	return r.safeUpdate(skillType, skillType.ID)
}

func (r *repository) DeleteSkillType(id int64) error {
	return checkRowsAffected(r.db.Delete(&models.SkillType{}, id))
}
