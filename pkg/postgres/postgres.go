package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Option interface applies options for pgx pool config.
type Option interface {
	Apply(cfg *pgxpool.Config)
}

// Connect function creates pgx connetions pool with given options and ping db.
func Connect(
	ctx context.Context,
	connString string,
	options ...Option,
) (conn *pgxpool.Pool, err error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from connection string: %w", err)
	}

	for _, option := range options {
		option.Apply(cfg)
	}

	conn, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool with config: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to pint database: %w", err)
	}

	return conn, nil
}

// Connect function creates pgx connetions pool with
// given options and ping db. It panics if any of errors occurred.
func MustConnect(
	ctx context.Context,
	connString string,
	options ...Option,
) (conn *pgxpool.Pool) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(fmt.Errorf("failed to parse config from connection string: %w", err))
	}

	for _, option := range options {
		option.Apply(cfg)
	}

	conn, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create pgx pool with config: %w", err))
	}

	if err := conn.Ping(ctx); err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	return conn
}
