package transform

import (
	"testing"
)

func TestROL(t *testing.T) {
	var n uint32 = 0b00000001111111111111111111111111
	var exp uint32 = 0b11111111111111111111100000001111
	got := rol(n, 11)
	if exp != got {
		t.Errorf("exp: %b, got: %b", exp, got)
	}
}

func TestT(t *testing.T) {
	got := T(12345)
	exp := uint32(408361784)

	if exp != got {
		t.Errorf("exp: %d, got: %d", exp, got)
	}
}
