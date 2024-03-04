package key

import "testing"

func TestZeros(t *testing.T) {
	got, _ := New(make([]byte, 32))
	if got != [8]uint32{0, 0, 0, 0, 0, 0, 0, 0} {
		t.Errorf("Error")
	}
}

func TestRandom(t *testing.T) {
	bytes := []byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 97, 98, 99, 100, 101, 102, 48, 57, 56, 55, 54, 53, 52, 51, 50, 49, 102, 101, 100, 99, 98, 97, 48, 57, 56, 55, 54, 53, 52, 51, 50, 49, 102, 101, 100, 99, 98, 97, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49}
	expected := [8]uint32{2018915346, 4023233424, 1130727177, 3135045153, 1130727177, 3135045153, 286331153, 286331153}
	got, _ := New(bytes)
	if got != expected {
		t.Errorf("Excpected %v, got %v", expected, got)
	}
}
