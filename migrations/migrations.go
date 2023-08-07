package migrate

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func Create(dbPool *pgxpool.Pool, ctx context.Context) error {

	db := stdlib.OpenDB(*dbPool.Config().ConnConfig)

	defer db.Close()

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

	return nil
}
