package store

import (
	"fcompressor/db"
	"fcompressor/db/model"

	"gorm.io/gorm"
)

type SessionStore interface {
	Init(id string) (model.Session, error)
}

type sessionStore struct {
	db *gorm.DB
}

func NewSessionStore() SessionStore {
	return &sessionStore{
		db: db.DB(),
	}
}

func (s *sessionStore) Init(id string) (model.Session, error) {
	session := model.Session{
		Id: id,
	}

	res := s.db.First(&session)
	if res.RowsAffected == 1 {
		return session, nil
	}

	res = s.db.Create(&session)
	return session, res.Error
}
