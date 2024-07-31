package service

import (
	"estrim/db/model"
)

type videoCompressor struct {
	*option
}

func newVideoCompressor(option *option) Compressor {
	return &videoCompressor{
		option: option,
	}
}

func (c *videoCompressor) Compress(target *model.File) error {
	return nil
}

func init() {
	registerFactory(model.Video, newVideoCompressor)
}
