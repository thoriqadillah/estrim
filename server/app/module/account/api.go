package account

import (
	"fcompressor/app"
	"fcompressor/common/response"
	"fcompressor/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type accountService struct {
	app          *fiber.App
	session      *session.Store
	sessionStore Store
}

func newService(app *fiber.App) app.Service {
	return &accountService{
		app:          app,
		sessionStore: NewStore(),
		session: session.New(session.Config{
			Expiration: env.Get("SESSION_EXP").Duration("720h"),
		}),
	}
}

func (s *accountService) initSession(ctx *fiber.Ctx) error {
	session, _ := s.session.Get(ctx)
	res, err := s.sessionStore.Init(session.ID())
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *accountService) CreateRoutes() {
	r := s.app.Group("/api/v1/user")

	r.Get("/", s.initSession)
}

func init() {
	app.RegisterService(newService)
}
