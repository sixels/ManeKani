package db

import (
	"context"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(db_url string) (*DB, error) {

	config, err := pgxpool.ParseConfig(db_url)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &DB{
		pool: pool,
	}, nil
}
