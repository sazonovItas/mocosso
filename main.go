package main

import (
	"fmt"
	"os"

	goversion "github.com/caarlos0/go-version"
	cmd "github.com/sazonovItas/mocosso/cmd/sso"
	"github.com/sazonovItas/mocosso/pkg/logger"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	defer logger.Sync()

	exit := func(code int) {
		if code != 0 {
			panic(fmt.Errorf("exit with code %d", code))
		}
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
