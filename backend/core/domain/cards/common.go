package cards

type FilterLevel struct {
	Level    *int32 `query:"level"`
	MaxLevel *int32 `query:"max_level"`
}
