package app

import (
	"context"

	"github.com/sazonovItas/auth-service/pkg/logger"
)

type Service interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error

	SetLogger(l logger.Logger) Service
}
