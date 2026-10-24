// Package aes is a from-scratch, readable implementation of the AES block
// cipher (the Advanced Encryption Standard, FIPS-197), written to learn how
// the algorithm actually works.
//
// # Do not use this in production
//
// This package is for study only. The Go standard library's crypto/aes is the
// one to use for real data, because it is:
//
//   - constant-time (this package uses table/branch operations that can leak
//     key material through timing and CPU-cache side channels), and
//   - hardware-accelerated (crypto/aes uses the CPU's AES-NI instructions and
//     is dozens of times faster).
//
// The goal here is clarity: given the same key and input, this package produces
// exactly the same bytes as crypto/aes and as the official NIST test vectors,
// but the code is written to be read and understood rather than to be safe or
// fast.
//
// The implementation is built bottom-up:
//
//	gf.go     finite-field arithmetic over GF(2^8)   <- you are here (Stage 1)
//	sbox.go   the substitution box
//	key.go    key expansion (the key schedule)
//	cipher.go the round functions and full block encrypt/decrypt
package aes
