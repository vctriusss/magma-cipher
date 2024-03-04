package simple

import (
	"reflect"
	"testing"

	"github.com/vctriusss/magma-cipher/internal/key"
)

func TestDecryption(t *testing.T) {
	keyString := "1234567890abcdef0987654321fedcba0987654321fedcba1111111111111111"
	k, _ := key.New([]byte(keyString))

	testCases := []struct {
		desc	string
		input []byte
		want []byte
		got []byte
	}{
		{
			desc: "one block",
			input: []byte{201, 55, 93, 9, 131, 28, 194, 41},
			want: []byte("asdf1234"),
		},
		{
			desc: "one symboled",
			input: []byte{39, 211, 58, 71, 86, 64, 118, 53},
			want: []byte("11111111"),
		},
		{
			desc: "padding",
			input: []byte{205, 76, 200, 120, 73, 186, 80, 24},
			want: append([]byte("pad???"), 0, 0),
		},
		{
			desc: "two blocks",
			input: []byte{201, 55, 93, 9, 131, 28, 194, 41, 39, 211, 58, 71, 86, 64, 118, 53},
			want: []byte("asdf123411111111"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Decrypt([]byte(tC.input), k)
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("Want: %v\nGot: %v", tC.want, got)
			}
		})
	}
}