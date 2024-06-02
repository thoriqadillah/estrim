package account

import (
	"fcompressor/db"

	"gorm.io/gorm"
)

type Store interface {
	Init(id string) (User, error)
}

type sessionStore struct {
	db *gorm.DB
}

func NewStore() Store {
	return &sessionStore{
		db: db.DB(),
	}
}

func (s *sessionStore) Init(id string) (User, error) {
	session := User{
		Id: id,
	}

	res := s.db.First(&session)
	if res.RowsAffected == 1 {
		return session, nil
	}

	res = s.db.Create(&session)
	return session, res.Error
}
