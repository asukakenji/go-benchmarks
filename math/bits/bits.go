// Package bits implements bit counting and manipulation
// functions for the predeclared unsigned integer types.
package bits

import "github.com/asukakenji/go-benchmarks"

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	is32Bit  = (^uint(0) >> 32 & 1) == 0
)

// UintSize is the size of a uint in bits.
const UintSize = uintSize

// --- LeadingZeros ---

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int {
	return int(benchmarks.SizeOfUintInBits) - Len(x)
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	return int(nlz8tab[x])
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	n := nlz8tab[x>>8]
	if n != 8 {
		return int(n)
	}
	return int(n + nlz8tab[x&0xff])
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	n := nlz8tab[x>>24]
	if n != 8 {
		return int(n)
	}
	n += nlz8tab[(x>>16)&0xff]
	if n != 16 {
		return int(n)
	}
	n += nlz8tab[(x>>8)&0xff]
	if n != 24 {
		return int(n)
	}
	return int(n + nlz8tab[x&0xff])
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	return 64 - Len64(x)
}

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len(x uint) int {
	if benchmarks.SizeOfUintInBits == 32 {
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

// --- TrailingZeros ---

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
func TrailingZeros(x uint) int {
	return 0
}

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func TrailingZeros8(x uint8) int {
	return 0
}

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func TrailingZeros16(x uint16) int {
	return 0
}

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int {
	return 0
}

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint) int {
	return 0
}

// --- OnesCount ---

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	if is32Bit {
		return OnesCount32(uint32(x))
	}
	return OnesCount64(uint64(x))
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	return int(((((uint32(x) * 0x08040201) >> 3) & 0x11111111) * 0x11111111) >> 28)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	return int((((uint64(x&0x7fff)*0x0002000400080010)&0x1111111111111111)*0x1111111111111111)>>60) + int((x>>15)&0x1)
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return int((x * 0x01010101) >> 24)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	return int((x * 0x0101010101010101) >> 56)
}
