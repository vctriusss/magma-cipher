package simple

import (
	"reflect"
	"testing"

	"github.com/vctriusss/magma-cipher/internal/key"
)

func TestEncryption(t *testing.T) {
	keyString := "1234567890abcdef0987654321fedcba0987654321fedcba1111111111111111"
	k, _ := key.New([]byte(keyString))

	testCases := []struct {
		desc	string
		input string
		want []byte
		got []byte
	}{
		{
			desc: "one block",
			input: "asdf1234",
			want: []byte{201, 55, 93, 9, 131, 28, 194, 41},
		},
		{
			desc: "one symboled",
			input: "11111111",
			want: []byte{39, 211, 58, 71, 86, 64, 118, 53},
		},
		{
			desc: "padding",
			input: "pad???",
			want: []byte{205, 76, 200, 120, 73, 186, 80, 24},
		},
		{
			desc: "two blocks",
			input: "asdf123411111111",
			want: []byte{201, 55, 93, 9, 131, 28, 194, 41, 39, 211, 58, 71, 86, 64, 118, 53},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Encrypt([]byte(tC.input), k)
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("Want: %v\nGot: %v", tC.want, got)
			}
		})
	}
}