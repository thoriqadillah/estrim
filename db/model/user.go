package model

import "time"

type Source = string

const (
	Session Source = "session"
	Email   Source = "email"
	Google  Source = "google"
)

type User struct {
	Id                 string    `json:"id"`
	Name               *string   `json:"name,omitempty"`
	Email              *string   `json:"email,omitempty"`
	Password           *string   `json:"-"`
	Source             Source    `json:"source" gorm:"default:session"`
	ResetPasswordToken *string   `json:"-"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`

	Files []File `json:"files" gorm:"foreignKey:UserId"`
}
