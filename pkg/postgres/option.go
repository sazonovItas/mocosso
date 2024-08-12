package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// option modifies pgx pool config.
type option func(cfg *pgxpool.Config)

// Apply method is implementation of Option interface.
func (op option) Apply(cfg *pgxpool.Config) {
	op(cfg)
}

// ConnectionTimeouts function returns option that modify connections timeouts settings.
func ConnectionTimeouts(
	connTimeout, idleConnTimeout, healthCheckTimeout, connJitterTimeout time.Duration,
) option {
	return func(cfg *pgxpool.Config) {
		cfg.MaxConnLifetime = connTimeout
		cfg.MaxConnIdleTime = idleConnTimeout
		cfg.HealthCheckPeriod = healthCheckTimeout
		cfg.MaxConnLifetimeJitter = connJitterTimeout
	}
}

// MaxConns functions returns option that modify maximum connections of pgx pool.
func MaxConns(maxConns int) option {
	return func(cfg *pgxpool.Config) {
		cfg.MaxConns = int32(maxConns)
	}
}

// beforeConn is template of before connections function.
type beforeConn func(context.Context, *pgx.ConnConfig) error

// BeforeConnect function returns option that modify before connection behavior.
func BeforeConnect(before beforeConn) option {
	return func(cfg *pgxpool.Config) {
		cfg.BeforeConnect = before
	}
}

// afterConn is template of after connection function.
type afterConn func(context.Context, *pgx.Conn) error

// BeforeConnect function returns option that modify after connection behavior.
func AfterConnect(after afterConn) option {
	return func(cfg *pgxpool.Config) {
		cfg.AfterConnect = after
	}
}
