package models

import "time"

// Image is the simplified view for frontend (computed from MiniatureFile + StorageFile)
type Image struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

// MiniatureFile is the junction table between miniature projects and storage files
type MiniatureFile struct {
	ID                 int64        `json:"id" gorm:"primaryKey"`
	MiniatureProjectID int64        `json:"miniatureProjectId" gorm:"column:miniature_project_id"`
	FileID             int64        `json:"fileId" gorm:"column:file_id"`
	Caption            string       `json:"caption"`
	DisplayOrder       int          `json:"displayOrder,omitempty" gorm:"column:display_order;default:0"`
	File               *StorageFile `json:"file,omitempty" gorm:"foreignKey:FileID"`
	CreatedAt          time.Time    `json:"createdAt" gorm:"column:created_at"`
}

func (MiniatureFile) TableName() string {
	return "miniatures.miniature_files"
}

// StorageFile represents files stored in S3/MinIO
type StorageFile struct {
	ID       int64     `json:"id" gorm:"primaryKey"`
	S3Key    string    `json:"s3Key" gorm:"column:s3_key"`
	S3Bucket string    `json:"s3Bucket" gorm:"column:s3_bucket"`
	FileName string    `json:"fileName" gorm:"column:file_name"`
	FileSize int64     `json:"fileSize" gorm:"column:file_size"`
	MimeType string    `json:"mimeType" gorm:"column:mime_type"`
	FileType string    `json:"fileType" gorm:"column:file_type"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (StorageFile) TableName() string {
	return "storage.files"
}
