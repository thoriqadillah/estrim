package model

import "time"

type Session struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
