package simple

import (
	"magma-cipher/internal/block"
	"magma-cipher/internal/key"
	"magma-cipher/internal/transform"
	"magma-cipher/internal/utils"
)

const (
	ROUNDS                 = 32
	ENC_ORDER_CHANGE_ROUND = 24
	DEC_ORDER_CHANGE_ROUND = 8
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

	return b.ToBytes()
}

func Encrypt(bytes []byte, key key.Key) []byte {
	var targetLen = ((len(bytes)-1)/block.SIZE_BYTES + 1) * block.SIZE_BYTES

	bytes = utils.Pad[byte](bytes, targetLen)
	byteBlocks := utils.ChunkSlice[byte](bytes, targetLen/block.SIZE_BYTES)
	res := make([]byte, 0)

	for _, b := range byteBlocks {
		res = append(res, encryptBlock(b, key)...)
	}
	return res
}
