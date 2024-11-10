package app

import (
	"context"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type (
	Service interface {
		Run(ctx context.Context) error
		Shutdown()
	}
)

type App struct {
	Cfg      any
	Services []Service
	Cleanups []func()

	ctx    context.Context
	cancel context.CancelFunc

	cleanUpDone atomic.Int32
}

func (a *App) MustRun(ctx context.Context) {
	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}

func (a *App) Run(parentCtx context.Context) error {
	a.ctx, a.cancel = context.WithCancel(parentCtx)

	wg, ctx := errgroup.WithContext(a.ctx)
	for _, svc := range a.Services {
		wg.Go(func() error {
			return svc.Run(ctx)
		})
	}
	<-ctx.Done()

	a.shutdownServices()

	return wg.Wait()
}

func (a *App) Shutdown() {
	if a.cancel != nil {
		a.cancel()
	}
}

func (a *App) Cleanup() {
	if !a.cleanUpDone.CompareAndSwap(0, 1) {
		return
	}

	for _, cleanup := range a.Cleanups {
		cleanup()
	}
}

func (a *App) shutdownServices() {
	for _, svc := range a.Services {
		svc.Shutdown()
	}
}
