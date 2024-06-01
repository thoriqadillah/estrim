package api

import (
	"fcompressor/app"

	"github.com/gofiber/fiber/v2"
)

type compressorService struct {
	app *fiber.App
}

func newService(app *fiber.App) app.Service {
	return &compressorService{
		app: app,
	}
}

func (s *compressorService) CreateRoutes() {

}

func init() {
	app.RegisterService(newService)
}
