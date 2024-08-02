package queue

import (
	"context"
	"estrim/db"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

func New(workers *river.Workers) *river.Client[pgx.Tx] {
	ctx := context.Background()

	dsn := db.Connection()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Panicln("Failed to connect database pool:", err)
	}

	queue, err := river.NewClient(riverpgxv5.New(pool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {
				MaxWorkers: 100,
			},
		},
		Workers: workers,
	})

	if err != nil {
		log.Panicln("Failed to connect queue:", err)
	}

	return queue
}
