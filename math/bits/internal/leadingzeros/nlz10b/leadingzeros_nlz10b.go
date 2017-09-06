package nlz10b

const u = 0

var table = [64]uint8{
	32, 20, 19, u, u, 18, u, 7,
	10, 17, u, u, 14, u, 6, u,
	u, 9, u, 16, u, u, 1, 26,
	u, 13, u, u, 24, 5, u, u,
	u, 21, u, 8, 11, u, 15, u,
	u, u, u, 2, 27, 0, 25, u,
	22, u, 12, u, u, 3, 28, u,
	23, u, 4, 29, u, u, 30, 31,
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right.
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x & ^(x >> 16)
	x = x * 0xFD7049FF
	return int(table[x>>26])
}

// LeadingZeros32NoMultiply returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32NoMultiply(x uint32) int {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right.
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x &^ (x >> 16)
	x = (x << 9) - x  // Multiply by 511.
	x = (x << 11) - x // Multiply by 2047.
	x = (x << 14) - x // Multiply by 16383.
	return int(table[x>>26])
}
