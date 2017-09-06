package nlz2a

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int {
	n := uint(uintSize)
	c := n >> 1
	for {
		y := x >> c
		if y != 0 {
			n = n - c
			x = y
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return int(n - x)
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	n := uint8(8)
	c := uint8(4)
	for {
		y := x >> c
		if y != 0 {
			n = n - c
			x = y
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return int(n - x)
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	n := uint16(16)
	c := uint16(8)
	for {
		y := x >> c
		if y != 0 {
			n = n - c
			x = y
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return int(n - x)
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	n := uint32(32)
	c := uint32(16)
	for {
		y := x >> c
		if y != 0 {
			n = n - c
			x = y
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return int(n - x)
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	n := uint64(64)
	c := uint64(32)
	for {
		y := x >> c
		if y != 0 {
			n = n - c
			x = y
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return int(n - x)
}
