package api

import (
	"fcompressor/app"
	"fcompressor/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

type compressorService struct {
	app     *fiber.App
	session *session.Store
}

func newService(app *fiber.App) app.Service {
	return &compressorService{
		app: app,
		session: session.New(session.Config{
			// Storage: store.NewSessionStorage(),
			Expiration: env.Get("SESSION_EXP").Duration("720h"),
		}),
	}
}

func (s *compressorService) initSession(ctx *fiber.Ctx) error {
	session, _ := s.session.Get(ctx)
	userId := session.Get("user_id")

	if userId == nil {
		id := uuid.NewString()
		session.Set("user_id", id)
		session.Save()

		userId = id
	}

	return ctx.SendString(userId.(string))
}

func (s *compressorService) CreateRoutes() {
	r := s.app.Group("/api/v1")

	r.Get("/user", s.initSession)
}

func init() {
	app.RegisterService(newService)
}
