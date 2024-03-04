package simple

import (
	"github.com/vctriusss/magma-cipher/internal/block"
	"github.com/vctriusss/magma-cipher/internal/key"
	"github.com/vctriusss/magma-cipher/internal/transform"
	"github.com/vctriusss/magma-cipher/internal/utils"
)

const (
	ROUNDS                 = 32
	ENC_ORDER_CHANGE_ROUND = 24
	DEC_ORDER_CHANGE_ROUND = ROUNDS - ENC_ORDER_CHANGE_ROUND
)

func encryptBlock(blck []byte, k key.Key) []byte {
	b := block.New(blck)

	for i := 0; i < ENC_ORDER_CHANGE_ROUND; i++ {
		b = transform.G(b, k[i%key.PARTS])
	}

	for i := ENC_ORDER_CHANGE_ROUND; i < ROUNDS; i++ {
		b = transform.G(b, k[key.PARTS-1-i%key.PARTS])
	}

	b[0], b[1] = b[1], b[0]

	return b.Bytes()
}

func Encrypt(bytes []byte, key key.Key) []byte {
	var targetLen = ((len(bytes)-1)/block.SIZE_BYTES + 1) * block.SIZE_BYTES

	bytes = utils.Pad(bytes, targetLen)
	byteBlocks := utils.ChunkSlice(bytes, targetLen/block.SIZE_BYTES)
	res := make([]byte, 0)

	for _, b := range byteBlocks {
		res = append(res, encryptBlock(b, key)...)
	}
	return res
}
