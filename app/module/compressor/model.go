package compressor

import (
	"estrim/app/module/compressor/service"
	"estrim/common"
	"estrim/db/model"

	"github.com/gofiber/fiber/v2"
	"github.com/riverqueue/river"
)

type Type = string

const (
	Image Type = "image"
	Video Type = "video"
	PDF   Type = "pdf"
)

type CreateFile struct {
	UserId string
	Path   string
	Name   string
	Type   Type
	Mime   string
	Size   int64
}

type UpdateFile struct {
	Path       *string `json:"path"`
	Compressed *bool   `json:"compressed"`
	Size       *int64  `json:"size"`
}

type CompressOption struct {
	quality int
}

func (c *CompressOption) toFunc() []service.Option {
	opts := make([]service.Option, 0)
	opts = append(opts,
		service.SetQuality(c.quality),
	)

	return opts
}

type CompressFile struct {
	File   model.File     `json:"file"`
	Option CompressOption `json:"option"`
}

func (CompressFile) Kind() string {
	return "file"
}

func (c *CompressFile) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		MaxAttempts: 3,
	}
}

func createArgs(file model.File, ctx *fiber.Ctx) CompressFile {
	return CompressFile{
		File: file,
		Option: CompressOption{
			quality: common.ParseString(ctx.FormValue("quality")).Int(50),
		},
	}
}
