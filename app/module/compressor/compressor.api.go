package compressor

import (
	"context"
	"estrim/app"
	"estrim/app/module/compressor/service"
	"estrim/common/response"
	"estrim/lib/auth"
	"estrim/lib/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/riverqueue/river"
)

type compressorService struct {
	*app.App
	storage storage.Storage
	store   Store
	river.WorkerDefaults[CompressFile]
}

func NewService(app *app.App) app.Service {
	return &compressorService{
		App:     app,
		storage: storage.New(),
		store:   NewStore(),
	}
}

func (s *compressorService) Work(ctx context.Context, job *river.Job[CompressFile]) error {
	file := job.Args.File
	option := job.Args.Option

	select {
	case <-ctx.Done():
		return context.Canceled
	default:
		compressor := service.NewCompressor(file.Type, option.toFunc()...)
		if err := compressor.Compress(&file); err != nil {
			return err
		}

		// TODO: update the file
		// TODO: notify user

		return nil
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

	args := createArgs(res, ctx)

	inserted, err := s.Queue.Insert(context.Background(), args, nil)
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	if err := s.store.AssignJob(res.Id, inserted.Job.ID); err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *compressorService) CreateRoutes() {
	r := s.Api.Group("/api/v1/compress")

	r.Post("/", auth.User, s.compress)
}

func (s *compressorService) CreateWorker(workers *river.Workers) {
	river.AddWorker(workers, s)
}
