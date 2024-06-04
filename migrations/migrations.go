package migrate

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

// Create runs database migrations.
func Create(dbPool *pgxpool.Pool) error {
	db := stdlib.OpenDB(*dbPool.Config().ConnConfig)

	if err := db.Ping(); err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// выполнение миграций
	if err := goose.Up(db, "migrations/schema"); err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
