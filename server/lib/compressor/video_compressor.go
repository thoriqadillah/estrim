package compressor

import (
	"fcompressor/api/model"
)

type videoCompressor struct {
	*option
}

func newVideoCompressor(option *option) Compressor {
	return &videoCompressor{
		option: option,
	}
}

func (c *videoCompressor) Compress(fileId string) {

}

func init() {
	registerFactory(model.Video, newVideoCompressor)
}
