package models

import "time"

type Profile struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	FullName     string    `json:"name" gorm:"column:full_name" binding:"required"`
	Title        string    `json:"title"`
	Bio          string    `json:"tagline"` // Short bio/tagline
	Email        string    `json:"email"`
	Phone        string    `json:"phone,omitempty"`
	Location     string    `json:"location,omitempty"`
	AvatarFileID *int64    `json:"avatarFileId,omitempty" gorm:"column:avatar_file_id"`
	ResumeFileID *int64    `json:"resumeFileId,omitempty" gorm:"column:resume_file_id"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Profile) TableName() string {
	return "portfolio.profile"
}
