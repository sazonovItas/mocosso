package postgres

import (
	"sync/atomic"

	"github.com/jackc/pgx/v5"
)

var _ atomic.Pointer[pgx.QueryTracer]
