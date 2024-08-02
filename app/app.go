package app

import (
	"context"
	"estrim/common/env"
	"estrim/lib/queue"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
)

type Service interface {
	CreateRoutes()
}

type Worker interface {
	CreateWorker(*river.Workers)
}

type Initter interface {
	Init() error
}

type Closer interface {
	Close() error
}

type ServiceFactory = func(app *App) Service

type App struct {
	Api      *fiber.App
	services []Service
	workers  *river.Workers
	Queue    *river.Client[pgx.Tx]
}

var services = make([]ServiceFactory, 0)

func RegisterService(s ServiceFactory) {
	services = append(services, s)
}

func New(fiber *fiber.App) *App {
	_services := make([]Service, 0)
	workers := river.NewWorkers()

	app := &App{
		Api:     fiber,
		workers: workers,
	}

	for _, service := range services {
		svc := service(app)

		if worker, ok := svc.(Worker); ok {
			worker.CreateWorker(workers)
		}

		if init, ok := svc.(Initter); ok {
			if err := init.Init(); err != nil {
				log.Fatal(err)
			}
		}

		svc.CreateRoutes()
		_services = append(_services, svc)
	}

	app.services = _services
	app.Queue = queue.New(workers)

	return app
}

func (a *App) Start() {
	ctx := context.Background()

	if err := a.Queue.Start(ctx); err != nil {
		log.Panicln("Failed to start queue:", err)
	}

	// graceful shutdown
	signals := []os.Signal{syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, os.Interrupt}
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, signals...)

	hardCtx, hardCancel := context.WithCancel(ctx)
	go func() {
		<-ch

		hardCancel()
		log.Println("Shutting down...")

		defer func() {
			a.Api.Shutdown()
			a.Queue.StopAndCancel(hardCtx)

			log.Println("Queue stopped")
		}()

		for _, service := range a.services {
			if closer, ok := service.(Closer); ok {
				if err := closer.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	port := env.Get("API_PORT").String(":8888")

	if err := a.Api.Listen(port); err != nil {
		log.Fatal(err)
	}
}
