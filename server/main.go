package main

import (
	"fcompressor/app"
	_ "fcompressor/env"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiber.Use(logger.New())

	api := app.New(fiber)
	api.Start()
}
