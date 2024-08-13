package httpapp

import (
	"context"
	"fmt"

	"github.com/sazonovItas/auth-service/internal/app"
	"github.com/sazonovItas/auth-service/pkg/logger"
)

type App struct {
	log logger.Logger
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (a *App) Shutdown(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (a *App) SetLogger(log logger.Logger) app.Service {
	const op = "app.httpapp.SetLogger"

	if log == nil {
		panic(fmt.Errorf("%s: logger must be not nil value", op))
	}

	a.log = log
	return a
}
