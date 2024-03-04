package utils

import (
	"encoding/binary"
	"encoding/hex"
)

func ChunkSlice[T any](arr []T, nChunks int) [][]T {
	chunks := make([][]T, nChunks)
	l := (len(arr)-1)/nChunks + 1
	for i := 0; i < nChunks-1; i++ {
		chunks[i] = arr[i*l : (i+1)*l]
	}
	chunks[nChunks-1] = arr[(nChunks-1)*l:]
	return chunks
}

func Pad[T any](arr []T, targetLen int) []T {
	if len(arr) == targetLen {
		return arr
	}
	return append(arr, make([]T, targetLen-len(arr))...)
}

func BytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func Uint32ToBytes(n uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, n)

	return bytes
}

func BytesToHex(bytes []byte) []byte {
	res := make([]byte, 2*len(bytes))
	hex.Encode(res, bytes)

	return res
}
