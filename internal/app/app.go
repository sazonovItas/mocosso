package app

import (
	httpapp "github.com/sazonovItas/mocosso/internal/app/http"
	"github.com/sazonovItas/mocosso/internal/config"
	authsvc "github.com/sazonovItas/mocosso/internal/service/auth"
	"github.com/sazonovItas/mocosso/pkg/app"
	"go.uber.org/zap"
)

func New(l *zap.Logger, cfg config.Config) (*app.App, error) {
	const op = "internal.app.New"

	// pgpool, err := postgres.ConnectPool(context.Background(), cfg.Postgres.PostrgresURI)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }
	// _ = pgpool
	//
	// rediscli, err := cacheredis.Connect(cfg.Redis.RedisURI)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }
	// _ = rediscli

	authSvc := authsvc.NewAuthService()

	httpApplication := httpapp.New(l, cfg, authSvc)

	return &app.App{
		Cfg:      cfg,
		Services: []app.Service{httpApplication},
		Cleanups: []func(){},
	}, nil
}
