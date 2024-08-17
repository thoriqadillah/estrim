package account

import (
	"estrim/app"
	"estrim/common/response"
	"estrim/lib/auth"

	"github.com/gofiber/fiber/v2"
)

type accountService struct {
	*app.App
	store Store
}

func NewService(app *app.App) app.Service {
	return &accountService{
		App:   app,
		store: NewStore(),
	}
}

func (s *accountService) initSession(ctx *fiber.Ctx) error {
	user := ctx.Locals("user_id").(string)

	res, err := s.store.Session(user)
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *accountService) refreshToken(ctx *fiber.Ctx) error {
	// TODO
	return response.Unauthorized(ctx)
}

func (s *accountService) CreateRoutes() {
	r := s.Api.Group("/api/v1/account")

	r.Get("/", auth.User, s.initSession)
	r.Get("/refresh-token", s.refreshToken)
}
