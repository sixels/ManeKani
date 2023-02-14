package filters

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type (
	FilterPagination struct {
		Page *uint `query:"page" form:"page"`
	}
	FilterLevels struct {
		Levels *CommaSeparatedInt32 `query:"levels" form:"levels"`
	}
	FilterIDs struct {
		IDs *CommaSeparatedUUID `query:"ids" form:"ids"`
	}
	FilterKinds struct {
		Kinds *CommaSeparatedString `query:"kinds" form:"kinds"`
	}
	FilterSlugs struct {
		Slugs *CommaSeparatedString `query:"slugs" form:"slugs"`
	}
	FilterDecks struct {
		Decks *CommaSeparatedUUID `query:"decks" form:"decks"`
	}
	FilterOwners struct {
		Owners *CommaSeparatedString `query:"owners" form:"owners"`
	}
	FilterSubjects struct {
		Subjects *CommaSeparatedUUID `query:"subjects" form:"subjects"`
	}
	FilterNames struct {
		Names *CommaSeparatedString `query:"names" form:"names"`
	}
)

type SeparateByComma[T any] interface {
	Separate() []T
}

type commaSeparated[T any] string
type CommaSeparatedUUID commaSeparated[uuid.UUID]
type CommaSeparatedString commaSeparated[string]
type CommaSeparatedInt32 commaSeparated[int32]

func (c commaSeparated[T]) string() string {
	return (string)(c)
}

func (c *commaSeparated[T]) Strings() []string {
	if c == nil {
		return nil
	}
	return strings.Split(c.string(), ",")
}

func (c *CommaSeparatedInt32) Separate() (nums []int32) {
	values := (*commaSeparated[any])(c).Strings()
	for _, v := range values {
		if number, err := strconv.Atoi(v); err == nil {
			nums = append(nums, int32(number))
		} else {
			nums = append(nums, -1)
		}
	}
	return nums
}

func (c *CommaSeparatedUUID) Separate() (uuids []uuid.UUID) {
	values := (*commaSeparated[any])(c).Strings()
	for _, v := range values {
		if id, err := uuid.Parse(v); err == nil {
			uuids = append(uuids, id)
		} else {
			uuids = append(uuids, uuid.Nil)
		}
	}
	return uuids
}

func (c *CommaSeparatedString) Separate() []string {
	return (*commaSeparated[any])(c).Strings()
}

// Applies a filter to a filter list
func ApplyFilter[P any, T any](
	fltrs []P,
	values []T,
	filter func(...T) P,
) []P {
	if len(values) > 0 {
		fltrs = append(fltrs, filter(values...))
	}
	return fltrs
}
