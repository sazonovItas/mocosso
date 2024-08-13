package storageapp

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/sazonovItas/auth-service/db/postgres"
	"github.com/sazonovItas/auth-service/internal/app"
	"github.com/sazonovItas/auth-service/internal/config"
	"github.com/sazonovItas/auth-service/pkg/logger"
	"github.com/sazonovItas/auth-service/pkg/postgres"
)

const connectTimeout = 3 * time.Second

type App struct {
	*db.Queries

	connPool *pgxpool.Pool
}

func New(cfg config.SectionStorage) (storage *App) {
	const op = "app.storgeapp.New"

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	pool, err := postgres.Connect(ctx, cfg.URI)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	storage = &App{
		Queries:  db.New(pool),
		connPool: pool,
	}

	return storage
}

// Run method is implementation of Service interface.
func (a *App) Run(ctx context.Context) error {
	return nil
}

// Shutdown method is implementation of Service interface.
func (a *App) Shutdown(ctx context.Context) error {
	a.connPool.Close()
	return nil
}

// SetLogger method is implementation of Service interface.
func (a *App) SetLogger(log logger.Logger) app.Service {
	return a
}
