package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectPool function creates pgx connetions pool with given options and ping db.
func ConnectPool(
	ctx context.Context,
	connString string,
	options ...PoolOption,
) (pool *pgxpool.Pool, err error) {
	const op = "pkg.postgres.ConnectPool"

	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse config from connection string: %w", op, err)
	}

	for _, option := range options {
		option.Apply(cfg)
	}

	pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create pgx pool with config: %w", op, err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %w", op, err)
	}

	return pool, nil
}

// MustConnectPool function creates pgx connetions pool with
// given options and ping db. It panics if any of errors occurred.
func MustConnectPool(
	ctx context.Context,
	connString string,
	options ...PoolOption,
) (pool *pgxpool.Pool) {
	pool, err := ConnectPool(ctx, connString, options...)
	if err != nil {
		panic(err)
	}

	return pool
}
