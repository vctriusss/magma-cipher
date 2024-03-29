package key

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/vctriusss/magma-cipher/internal/utils"
)

const (
	SIZE       = 256
	SIZE_BYTES = SIZE / 8
	PARTS      = SIZE / 32
)

type Key [PARTS]uint32

func New(bytes []byte) (Key, error) {
	key := Key{}
	keyBytes := make([]byte, SIZE_BYTES)

	n, err := hex.Decode(keyBytes, bytes)
	if err != nil || n != SIZE_BYTES {
		return key, errors.New("Key must be a hex string with length of 256 bits (or 64 hex-symbols)")
	}

	for i, ch := range utils.ChunkSlice(keyBytes, PARTS) {
		key[i] = utils.BytesToUint32(ch)
	}

	return key, nil
}

func Generate() (Key, error) {
	keyBytes := make([]byte, SIZE_BYTES)

	_, err := rand.Read(keyBytes)
	if err != nil {
		return Key{}, err
	}

	return New(utils.BytesToHex(keyBytes))
}

func (k Key) Bytes() []byte {
	res := make([]byte, 0)

	for _, b := range k {
		res = append(res, utils.Uint32ToBytes(b)...)
	}

	return res
}
