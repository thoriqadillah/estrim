package main

import (
	"embed"
	"estrim/app"
	_ "estrim/app/module"
	"estrim/common/response"
	"estrim/db"
	"estrim/lib/auth"
	"io/fs"
	"log"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	godotenv.Load()
}

//go:embed web/dist/*
var web embed.FS

func frontend() http.FileSystem {
	dist, err := fs.Sub(web, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(dist)
}

func notFound(ctx *fiber.Ctx) error {
	html, err := web.ReadFile("web/dist/index.html")
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return ctx.Type("html").Send(html)
}

func main() {
	db.Open()

	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiber.Use("/", filesystem.New(filesystem.Config{
		Root: frontend(),
	}))

	fiber.Use(
		logger.New(),
		auth.Session,
	)

	api := app.New(fiber)
	fiber.Use("*", notFound)

	api.Start()
}
