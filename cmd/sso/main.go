package main

import (
	"context"

	"github.com/sazonovItas/auth-service/internal/app"
	"github.com/sazonovItas/auth-service/internal/config"
	"github.com/sazonovItas/auth-service/pkg/logger"
	"github.com/sazonovItas/auth-service/pkg/logger/zapper"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	l, err := zapper.NewZapLogger(
		logger.ZapLevel(logger.LevelFromString(config.Log.Level)),
		config.Log.Format,
		config.Log.LogPath,
	)
	if err != nil {
		panic(err)
	}

	log := logger.NewZapInterceptor(l)
	application := app.New(log, config)

	go func() {
		err := application.Run(context.TODO())
		if err != nil {
			log.Error("faied to run", "error", err.Error())
		}
	}()

	application.WaitForShutdown()
}
