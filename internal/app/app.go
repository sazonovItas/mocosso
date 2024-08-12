package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sazonovItas/auth-service/internal/config"
	"github.com/sazonovItas/auth-service/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type App struct {
	log logger.Logger
	cfg *config.ConfigYaml

	services []Service
}

// New function create a app with given logger and shutdown timeout.
func New(l logger.Logger, cfg *config.ConfigYaml) *App {
	return &App{
		log: l,
		cfg: cfg,
	}
}

// Add method append new service to app.
func (a *App) Add(service Service) {
	a.services = append(a.services, service)
}

// Run method start all services and wait until all done.
func (a *App) Run(ctx context.Context) error {
	a.log.Info("starting app")

	g, ctx := errgroup.WithContext(ctx)
	for _, service := range a.services {
		g.Go(func() error {
			return service.Run(ctx)
		})
	}

	return g.Wait()
}

// Stop method shutdown all servcies and wait until all stops.
func (a *App) Shutdown(ctx context.Context) error {
	a.log.Info("stopping app")

	g, ctx := errgroup.WithContext(ctx)
	for _, service := range a.services {
		g.Go(func() error {
			return service.Shutdown(ctx)
		})
	}

	return g.Wait()
}

// WaitForShutdown method waits for termination signal and then shutdown all services.
func (a *App) WaitForShutdown() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Core.ShutdownTimeout)
	defer cancel()

	err := a.Shutdown(ctx)
	if err != nil {
		a.log.Error("error during stopping app", "error", err.Error())
	}
}
