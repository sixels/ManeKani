package ent

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/sixels/manekani/core/ports/transactions"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Client = ent.AnyClient

type EntRepository struct {
	Client
}

func NewRepository() (*EntRepository, error) {
	dbUrl := os.Getenv("MANEKANI_DB_URL")

	client, err := open(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open a connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}

	return &EntRepository{
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

func (repo *EntRepository) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	client, ok := repo.Client.(*ent.Client)
	if !ok {
		return nil, fmt.Errorf("Cannot start a transaction out of a transaction")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	txRepo := *repo
	txRepo.Client = tx

	return &txRepo, nil
}

func (repo *EntRepository) Rollback() error {
	tx, ok := repo.Client.(*ent.Tx)
	if !ok {
		return fmt.Errorf("Can't rollback because this is not a transaction")
	}
	return tx.Rollback()
}
func (repo *EntRepository) Commit() error {
	tx, ok := repo.Client.(*ent.Tx)
	if !ok {
		return nil
	}
	return tx.Commit()
}
