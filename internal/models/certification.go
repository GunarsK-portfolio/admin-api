package models

import "time"

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

func (Certification) TableName() string {
	return "certifications"
}
