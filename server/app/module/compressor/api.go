package compressor

import (
	"fcompressor/app"
	"fcompressor/common/response"
	"fcompressor/lib/auth"
	"fcompressor/lib/storage"

	"github.com/gofiber/fiber/v2"
)

type compressorService struct {
	app     *fiber.App
	storage storage.Storage
	store   Store
}

func newService(app *fiber.App) app.Service {
	return &compressorService{
		app:     app,
		storage: storage.New(),
		store:   NewStore(),
	}
}

func (s *compressorService) compress(ctx *fiber.Ctx) error {
	user := ctx.Locals("user_id").(string)

	file, err := ctx.FormFile("file")
	if err != nil {
		return response.BadRequest(ctx, err)
	}

	reader, err := file.Open()
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	path, err := s.storage.Save(file.Filename, reader)
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	res, err := s.store.Create(CreateFile{
		UserId: user,
		Path:   path,
		Name:   file.Filename,
		Type:   ctx.FormValue("type"),
		Mime:   "",
		Size:   file.Size,
	})

	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	// TODO: compress in the background

	return response.Ok(ctx, res)
}

func (s *compressorService) CreateRoutes() {
	r := s.app.Group("/api/v1/compress")

	r.Post("/", auth.User, s.compress)
}

func init() {
	app.RegisterService(newService)
}
