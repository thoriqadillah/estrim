package main

import (
	"estrim/app"
	_ "estrim/app/module"
	"estrim/db"
	"estrim/lib/auth"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	godotenv.Load()
}

func main() {
	db.Open()

	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiber.Use(
		logger.New(),
		auth.Session,
	)

	api := app.New(fiber)
	api.Start()
}
