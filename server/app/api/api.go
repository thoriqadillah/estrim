package api

import (
	"fcompressor/app"
	"fcompressor/app/store"
	"fcompressor/common/response"
	"fcompressor/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type compressorService struct {
	app          *fiber.App
	session      *session.Store
	sessionStore store.SessionStore
}

func newService(app *fiber.App) app.Service {
	return &compressorService{
		app:          app,
		sessionStore: store.NewSessionStore(),
		session: session.New(session.Config{
			Expiration: env.Get("SESSION_EXP").Duration("720h"),
		}),
	}
}

func (s *compressorService) initSession(ctx *fiber.Ctx) error {
	session, _ := s.session.Get(ctx)
	res, err := s.sessionStore.Init(session.ID())
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *compressorService) CreateRoutes() {
	r := s.app.Group("/api/v1")

	r.Get("/user", s.initSession)
}

func init() {
	app.RegisterService(newService)
}
