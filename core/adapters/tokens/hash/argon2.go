package hash

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	ARGON2_PARAM_TIME     uint32 = 1
	ARGON2_PARAM_MEMORY   uint32 = 64 * 1024
	ARGON2_PARAM_THREAD   uint8  = 4
	ARGON2_PARAM_HASH_LEN uint32 = 32
	SALT_LEN              int    = 8
)

func Argon2IDHash(data []byte, salt []byte) string {
	hash := argon2.IDKey(data, salt, ARGON2_PARAM_TIME, ARGON2_PARAM_MEMORY, ARGON2_PARAM_THREAD, ARGON2_PARAM_HASH_LEN)

	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)

	argon2Hash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		ARGON2_PARAM_MEMORY, ARGON2_PARAM_TIME, ARGON2_PARAM_THREAD,
		saltEncoded,
		hashEncoded,
	)

	return argon2Hash
}
