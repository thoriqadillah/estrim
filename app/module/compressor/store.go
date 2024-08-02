package compressor

import (
	"estrim/db"
	"estrim/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store interface {
	Create(data CreateFile) (model.File, error)
	AssignJob(fileId string, jobId int64) error
}

type gormStore struct {
	db *gorm.DB
}

func NewStore() Store {
	return &gormStore{
		db: db.DB(),
	}
}

func (s *gormStore) Create(data CreateFile) (model.File, error) {
	file := model.File{
		Id:     uuid.NewString(),
		UserId: data.UserId,
		Name:   data.Name,
		Path:   data.Path,
		Mime:   data.Mime,
		Type:   data.Type,
		Size:   data.Size,
	}

	res := s.db.Create(&file)
	return file, res.Error
}

func (s *gormStore) AssignJob(fileId string, jobId int64) error {
	return s.db.Model(&model.File{}).
		Where("id = ?", fileId).
		Updates(
			map[string]interface{}{
				"job_id": jobId,
			}).Error
}
