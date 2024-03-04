package block

import "testing"

func Test(t *testing.T) {
	arr := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	exp := [2]uint32{67305985, 134678021}
	got := New(arr)
	if got != exp {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}
