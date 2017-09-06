package nlz10

const u = 0

var table = [64]uint8{
	32, 31, u, 16, u, 30, 3, u,
	15, u, u, u, 29, 10, 2, u,
	u, u, 12, 14, 21, u, 19, u,
	u, 28, u, 25, u, 9, 1, u,
	17, u, 4, u, u, u, 11, u,
	13, 22, 20, u, 26, u, u, 18,
	5, u, u, 23, u, 27, u, 6,
	u, 24, 7, u, 8, u, 0, u,
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right.
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x * 0x06EB14F9 // Multiplier is 7*255**3.
	return int(table[x>>26])
}

// LeadingZeros32NoMultiply returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32NoMultiply(x uint32) int {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right.
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = (x << 3) - x // Multiply by 7.
	x = (x << 8) - x // Multiply by 255.
	x = (x << 8) - x // Again.
	x = (x << 8) - x // Again.
	return int(table[x>>26])
}
