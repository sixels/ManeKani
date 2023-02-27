package util

import (
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/ent"

	"github.com/jackc/pgtype"
)

func UpdateValue[T any, U any](value *T, setter func(T) U) {
	if value != nil {
		setter(*value)
	}
}

func UpdateValues[T any, U any](value *[]T, setter func(...T) U) {
	if value != nil {
		setter((*value)...)
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

func Ptr[T any](t T) *T {
	return &t
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
