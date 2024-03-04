package utils

import "testing"
import "reflect"


func TestNotEq(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	exp := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	got := ChunkSlice(arr, 3)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}

func TestEq(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	exp := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	got := ChunkSlice(arr, 3)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}

func TestUintToBytes(t *testing.T) {
	var n uint32 = 0b10101010111111110000000011110000
	exp := []byte{0b11110000, 0b00000000, 0b11111111, 0b10101010}
	got := Uint32ToBytes(n)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("exp: %b, got: %b", exp, got)
	}
}

func TestBytesToUint(t *testing.T) {
	bytes := []byte{0b10101010, 0b11111111, 0b00000000, 0b11110000}
	var exp uint32 = 0b11110000000000001111111110101010
	got := BytesToUint32(bytes)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("exp: %b, got: %b", exp, got)
	}
}

func TestTwoWays(t *testing.T) {
	var n uint32 = 0b10101010111111110000000011110000
	bytes := []byte{0b10101010, 0b11111111, 0b00000000, 0b11110000}
	if n != BytesToUint32(Uint32ToBytes(n)) {
		t.Errorf("error BUB")
	}

	if !reflect.DeepEqual(bytes, Uint32ToBytes(BytesToUint32(bytes))) {
		t.Errorf("error UBU")
	}
}