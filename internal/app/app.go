package app

import (
	"context"
	"fmt"

	"github.com/sazonovItas/mocosso/pkg/app"
	"github.com/sazonovItas/mocosso/pkg/postgres"
	cacheredis "github.com/sazonovItas/mocosso/pkg/redis"
	"go.uber.org/zap"
)

func New(log *zap.Logger, cfg Config) (*app.App, error) {
	const op = "internal.app.New"

	pgpool, err := postgres.ConnectPool(context.Background(), cfg.Postgres.PostrgresURI)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_ = pgpool

	rediscli, err := cacheredis.Connect(cfg.Redis.RedisURI)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_ = rediscli

	return &app.App{
		Cfg: cfg,
	}, nil
}
