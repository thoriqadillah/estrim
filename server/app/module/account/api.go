package account

import (
	"fcompressor/app"
	"fcompressor/common/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type accountService struct {
	app   *fiber.App
	store Store
}

func newService(app *fiber.App) app.Service {
	return &accountService{
		app:   app,
		store: NewStore(),
	}
}

func (s *accountService) initSession(ctx *fiber.Ctx) error {
	session := ctx.Locals("session").(*session.Session)

	res, err := s.store.Session(session.ID())
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *accountService) CreateRoutes() {
	r := s.app.Group("/api/v1/account")

	r.Get("/session", s.initSession)
}

func init() {
	app.RegisterService(newService)
}
