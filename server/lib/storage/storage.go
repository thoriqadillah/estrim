package storage

import (
	"fcompressor/env"
	"io"
	"log"
)

type Storage interface {
	Serve(filename string) (io.Reader, error)
	Save(filename string, src io.Reader) (path string, err error)
	Remove(filename string) error
}

type Driver = string

var factories = map[Driver]Storage{}

func register(driver Driver, storage Storage) {
	factories[driver] = storage
}

func New(driver ...Driver) Storage {
	d := env.Get("STORAGE_DRIVER").String(Local)
	log.Println(factories[d])

	if len(driver) > 0 {
		d = driver[0]
	}

	return factories[d]
}
