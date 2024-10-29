package drivers

import (
	"context"
	"database/sql"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"alan-tabeo-test-task/src/logging"
)

type PostgreSQL struct {
	DB         *gorm.DB
	Connection *sql.DB
}

// InitPostgreSQL initializes and returns a new postgreSQL connection.
func InitPostgreSQL(_ context.Context, dsn string) (*PostgreSQL, error) {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
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

	return &PostgreSQL{
		DB:         db,
		Connection: connection,
	}, nil
}
