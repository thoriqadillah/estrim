package compressor

import (
	"bytes"
	"fcompressor/api/model"
	"io"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

type imageCompressor struct {
	*option
}

func newImageCompressor(option *option) Compressor {
	return &imageCompressor{
		option: option,
	}
}

func (c *imageCompressor) Compress(fileId string) {
	target, err := c.store.Get(fileId)
	if err != nil {
		// TODO
	}

	name := filepath.Base(target.Path)
	file, err := c.storage.Serve(name)
	if err != nil {
		// TODO
	}

	b, err := io.ReadAll(file)
	if err != nil {
		// TODO
	}

	converted, err := bimg.NewImage(b).Convert(bimg.WEBP)
	if err != nil {
		// TODO
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{
		Quality: 90,
	})
	if err != nil {
		// TODO
	}

	r := bytes.NewReader(processed)

	ext := filepath.Ext(name)
	newName := strings.Replace(name, ext, ".webp", 1)

	path, err := c.storage.Save(newName, r)
	if err != nil {
		// TODO
	}

	if err := c.storage.Remove(name); err != nil {
		// TODO
	}

	compressed := true
	_, err = c.store.Update(fileId, model.UpdateFile{
		Path:       &path,
		Compressed: &compressed,
	})

	if err != nil {
		// TODO
	}

}

func init() {
	registerFactory(model.Image, newImageCompressor)
}
