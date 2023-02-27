package common

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

func TextArray(name string, optional bool) ent.Field {
	builder := field.Other(name, pgtype.TextArray{}).SchemaType(map[string]string{
		dialect.Postgres: "TEXT[]",
	})

	if optional {
		builder = builder.Optional()
	}

	return builder
}
