package compressor

import (
	"fcompressor/api/model"
	"fcompressor/api/store"
	"fcompressor/env"
	"fcompressor/lib/storage"
)

type Compressor interface {
	Compress(fileId string)
}

type option struct {
	store   store.Store
	storage storage.Storage
}

type Option = func(opt *option)

func UseStore(store store.Store) Option {
	return func(opt *option) {
		opt.store = store
	}
}

func UseStorage(driver storage.Driver) Option {
	return func(opt *option) {
		opt.storage = storage.New(driver)
	}
}

type CompressorFactory = func(opt *option) Compressor

var factories = map[string]CompressorFactory{}

func New(name model.Type, options ...Option) Compressor {
	option := &option{
		storage: storage.New(env.Get("STORAGE_DRIVER").String(storage.Local)),
	}

	return factories[name](option)
}

func registerFactory(name model.Type, f CompressorFactory) {
	factories[name] = f
}
