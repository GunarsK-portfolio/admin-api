package models

import "time"

type MiniatureTheme struct {
	ID           int64              `json:"id" gorm:"primaryKey"`
	Name         string             `json:"name" binding:"required"`
	Description  string             `json:"description"`
	CoverImageID *int64             `json:"coverImageId,omitempty" gorm:"column:cover_image_id"`
	DisplayOrder int                `json:"displayOrder,omitempty" gorm:"column:display_order;default:0"`
	Miniatures   []MiniatureProject `json:"miniatures,omitempty" gorm:"foreignKey:ThemeID"`
	CreatedAt    time.Time          `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time          `json:"updatedAt" gorm:"column:updated_at"`
}

func (MiniatureTheme) TableName() string {
	return "miniatures.miniature_themes"
}

type MiniatureProject struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	ThemeID       *int64    `json:"themeId,omitempty" gorm:"column:theme_id"`
	Title         string    `json:"name" gorm:"column:title" binding:"required"` // Frontend expects "name"
	Description   string    `json:"description"`
	CompletedDate *string   `json:"completedDate,omitempty" gorm:"column:completed_date;type:date"`
	Scale         string    `json:"scale,omitempty"`
	Manufacturer  string    `json:"manufacturer,omitempty"`
	TimeSpent     *float64  `json:"timeSpent,omitempty" gorm:"column:time_spent;type:numeric(6,2)"` // Hours as decimal
	Difficulty    string    `json:"difficulty,omitempty"`
	DisplayOrder  int       `json:"displayOrder,omitempty" gorm:"column:display_order;default:0"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`

	// Associations (loaded with Preload)
	Theme          *MiniatureTheme `json:"theme,omitempty" gorm:"foreignKey:ThemeID"`
	MiniatureFiles []MiniatureFile `json:"-" gorm:"foreignKey:MiniatureProjectID"` // Database relation
	Images         []Image         `json:"images,omitempty" gorm:"-"`              // Computed for frontend from MiniatureFiles
	Techniques     []string        `json:"techniques,omitempty" gorm:"-"`
}

func (MiniatureProject) TableName() string {
	return "miniatures.miniature_projects"
}
