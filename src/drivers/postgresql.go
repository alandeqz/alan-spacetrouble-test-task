package drivers

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"alan-tabeo-test-task/src/config"
	"alan-tabeo-test-task/src/logging"
)

type PostgreSQL struct {
	DB         *gorm.DB
	Connection *sql.DB
}

// InitPostgreSQL initializes and returns a new postgreSQL connection.
func InitPostgreSQL(_ context.Context, cfg *config.Configuration, baseFS embed.FS) (*PostgreSQL, error) {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  cfg.DSN,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
	)
	if err != nil {
		slog.Error("failed to connect to PostgreSQL", logging.Error, err.Error())

		return nil, err
	}

	connection, err := db.DB()
	if err != nil {
		slog.Error("failed to get the database connection", logging.Error, err.Error())

		return nil, err
	}

	slog.Info("successfully connected to PostgreSQL")

	if err = runMigrations(connection, cfg.SchemaName, "migrations", baseFS); err != nil {
		return nil, err
	}

	return &PostgreSQL{
		DB:         db,
		Connection: connection,
	}, nil
}

// runMigrations runs the migrations.
func runMigrations(db *sql.DB, schemaName, migrationsDirectory string, baseFS embed.FS) error {
	goose.SetBaseFS(baseFS)

	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("failed to set the dialect", logging.Error, err.Error())

		return err
	}

	goose.SetTableName(fmt.Sprintf("%s.goose_db_version", schemaName))

	if err := goose.Up(db, migrationsDirectory); err != nil {
		slog.Error("failed to run the migrations", logging.Error, err.Error())

		return err
	}

	return nil
}
