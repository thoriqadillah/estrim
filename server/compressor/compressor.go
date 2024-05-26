package compressor

import (
	"fcompressor/model"
	"fcompressor/store"
)

type Compressor interface {
	Compress()
}

type option struct {
	store store.Store
}

type Option = func(opt *option)

func UseStore(store store.Store) Option {
	return func(opt *option) {
		opt.store = store
	}
}

type CompressorFactory = func(id string, opt *option) Compressor

var factories = map[string]CompressorFactory{}

func New(name model.Type, id string, options ...Option) Compressor {
	option := &option{}

	return factories[name](id, option)
}

func registerFactory(name model.Type, f CompressorFactory) {
	factories[name] = f
}
