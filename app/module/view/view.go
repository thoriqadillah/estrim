package view

import (
	"embed"
	"estrim/app"
	"estrim/common/response"
	"estrim/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type viewService struct {
	*app.App
	web embed.FS
}

func NewService(app *app.App) app.Service {
	return &viewService{
		App: app,
		web: web.Web,
	}
}

func (s *viewService) serve404(ctx *fiber.Ctx) error {
	html, err := s.web.ReadFile("dist/index.html")
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return ctx.Type("html").Send(html)
}

func (s *viewService) CreateRoutes() {
	s.Api.Use("/", filesystem.New(filesystem.Config{
		Root: web.ServeWeb(),
	}))

	s.Api.Use("*", s.serve404)
}
