package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PoolOption interface applies options for pgx pool config.
type PoolOption interface {
	Apply(cfg *pgxpool.Config)
}

// option modifies pgx pool config.
type option func(cfg *pgxpool.Config)

// Apply method is implementation of Option interface.
func (op option) Apply(cfg *pgxpool.Config) {
	op(cfg)
}

// WithTracer function returns option that set connection tracer.
func WithTracer(tracer pgx.QueryTracer) option {
	return func(cfg *pgxpool.Config) {
		cfg.ConnConfig.Tracer = tracer
	}
}

// beforeConnFunc is template of before connections function.
type beforeConnFunc func(context.Context, *pgx.ConnConfig) error

// WithBeforeConnect function returns option that modify before connection behavior.
func WithBeforeConnect(before beforeConnFunc) option {
	return func(cfg *pgxpool.Config) {
		cfg.BeforeConnect = before
	}
}

// afterConnFunc is template of after connection function.
type afterConnFunc func(context.Context, *pgx.Conn) error

// WithAfterConnect function returns option that modify after connection behavior.
func WithAfterConnect(after afterConnFunc) option {
	return func(cfg *pgxpool.Config) {
		cfg.AfterConnect = after
	}
}
