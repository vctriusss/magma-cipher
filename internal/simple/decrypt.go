package simple

import (
	"github.com/vctriusss/magma-cipher/internal/block"
	"github.com/vctriusss/magma-cipher/internal/key"
	"github.com/vctriusss/magma-cipher/internal/transform"
	"github.com/vctriusss/magma-cipher/internal/utils"
)

func decryptBlock(blck []byte, k key.Key) []byte {
	b := block.New(blck)

	for i := 0; i < DEC_ORDER_CHANGE_ROUND; i++ {
		b = transform.G(b, k[i%key.PARTS])
	}

	for i := DEC_ORDER_CHANGE_ROUND; i < ROUNDS; i++ {
		b = transform.G(b, k[key.PARTS-1-i%key.PARTS])
	}

	b[0], b[1] = b[1], b[0]

	return b.Bytes()
}

func Decrypt(bytes []byte, key key.Key) []byte {
	byteBlocks := utils.ChunkSlice(bytes, len(bytes)/block.SIZE_BYTES)
	res := make([]byte, 0)

	for _, b := range byteBlocks {
		res = append(res, decryptBlock(b, key)...)
	}

	return res
}
