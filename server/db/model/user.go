package model

import "time"

type User struct {
	Id        uint `gorm:"primaryKey;index;autoIncrement"`
	Key       string
	Value     []byte    `gorm:"type:bytea"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
