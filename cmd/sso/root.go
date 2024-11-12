package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	goversion "github.com/caarlos0/go-version"
	"github.com/sazonovItas/mocosso/internal/app"
	"github.com/sazonovItas/mocosso/internal/config"
	configloader "github.com/sazonovItas/mocosso/pkg/config"
	"github.com/sazonovItas/mocosso/pkg/logger"
	"github.com/spf13/cobra"
)

type rootCmd struct {
	cmd  *cobra.Command
	opts rootOpts
	exit func(int)
}

type rootOpts struct {
	configPath string
}

func Execute(version goversion.Info, exit func(int), args []string) {
	newRootCmd(version, exit).Execute(args)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		cmd.exit(1)
	}
}

func newRootCmd(version goversion.Info, exit func(int)) *rootCmd {
	const op = "cmd.root.Run"

	root := &rootCmd{
		exit: exit,
	}

	cmd := &cobra.Command{
		Use:               "mocosso [command]",
		Short:             "Run HTTP/GRPC sso.",
		Long:              "Run HTTP/GRPC Single Sign-on.",
		Version:           version.String(),
		SilenceUsage:      false,
		SilenceErrors:     false,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			var cfg config.Config
			if err = configloader.Load(
				&cfg,
				root.opts.configPath,
				true,
				configloader.WithEnvs(config.ConfigEnvPrefix),
				configloader.WithConfigName(config.ConfigFileName),
				configloader.WithConfigType("yaml"),
				configloader.WithConfigPaths(config.ConfigHomeDir, config.ConfigEtcDir, "."),
				configloader.WithDefaults(map[string]any{
					// core section defaults
					"core.env":              "development",
					"core.service_name":     "mocosso",
					"core.shutdown_timeout": 20 * time.Second,

					// auth sectionj defaults
					"auth.access_token_liveness":  15 * time.Minute,
					"auth.refresh_token_liveness": 30 * 24 * time.Hour,

					// log section defaults
					"log.encoding":   "json",
					"log.level":      "info",
					"log.error_log":  "stderr",
					"log.access_log": "stdout",

					// http section defaults
					"http.enabled": true,
					"http.port":    "8080",
					"http.ssl":     false,

					// grpc section defaults
					"grpc.enabled": false,
					"grpc.port":    "9090",
					"grpc.ssl":     true,
				}),
			); err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}

			if err = logger.ConfigureLogger(
				logger.WithLevel(logger.ParseLevel(cfg.Log.Level)),
				logger.WithEncoding(cfg.Log.Encoding),
				logger.WithOutputPaths([]string{cfg.Log.AccessLogPath}),
				logger.WithErrorOutputPaths([]string{cfg.Log.ErrorLogPath}),
			); err != nil {
				return fmt.Errorf("%s: failed to configure logger: %w", op, err)
			}

			application, err := app.New(logger.CreateLogger(), cfg)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}

			ctx, stop := signal.NotifyContext(
				context.Background(),
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT,
			)
			defer stop()

			if err = application.Run(ctx); err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			defer application.Cleanup()

			return nil
		},
	}
	cmd.SetVersionTemplate("{{ .Version }}")
	cmd.PersistentFlags().
		StringVarP(&root.opts.configPath, "config", "c", "", "Specify path to config file.")

	cmd.AddCommand(newMigrateCmd().cmd)

	root.cmd = cmd
	return root
}
