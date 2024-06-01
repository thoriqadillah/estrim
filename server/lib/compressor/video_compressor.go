package compressor

import (
	"fcompressor/model"
)

type videoCompressor struct {
	*option
}

func newVideoCompressor(option *option) Compressor {
	return &videoCompressor{
		option: option,
	}
}

func (c *videoCompressor) Compress(target *model.File) (*model.File, error) {
	return nil, nil
}

func init() {
	registerFactory(model.Video, newVideoCompressor)
}
