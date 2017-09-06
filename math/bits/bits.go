// Package bits implements bit counting and manipulation
// functions for the predeclared unsigned integer types.
package bits

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	is32Bit  = (^uint(0) >> 32 & 1) == 0
)

// UintSize is the size of a uint in bits.
const UintSize = uintSize

// --- LeadingZeros ---

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int {
	return 0
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	return 0
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	return 0
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	return 0
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	return 0
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
