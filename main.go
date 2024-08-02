package main

import (
	"embed"
	"estrim/app"
	_ "estrim/app/module/account"
	_ "estrim/app/module/compressor"
	_ "estrim/app/module/storage"
	"estrim/db"
	"estrim/lib/auth"
	"io/fs"
	"log"
	"net/http"

	"github.com/goccy/go-json"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed web/dist/*
var web embed.FS

func frontend() http.FileSystem {
	dist, err := fs.Sub(web, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(dist)
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
	api.Start()
}
