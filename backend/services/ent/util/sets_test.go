package util_test

import (
	"testing"

	"sixels.io/manekani/services/ent/util"

	"github.com/stretchr/testify/assert"
)

func TestDiffs(t *testing.T) {
	a := []string{"foo", "bar", "baz"}
	b := []string{"foo", "bar", "ba!"}
	dups := []string{"foo", "foo"}
	dupsExclusive := []string{"foo", "ba!", "ba!"}

	assert.Equal(t,
		[]string{"ba!"}, util.DiffStrings(a, b))
	assert.Equal(t,
		[]string(nil), util.DiffStrings(a, dups))
	assert.Equal(t,
		[]string{"bar", "baz"}, util.DiffStrings(dups, a))
	assert.Equal(t,
		[]string{"ba!"}, util.DiffStrings(a, dupsExclusive))
	assert.Equal(t,
		[]string{"bar", "baz"}, util.DiffStrings(dupsExclusive, a))
}
