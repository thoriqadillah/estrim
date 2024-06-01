package store

import "fcompressor/model"

type Store interface {
	Get(id string) (model.File, error)
	Save(data model.CreateFile) (id model.File, err error)
	Update(id string, data model.UpdateFile) (model.File, error)
}
