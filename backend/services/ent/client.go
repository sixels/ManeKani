package ent

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type EntClient struct {
	*ent.Client
}

func Connect() (*EntClient, error) {
	dbUrl := os.Getenv("MANEKANI_DB_URL")

	client, err := open(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open a connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}

	return &EntClient{
		Client: client,
	}, nil
}

// Open a new connection
func open(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
