package compressor

import (
	"fcompressor/app/store"
	"fcompressor/env"
	"fcompressor/lib/storage"
	"fcompressor/model"
)

type Compressor interface {
	Compress(target *model.File) (*model.File, error)
}

type option struct {
	storage storage.Storage
}

type Option = func(opt *option)

func UseStorage(driver storage.Driver) Option {
	return func(opt *option) {
		opt.storage = storage.New(driver)
	}
}

type CompressorFactory = func(opt *option) Compressor

var factories = map[string]CompressorFactory{}

func New(name model.Type, store store.Store, options ...Option) Compressor {
	option := &option{
		storage: storage.New(env.Get("STORAGE_DRIVER").String(storage.Local)),
	}

	return factories[name](option)
}

func registerFactory(name model.Type, f CompressorFactory) {
	factories[name] = f
}
