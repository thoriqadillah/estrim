package app

import (
	"estrim/common/env"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type App interface {
	Start()
}

type Service interface {
	CreateRoutes()
}

type Initter interface {
	Init() error
}

type Closer interface {
	Close() error
}

type ServiceFactory = func(app *fiber.App) Service

type api struct {
	*fiber.App
	services []Service
}

var services = make([]ServiceFactory, 0)

func RegisterService(s ServiceFactory) {
	services = append(services, s)
}

func New(app *fiber.App) App {
	_services := make([]Service, 0)
	for _, service := range services {
		_services = append(_services, service(app))
	}

	return &api{
		services: _services,
		App:      app,
	}
}

func (a *api) Start() {
	for _, service := range a.services {
		if init, ok := service.(Initter); ok {
			if err := init.Init(); err != nil {
				log.Fatal(err)
			}
		}

		service.CreateRoutes()
	}

	// graceful shutdown
	signals := []os.Signal{syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, os.Interrupt}
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, signals...)

	go func() {
		<-ch
		log.Println("Shutting down...")

		defer a.Shutdown()

		for _, service := range a.services {
			if closer, ok := service.(Closer); ok {
				if err := closer.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	port := env.Get("API_PORT").String(":8888")

	if err := a.Listen(port); err != nil {
		log.Fatal(err)
	}
}
