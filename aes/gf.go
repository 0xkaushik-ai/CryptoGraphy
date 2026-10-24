// The whole job os this file is to implement the Galois field GF(2^8) used in AES. The field is defined by the irreducible polynomial x^8 + x^4 + x^3 + x + 1. The elements of the field are represented as bytes (0-255), and the operations of addition and multiplication are defined modulo this polynomial.
// This is because polynomial over a finite field give you arithmetic where every operation is reversible ans mixes bits throughly,which is what precisly cipher need .

// The build order is these three functions, which are the core of the Galois field arithmetic:
// add - XOR(the floor everything stands on)
// xtime - multiply x, with the 0*1B reduction(the core trick)
// gfMul - general multipy butild from xtime and add

package aes

// variable
const reductionPoly = 0x1B // The irreducible polynomial x^8 + x^4 + x^3 + x + 1, represented as a byte

// Designing a function

func gfAdd(a, b byte) byte {
	return a ^ b // Addition in GF(2^8) is just XOR
}

// here xtime multiplies a by x. Multiplying a polynomical by x shifts every power up by one . which is left shift  of the byte
func xtime(a byte) byte {
	highBitSet := a&0x80 != 0 // Check if the high bit is set
	a <<= 1                   // a is byte uint8 so the shift already drops any  9th bit
	if highBitSet {
		a ^= reductionPoly // If the high bit was set, reduce modulo the irreducible polynomial
	}

	return a
}

func gfMul(a, b byte) byte {
	var product byte
	for b != 0 {
		if b&1 != 0 { // the loweset remaining bit of b is set
			product ^= a // add the current power of two multiple of a
		}
		a = xtime(a)
		b >>= 1 // shift b to the right, effectively dividing by 2
	}

	return product
}
