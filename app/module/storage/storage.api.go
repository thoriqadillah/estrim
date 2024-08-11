package storage

import (
	"estrim/app"
	"estrim/common/response"
	"estrim/lib/storage"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type storageService struct {
	*app.App
	storage storage.Storage
}

func NewService(app *app.App) app.Service {
	return &storageService{
		App:     app,
		storage: storage.New(),
	}
}

func (s *storageService) serveFile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	file, err := s.storage.Serve(id)
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	ext := filepath.Ext(id)
	mimetype := mime.TypeByExtension(ext)

	ctx.Set("Content-Type", mimetype)
	return ctx.SendStream(file)
}

func (s *storageService) CreateRoutes() {
	r := s.Api.Group("/api/v1/storage")

	r.Get("/:id", s.serveFile)
}

func init() {
	app.RegisterService(NewService)
}
