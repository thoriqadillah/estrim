package service

import (
	"bytes"
	"estrim/db/model"
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

func (c *imageCompressor) Compress(target *model.File) (*model.File, error) {
	name := filepath.Base(target.Path)
	file, err := c.storage.Serve(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	converted, err := bimg.NewImage(b).Convert(bimg.WEBP)
	if err != nil {
		return nil, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{
		Quality: c.quality,
	})
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(processed)

	ext := filepath.Ext(name)
	newName := strings.Replace(name, ext, ".webp", 1)

	path, err := c.storage.Save(newName, r)
	if err != nil {
		return nil, err
	}

	if err := c.storage.Remove(name); err != nil {
		return nil, err
	}

	target.Path = path
	target.Mime = "image/webp"
	target.Size = int64(len(processed))
	target.IsCompressed = true

	return target, nil
}

func init() {
	registerFactory(model.Image, newImageCompressor)
}
