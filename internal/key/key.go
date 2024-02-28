package key

import  (
	"magma-cipher/internal/utils"
	"encoding/hex"
	"errors"
)

const (
	SIZE = 256
	SIZE_BYTES = SIZE / 8
	PARTS = SIZE / 32
)

type Key [PARTS]uint32

func New(bytes []byte) (Key, error) {
	key := Key{}
	keyBytes := make([]byte, SIZE_BYTES)

	n, err := hex.Decode(keyBytes, bytes)
	if err != nil || n != SIZE_BYTES {
		return key, errors.New("Key must be a hex string with length of 256 bits (or 64 hex-symbols)")
	}

	for i, ch := range utils.ChunkSlice[byte](keyBytes, PARTS) {
		key[i] = utils.BytesToUint32(ch)
	}

	return key, nil
}
