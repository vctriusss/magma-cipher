package transform

import (
	"github.com/vctriusss/magma-cipher/internal/block"
	"github.com/vctriusss/magma-cipher/internal/utils"
)

var PI = [8][16]byte{
	{12, 4, 6, 2, 10, 5, 11, 9, 14, 8, 13, 7, 0, 3, 15, 1},
	{6, 8, 2, 3, 9, 10, 5, 12, 1, 14, 4, 7, 11, 13, 0, 15},
	{11, 3, 5, 8, 2, 15, 10, 13, 14, 1, 7, 4, 12, 9, 6, 0},
	{12, 8, 2, 1, 13, 4, 15, 6, 7, 0, 10, 5, 3, 14, 9, 11},
	{7, 15, 5, 10, 8, 1, 6, 13, 0, 9, 3, 14, 11, 4, 2, 12},
	{5, 13, 15, 6, 9, 2, 12, 10, 11, 7, 8, 1, 4, 3, 14, 0},
	{8, 14, 2, 5, 6, 9, 1, 12, 15, 4, 11, 0, 13, 10, 3, 7},
	{1, 7, 14, 13, 0, 5, 8, 3, 4, 15, 10, 6, 9, 12, 11, 2},
}

func T(bi uint32) uint32 {
	bytes := utils.Uint32ToBytes(bi)

	for i := 0; i < 4; i++ {
		var p1 byte = PI[i*2][bytes[i]&0x0f]
		var p2 byte = PI[i*2+1][(bytes[i]&0xf0)>>4]

		bytes[i] = (p2 << 4) | p1
	}
	return utils.BytesToUint32(bytes)
}

func G(b block.Block, ki uint32) block.Block {
	tmp := b[1]

	b[1] += ki
	b[1] = T(b[1])
	b[1] = rol(b[1], 11)
	b[1] = b[0] ^ b[1]
	b[0] = tmp
	return b
}

func sumMod32(n1, n2 uint32) uint32 {
	res := make([]byte, 4)

	b1 := utils.Uint32ToBytes(n1)
	b2 := utils.Uint32ToBytes(n2)

	var tmp uint32
	for i := 3; i >= 0; i-- {
		tmp = tmp>>8 + uint32(b1[i]+b2[i])
		res[i] = byte(tmp & 0xff)
	}

	return utils.BytesToUint32(res)
}

func rol(n uint32, nPos int) uint32 {
	return ((n << 11) & (1<<32 - 1)) | ((n >> (32 - 11)) & (1<<32 - 1))
}
