package models

import "time"

type MiniatureProject struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CompletedDate *string   `json:"completed_date"`
	DisplayOrder  int       `json:"display_order"`
	Images        []Image   `json:"images,omitempty" gorm:"foreignKey:MiniatureProjectID"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (MiniatureProject) TableName() string {
	return "miniature_projects"
}
