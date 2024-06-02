package account

import (
	"fcompressor/db"
	"fcompressor/db/model"

	"gorm.io/gorm"
)

type Store interface {
	Session(id string) (model.User, error)
}

type gormStore struct {
	db *gorm.DB
}

func NewStore() Store {
	return &gormStore{
		db: db.DB(),
	}
}

func (s *gormStore) Session(id string) (model.User, error) {
	session := model.User{
		Id: id,
	}

	res := s.db.First(&session)
	if res.RowsAffected == 1 {
		return session, nil
	}

	res = s.db.Create(&session)
	return session, res.Error
}
