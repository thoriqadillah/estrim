package storage

import (
	"io"
	"os"
	"path/filepath"
)

const Local Driver = "local"

type localStorage struct {
	dir string
}

func newLocalStorage() Storage {
	wd, _ := os.Getwd()
	storageDir := filepath.Join(wd, "lib", "storage", "uploads")
	os.MkdirAll(storageDir, os.ModePerm)

	return &localStorage{
		dir: storageDir,
	}
}

func (s *localStorage) Serve(filename string) (io.ReadCloser, error) {
	path := filepath.Join(s.dir, filename)

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *localStorage) Save(filename string, src io.Reader) (path string, err error) {
	path = filepath.Join(s.dir, filename)
	dst, err := os.Create(path)
	if err != nil {
		return "", nil
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}

func (s *localStorage) Remove(filename string) error {
	path := filepath.Join(s.dir, filename)
	return os.Remove(path)
}

func init() {
	register(Local, newLocalStorage())
}
