package compressor

import (
	"fcompressor/model"
	"fcompressor/store"
)

type imageCompressor struct {
	id    string
	store store.Store
}

func newImageCompressor(id string, option *option) Compressor {
	return &imageCompressor{
		id:    id,
		store: option.store,
	}
}

func (c *imageCompressor) Compress() {

}

func init() {
	registerFactory(model.Image, newImageCompressor)
}
