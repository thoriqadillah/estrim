package storage

import (
	"estrim/common/env"
	"io"
)

type StorageObject struct {
	Name   string
	Mime   string
	Stream io.ReadCloser
}

type Storage interface {
	// Serve will serve the saved file
	// Pathlike is typically a URL to the file
	Serve(pathlike string) (io.ReadCloser, error)

	// Save will save a file using the filename and the reader
	Save(filename string, src io.Reader) (path string, err error)

	// Remove will remove a file from storage
	Remove(pathlike string) error
}

type Driver = string

var factories = map[Driver]Storage{}

func register(driver Driver, storage Storage) {
	factories[driver] = storage
}

func New(driver ...Driver) Storage {
	d := env.Get("STORAGE_DRIVER").String(Local)

	if len(driver) > 0 {
		d = driver[0]
	}

	return factories[d]
}
