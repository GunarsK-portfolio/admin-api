package models

import "time"

type Profile struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name"`
	Title     string    `json:"title"`
	Bio       string    `json:"bio"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Profile) TableName() string {
	return "profile"
}
