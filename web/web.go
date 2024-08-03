package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/*
var Web embed.FS

func ServeWeb() http.FileSystem {
	dist, err := fs.Sub(Web, "dist")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(dist)
}
