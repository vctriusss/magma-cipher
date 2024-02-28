package random

import (
	"crypto/rand"
	"magma-cipher/internal/key"
	"magma-cipher/internal/utils"
	"math"
)

const MaxUint32 = int64(math.MaxUint32)

func RandKey() key.Key {
	k := key.Key{}

	b := make([]byte, 32)
	rand.Read(b)

	k[0] = utils.BytesToUint32(b[:16])
	k[1] = utils.BytesToUint32(b[16:])

	return k
}
