package ulident

import (
	"crypto/rand"
	"io"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

func ID() ent.Field {
	return field.Bytes("id").
		GoType(ulid.ULID{}).
		DefaultFunc(NewULID)
}

func NewULID() ulid.ULID {
	id := ulid.MustNew(ulid.Now(), getEntropy())
	return id
}

var (
	entropy     io.Reader
	entropyOnce sync.Once
)

// getEntropy returns a thread-safe per process monotonically increasing
// entropy source.
func getEntropy() io.Reader {
	entropyOnce.Do(func() {
		entropy = &ulid.LockedMonotonicReader{
			MonotonicReader: ulid.Monotonic(rand.Reader, 0),
		}
	})
	return entropy
}
