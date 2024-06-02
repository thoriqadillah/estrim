package storage

import (
	"fcompressor/env"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const Local Driver = "local"

type localStorage struct {
	dir     string
	baseUrl string
}

func newLocalStorage() Storage {
	wd, _ := os.Getwd()
	storageDir := filepath.Join(wd, "lib", "storage", "uploads")
	os.MkdirAll(storageDir, os.ModePerm)

	return &localStorage{
		dir:     storageDir,
		baseUrl: env.Get("BASE_URL").String(),
	}
}

func (s *localStorage) Serve(pathlike string) (io.ReadCloser, error) {
	filename := filepath.Base(pathlike)
	path := filepath.Join(s.dir, filename)

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *localStorage) Save(filename string, src io.Reader) (path string, err error) {
	id := uuid.NewString() + filepath.Ext(filename)
	path = filepath.Join(s.dir, id)

	dst, err := os.Create(path)
	if err != nil {
		return "", nil
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/storage/%s", s.baseUrl, id), nil
}

func (s *localStorage) Remove(pathlike string) error {
	filename := filepath.Base(pathlike)
	path := filepath.Join(s.dir, filename)

	return os.Remove(path)
}

func init() {
	register(Local, newLocalStorage())
}
