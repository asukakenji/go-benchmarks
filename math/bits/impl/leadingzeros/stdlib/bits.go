package stdlib

import "github.com/asukakenji/go-benchmarks"

// --- LeadingZeros ---

// LeadingZeros returns the number of leading zero bits in x; the result is the size of uint in bits for x == 0.
func LeadingZeros(x uint) int { return int(benchmarks.LogSizeOfUintInBytes) - Len(x) }

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int { return 8 - Len8(x) }

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int { return 16 - Len16(x) }

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int { return 32 - Len32(x) }

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int { return 64 - Len64(x) }

// LeadingZerosPtr returns the number of leading zero bits in x; the result is the size of uintptr in bits for x == 0.
func LeadingZerosPtr(x uintptr) int { return int(benchmarks.LogSizeOfUintInBytes) - LenPtr(x) }

// --- Len ---

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len(x uint) int {
	if benchmarks.LogSizeOfUintInBytes == 32 {
		return Len32(uint32(x))
	}
	return Len64(uint64(x))
}

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len8(x uint8) int {
	return int(len8tab[x])
}

// Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len16(x uint16) (n int) {
	if x >= 1<<8 {
		x >>= 8
		n = 8
	}
	return n + int(len8tab[x])
}

// Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len32(x uint32) (n int) {
	if x >= 1<<16 {
		x >>= 16
		n = 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(len8tab[x])
}

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(len8tab[x])
}

// LenPtr returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func LenPtr(x uintptr) int {
	if benchmarks.LogSizeOfUintInBytes == 32 {
		return Len32(uint32(x))
	}
	return Len64(uint64(x))
}
