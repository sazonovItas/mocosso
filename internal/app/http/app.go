package httpapp

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sazonovItas/mocosso/internal/config"
	httpserverv1 "github.com/sazonovItas/mocosso/internal/handler/http/v1"
	authhandlerv1 "github.com/sazonovItas/mocosso/internal/handler/http/v1/auth"
	authsvc "github.com/sazonovItas/mocosso/internal/service/auth"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
)

const (
	// Default timeout to read header of http request.
	defaultReadHeaderTimeout = 2 * time.Second
)

type app struct {
	l   *zap.Logger
	cfg config.Config

	server *http.Server
}

func New(l *zap.Logger, cfg config.Config, authSvc *authsvc.AuthService) *app {
	const op = "internal.app.http.New"

	authHandler := authhandlerv1.NewAuthHandler(authSvc)

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			httpserverv1.RegisterServer(r, authHandler)
		})
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", cfg.Core.Host, cfg.HTTP.Port),
		Handler:           router,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	if err := http2.ConfigureServer(srv, &http2.Server{}); err != nil {
		panic(err)
	}

	return &app{
		l:      l.With(zap.String("app", "http.app")),
		cfg:    cfg,
		server: srv,
	}
}

func (a *app) MustRun(ctx context.Context) {
	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}

func (a *app) Run(ctx context.Context) (err error) {
	const op = "internal.app.http.Run"

	a.l.Info("start running http app", zap.String("address", a.server.Addr))

	if a.cfg.HTTP.SSL {
		err = a.server.ListenAndServeTLS(a.cfg.HTTP.CertPath, a.cfg.HTTP.KeyPath)
	} else {
		err = a.server.ListenAndServe()
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.l.Error("failed to run", zap.Error(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *app) Shutdown() {
	const op = "internal.app.http.Shutdown"

	a.l.Info("shutdown http app")

	shutdownContext, cancel := context.WithTimeout(context.Background(), a.cfg.Core.ShutdownTimeout)
	defer cancel()

	if err := a.server.Shutdown(shutdownContext); err != nil {
		a.l.Error("failed to shutdown", zap.Error(err))
	}
}
