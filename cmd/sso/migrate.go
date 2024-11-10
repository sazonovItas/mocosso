package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sazonovItas/mocosso/migrations"
	"github.com/spf13/cobra"
)

const (
	// connStringFlag is connection string flag for database.
	connStringFlag = "db"

	// driverStringFlag is driver string flag for database.
	driverStringFlag = "driver"

	// pingTimeout is ping timeout for database connection.
	pingTimeout = 5 * time.Second
)

var ErrDBStringNotSpecified = errors.New("connection string is not specified")

type migrateCmd struct {
	cmd  *cobra.Command
	opts migrateOpts
}

type migrateOpts struct {
	connString string
	driver     string
}

func newMigrateCmd() *migrateCmd {
	root := &migrateCmd{}

	cmd := &cobra.Command{
		Use:               "migrate [flags] [command]",
		Short:             "Run migrations on database.",
		Long:              "Run migrations on database for sso service.",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.PersistentFlags().
		StringVar(&root.opts.connString, connStringFlag, "", "Specify connection string for database.")
	cmd.MarkFlagsOneRequired(connStringFlag)
	cmd.PersistentFlags().
		StringVar(&root.opts.driver, driverStringFlag, "pgx", "Specify db driver for connection. Available drivers: pgx.")

	cmd.AddCommand(
		newMigrateUpCmd(&root.opts.connString, &root.opts.driver),
		newMigrateDownCmd(&root.opts.connString, &root.opts.driver),
	)

	root.cmd = cmd
	return root
}

func newMigrateUpCmd(connString, driver *string) *cobra.Command {
	const op = "cmd.sso.migrate"

	cmd := &cobra.Command{
		Use:               "up",
		Short:             "Up migrations on database.",
		Long:              "Up migrations on database for sso service.",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := connectDB(*driver, *connString)
			if err != nil {
				return fmt.Errorf("%s: failed to connect db: %w", op, err)
			}

			goose.SetBaseFS(migrations.FS)
			if err := goose.Up(db, ".", goose.WithNoVersioning()); err != nil {
				return fmt.Errorf("%s: failed to up migrations: %w", op, err)
			}

			return nil
		},
	}

	return cmd
}

func newMigrateDownCmd(connString, driver *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "down",
		Short:             "Down migrations on database.",
		Long:              "Down migrations on database for sso service.",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := connectDB(*driver, *connString)
			if err != nil {
				return fmt.Errorf("failed to connect db: %w", err)
			}

			goose.SetBaseFS(migrations.FS)
			if err := goose.Down(db, ".", goose.WithNoVersioning(), goose.WithAllowMissing()); err != nil {
				return fmt.Errorf("failed to up migrations: %w", err)
			}

			return nil
		},
	}

	return cmd
}

func connectDB(driver, connString string) (*sql.DB, error) {
	const op = "cmd.sso.connectDB"

	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create db from conn string: %w", op, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %w", op, err)
	}

	return db, nil
}
