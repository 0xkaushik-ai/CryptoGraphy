package aes

import "testing"

// The known-answer values below are taken from the AES specification (FIPS-197
// and the Rijndael design document). Matching them is our proof that the field
// arithmetic is correct, not merely self-consistent.

func TestXtime(t *testing.T) {
	// Repeatedly doubling 0x57 walks 0x57 * 2, * 4, * 8, * 16, as worked out in
	// the AES spec.
	tests := []struct{ in, want byte }{
		{0x57, 0xAE}, // 0x57 * 2  (no overflow)
		{0xAE, 0x47}, // 0x57 * 4  (overflow -> reduce with 0x1B)
		{0x47, 0x8E}, // 0x57 * 8
		{0x8E, 0x07}, // 0x57 * 16 (overflow -> reduce)
		{0x01, 0x02}, // x * 1 = x
		{0x80, 0x1B}, // top bit set: 0x80<<1 = 0x00, then ^0x1B = 0x1B
	}
	for _, tc := range tests {
		if got := xtime(tc.in); got != tc.want {
			t.Errorf("xtime(%#02x) = %#02x, want %#02x", tc.in, got, tc.want)
		}
	}
}

func TestGfMulKnownValues(t *testing.T) {
	tests := []struct{ a, b, want byte }{
		{0x57, 0x13, 0xFE}, // the canonical FIPS-197 example
		{0x57, 0x83, 0xC1}, // another spec example
		{0x57, 0x02, 0xAE}, // multiplying by 2 agrees with xtime
		{0x57, 0x01, 0x57}, // multiplying by 1 is the identity
		{0xFF, 0x00, 0x00}, // multiplying by 0 is 0
	}
	for _, tc := range tests {
		if got := gfMul(tc.a, tc.b); got != tc.want {
			t.Errorf("gfMul(%#02x, %#02x) = %#02x, want %#02x", tc.a, tc.b, got, tc.want)
		}
	}
}

// TestGfAddSelfInverse checks that XOR-addition is its own inverse:
// adding b twice gets you back to a.
func TestGfAddSelfInverse(t *testing.T) {
	for a := 0; a < 256; a++ {
		for b := 0; b < 256; b++ {
			if got := gfAdd(gfAdd(byte(a), byte(b)), byte(b)); got != byte(a) {
				t.Fatalf("gfAdd not self-inverse at a=%#02x b=%#02x", a, b)
			}
		}
	}
}

// TestGfMulFieldLaws checks the algebraic laws a correct field must obey, over
// every pair of bytes. These catch bugs that a handful of fixed examples miss.
func TestGfMulFieldLaws(t *testing.T) {
	for a := 0; a < 256; a++ {
		A := byte(a)
		if got := gfMul(A, 1); got != A {
			t.Fatalf("identity failed: gfMul(%#02x,1) = %#02x", A, got)
		}
		if got := gfMul(A, 0); got != 0 {
			t.Fatalf("zero failed: gfMul(%#02x,0) = %#02x", A, got)
		}
		for b := 0; b < 256; b++ {
			B := byte(b)
			// Commutativity: a*b == b*a.
			if gfMul(A, B) != gfMul(B, A) {
				t.Fatalf("not commutative at a=%#02x b=%#02x", A, B)
			}
			// Distributivity: a*(b+c) == a*b + a*c, tested with c = a.
			left := gfMul(A, gfAdd(B, A))
			right := gfAdd(gfMul(A, B), gfMul(A, A))
			if left != right {
				t.Fatalf("not distributive at a=%#02x b=%#02x", A, B)
			}
		}
	}
}
