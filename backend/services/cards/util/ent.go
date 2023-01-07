package util

import (
	"context"
	"fmt"

	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/services/cards/ent"

	"github.com/jackc/pgtype"
)

func UpdateValue[T any, U any](value *T, setter func(T) U) {
	if value != nil {
		setter(*value)
	}
}

func UpdateTextArray[U any](value *[]string, setter func(pgtype.TextArray) U) {
	if value != nil {
		setter(ToPgTextArray(*value))
	}
}

func MapArray[T any, U any](values []T, mapper func(T) U) []U {
	array := make([]U, len(values))

	for i, val := range values {
		array[i] = mapper(val)
	}

	return array
}

func ParseEntError(err error) *errors.Error {
	var domainError *errors.Error
	switch e := err.(type) {
	case *ent.ValidationError:
		domainError = errors.InvalidRequest(e.Error())
	case *ent.NotFoundError:
		domainError = errors.NotFound(e.Error())
	case *ent.ConstraintError:
		domainError = errors.Conflict(e.Error())
	case *ent.NotSingularError:
		domainError = errors.Unknown(e)
	default:
		domainError = errors.Unknown(err)
	}
	return domainError
}

func WithTx[T any](ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (*T, error)) (*T, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	ret, err := fn(tx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return ret, commit(tx)
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
