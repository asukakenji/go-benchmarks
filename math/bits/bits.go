// Package bits implements bit counting and manipulation
// functions for the predeclared unsigned integer types.
package bits

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	is32Bit  = (^uint(0) >> 32 & 1) == 0
)

// UintSize is the size of a uint in bits.
const UintSize = uintSize

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	if is32Bit {
		return OnesCount32(uint32(x))
	}
	return OnesCount64(uint64(x))
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	return int(pop8tab[x&0xff])
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	return int(pop8tab[x&0xff] + pop8tab[x>>8])
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
