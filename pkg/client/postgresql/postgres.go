package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Ping(ctx context.Context) error
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectionConfig(cfg Config) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return config, nil
}

func ConnectToDB(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	connCfg, err := ConnectionConfig(cfg)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(ctx, connCfg)

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	logrus.Println("Connection succes")
	return db, nil
}
