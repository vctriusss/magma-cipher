package block

import (
	"github.com/vctriusss/magma-cipher/internal/utils"
)

const (
	SIZE       = 64
	SIZE_BYTES = SIZE / 8
	PARTS      = SIZE / 32
)

type Block [PARTS]uint32

func New(bytes []byte) Block {
	block := Block{}
	for i, ch := range utils.ChunkSlice[byte](bytes, PARTS) {
		block[i] = utils.BytesToUint32(ch)
	}

	return block
}

func (b Block) ToBytes() []byte {
	return append(utils.Uint32ToBytes(b[1]), utils.Uint32ToBytes(b[0])...)
}
