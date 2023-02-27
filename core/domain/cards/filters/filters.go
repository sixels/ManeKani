package filters

import (
	"strconv"
	"strings"
	"time"

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
	FilterCards struct {
		Cards *CommaSeparatedUUID `query:"cards" form:"cards"`
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
	FilterCreatedTime struct {
		CreatedAfter  *time.Time `query:"created_after" form:"created_after"`
		CreatedBefore *time.Time `query:"created_before" form:"created_before"`
	}
)

type SeparateByComma[T any] interface {
	Separate() []T
	Only() *T
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

func (c *CommaSeparatedInt32) Only() *int32 {
	values := (*commaSeparated[any])(c).Strings()
	if len(values) == 0 {
		return nil
	}

	num := int32(-1)
	if number, err := strconv.Atoi(values[0]); err == nil {
		num = int32(number)
	}
	return &num
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

func (c *CommaSeparatedUUID) Only() *uuid.UUID {
	values := (*commaSeparated[any])(c).Strings()
	if len(values) == 0 {
		return nil
	}

	uid := uuid.Nil
	if id, err := uuid.Parse(values[0]); err == nil {
		uid = id
	}
	return &uid
}

func (c *CommaSeparatedString) Separate() []string {
	return (*commaSeparated[any])(c).Strings()
}

func (c *CommaSeparatedString) Only() *string {
	values := (*commaSeparated[any])(c).Strings()
	if len(values) == 0 {
		return nil
	}

	return &values[0]
}

type Filter[P any] []P

func NewFilter[P any](init []P) *Filter[P] {
	return (*Filter[P])(&init)
}

func (f *Filter[P]) Filters() []P {
	return ([]P)(*f)
}

func With[P, T any](f *Filter[P], value *T, filter func(T) P) *Filter[P] {
	if value != nil {
		filters := f.Filters()
		*f = (Filter[P])(append(filters, filter(*value)))
	}
	return f
}

func In[P, T any](f *Filter[P], values []T, filter func(...T) P) *Filter[P] {
	if len(values) > 0 {
		filters := f.Filters()
		*f = (Filter[P])(append(filters, filter(values...)))
	}
	return f
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
