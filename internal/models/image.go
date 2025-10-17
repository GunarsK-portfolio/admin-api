package models

import "time"

type Image struct {
	ID                 int64     `json:"id" gorm:"primaryKey"`
	MiniatureProjectID *int64    `json:"miniature_project_id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	S3Key              string    `json:"-"`
	S3Bucket           string    `json:"-"`
	URL                string    `json:"url"`
	DisplayOrder       int       `json:"display_order"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (Image) TableName() string {
	return "images"
}
