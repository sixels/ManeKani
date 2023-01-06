package cards

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"sixels.io/manekani/services/cards/ent"
	"sixels.io/manekani/services/cards/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type CardsRepository struct {
	client *ent.Client
}

func NewRepository(ctx context.Context) (CardsRepository, error) {
	dbUrl := os.Getenv("MANEKANI_DB_URL")

	client, err := open(dbUrl)
	if err != nil {
		return CardsRepository{}, fmt.Errorf("failed to open a connection to postgres: %v", err)
	}

	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		client.Close()
		return CardsRepository{}, fmt.Errorf("failed to migrate: %v", err)
	}

	repo := CardsRepository{
		client: client,
	}

	return repo, nil
}

func (repo CardsRepository) Close() {
	repo.client.Close()
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