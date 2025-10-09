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

type Certification struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Issuer        string    `json:"issuer"`
	IssueDate     string    `json:"issue_date"`
	ExpiryDate    *string   `json:"expiry_date"`
	CredentialID  string    `json:"credential_id"`
	CredentialURL string    `json:"credential_url"`
	DisplayOrder  int       `json:"display_order"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

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

func (Profile) TableName() string {
	return "profile"
}

func (WorkExperience) TableName() string {
	return "work_experience"
}

func (Certification) TableName() string {
	return "certifications"
}

func (MiniatureProject) TableName() string {
	return "miniature_projects"
}

func (Image) TableName() string {
	return "images"
}
