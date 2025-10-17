package models

import "time"

type WorkExperience struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Company      string    `json:"company"`
	Position     string    `json:"position"`
	Description  string    `json:"description"`
	StartDate    string    `json:"start_date"`
	EndDate      *string   `json:"end_date"`
	IsCurrent    bool      `json:"is_current"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (WorkExperience) TableName() string {
	return "work_experience"
}
