package transactions

import (
	"context"
)

type TransactionalRepository interface {
	BeginTransaction(ctx context.Context) (TransactionalRepository, error)
	Rollback() error
	Commit() error
}

type Transaction struct {
	repos []TransactionalRepository
	ctx   context.Context
}

func (tx *Transaction) registerRepository(repo TransactionalRepository) {
	tx.repos = append(tx.repos, repo)
}

func (tx *Transaction) Rollback() {
	for _, repo := range tx.repos {
		repo.Rollback()
	}
}

func (tx *Transaction) Commit() {
	for _, repo := range tx.repos {
		repo.Commit()
	}
}

func (tx *Transaction) Run(txFn func(ctx context.Context) error) error {
	_, err := RunWithResult(tx, func(ctx context.Context) (struct{}, error) {
		return struct{}{}, txFn(ctx)
	})
	return err
}

func RunWithResult[R any](tx *Transaction, txFn func(ctx context.Context) (R, error)) (result R, err error) {
	defer func() {
		if unwind := recover(); unwind != nil {
			tx.Rollback()
			panic(unwind)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	result, err = txFn(tx.ctx)
	if err != nil {
		tx.Rollback()
		return result, err
	}
	return result, err
}

func Begin(ctx context.Context) *Transaction {
	return &Transaction{
		ctx:   ctx,
		repos: []TransactionalRepository{},
	}
}

func MakeTransactional[R TransactionalRepository](tx *Transaction, repo R) (R, error) {
	txRepo, err := repo.BeginTransaction(tx.ctx)
	if err != nil {
		return txRepo.(R), err
	}
	tx.registerRepository(txRepo)
	return txRepo.(R), nil
}
