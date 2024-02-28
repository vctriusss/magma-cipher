package random

import (
	"crypto/rand"
	"math"

	"github.com/vctriusss/magma-cipher/internal/key"
	"github.com/vctriusss/magma-cipher/internal/utils"
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
