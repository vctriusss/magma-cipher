package simple

import (
	"magma-cipher/internal/block"
	"magma-cipher/internal/key"
	"magma-cipher/internal/transform"
	"magma-cipher/internal/utils"
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

	return b.ToBytes()
}

func Decrypt(bytes []byte, key key.Key) []byte {
	byteBlocks := utils.ChunkSlice[byte](bytes, len(bytes) / block.SIZE_BYTES)
	res := make([]byte, 0)
	
	for _, b := range byteBlocks {
		res = append(res, decryptBlock(b, key)...)
	}

	return res
}
