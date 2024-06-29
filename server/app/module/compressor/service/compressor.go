package service

import (
	"estrim/common/env"
	"estrim/db/model"
	"estrim/lib/storage"
)

type Compressor interface {
	Compress(target *model.File) (*model.File, error)
}

type option struct {
	quality int
	storage storage.Storage
}

type Option = func(opt *option)

func UseStorage(driver storage.Driver) Option {
	return func(opt *option) {
		opt.storage = storage.New(driver)
	}
}

func SetQuality(quality int) Option {
	return func(opt *option) {
		opt.quality = quality
	}
}

type CompressorFactory = func(opt *option) Compressor

var factories = map[string]CompressorFactory{}

func NewCompressor(name model.Type, options ...Option) Compressor {
	option := &option{
		quality: 50,
		storage: storage.New(env.Get("STORAGE_DRIVER").String(storage.Local)),
	}

	return factories[name](option)
}

func registerFactory(name model.Type, f CompressorFactory) {
	factories[name] = f
}
