package model

import "time"

type Type = string

const (
	Image Type = "image"
	Video Type = "video"
	PDF   Type = "pdf"
)

type File struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey;index"`
	UserId       string    `json:"user_id"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Mime         string    `json:"mime"`
	Type         Type      `json:"type"`
	Size         int64     `json:"size"`
	IsCompressed bool      `json:"is_compressed" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`

	Owner User `json:"-" gorm:"foreignKey:UserId"`
}
