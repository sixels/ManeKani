package util

import "github.com/jackc/pgtype"

// Convert a string array into a postgres TextArray
func IntoPgTextArray(elements []string) pgtype.TextArray {
	status := pgtype.Null
	if elements != nil {
		status = pgtype.Present
	}

	return pgtype.TextArray{
		Status: status,
		Dimensions: []pgtype.ArrayDimension{
			{
				Length:     int32(len(elements)),
				LowerBound: 1,
			},
		},
		Elements: mapText(elements),
	}
}

func FromPgTextArray(txtArray pgtype.TextArray) []string {
	if txtArray.Status != pgtype.Present {
		return nil
	}

	array := make([]string, len(txtArray.Elements))

	for i, e := range txtArray.Elements {
		array[i] = e.String
	}

	return array
}

// Maps a string array to a postgres Text array
func mapText(elements []string) []pgtype.Text {
	mapping := make([]pgtype.Text, len(elements))

	for i, e := range elements {
		mapping[i] = pgtype.Text{String: e, Status: pgtype.Present}
	}

	return mapping
}
