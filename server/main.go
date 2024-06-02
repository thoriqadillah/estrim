package main

import (
	"fcompressor/app"
	_ "fcompressor/app/module/account"
	"fcompressor/db"
	"fcompressor/env"
	_ "fcompressor/env"
	"fcompressor/lib/auth"

	"github.com/goccy/go-json"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.Open()

	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiber.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowCredentials: env.Get("CORS_ORIGINS").String("*") != "*",
			AllowOrigins:     env.Get("CORS_ORIGINS").String("*"),
		}),
		auth.Session,
	)

	api := app.New(fiber)
	api.Start()
}
