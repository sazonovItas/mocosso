package main

import (
	"context"
	"os"

	goversion "github.com/caarlos0/go-version"
	cmd "github.com/sazonovItas/mocosso/cmd/sso"
	"github.com/sazonovItas/mocosso/pkg/logger"
	"go.uber.org/zap"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	defer logger.Sync()

	exit := func(code int) {
		logger.ErrorContext(
			context.Background(),
			"failed to execute command",
			zap.Int("code", code),
		)
	}

	info := buildVersion(version, commit, date)
	cmd.Execute(info, exit, os.Args[1:])
}

func buildVersion(version, commit, date string) goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(
			"mocosso",
			"Run HTTP/GRPC single sign on.",
			"",
		),
		func(i *goversion.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			if date != "" {
				i.BuildDate = date
			}
			if version != "" {
				i.GitVersion = version
			}
		},
	)
}
