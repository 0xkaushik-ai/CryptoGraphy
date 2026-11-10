package aes

import "testing"

func TestSboxKnownValues(t *testing.T) {

	tests := []struct{ in, want byte }{
		{0x00, 0x63},
		{0x01, 0x7c},
		{0x10, 0xca},
		{0x53, 0xed},
		{0xff, 0x16},
	}

	for _, tc := range tests {
		if got := sbox[tc.in]; got != tc.want {
			t.Errorf(
				"sbox[%#02x] = %#02x,want %#02x", tc.in, got, tc.want)
		}
	}
}

// the inverse S BOX must undo the forward S-box for every byte

func TestSboxRountTrip(t *testing.T) {
	for i := 0; i < 256; i++ {
		if got := invSbox[sbox[i]]; got != byte(i) {
			t.Fatalf("invSbox[sbox[%#02x]] = %#02x, want %#02x", i, got, i)
		}
	}
}

func TestSboxNoFixedPoints(t *testing.T) {
	for i := 0; i < 256; i++ {
		if sbox[i] == byte(i) {
			t.Errorf("sbox has fixed point at %#02x", i)

		}

	}
}
