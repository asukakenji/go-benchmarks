package nlz1

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	if x == 0 {
		return 8
	}

	n := 0
	if x <= 0x0f {
		n = n + 4
		x = x << 4
	}
	if x <= 0x3f {
		n = n + 2
		x = x << 2
	}
	if x <= 0x7f {
		n = n + 1
	}
	return n
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}

	n := 0
	if x <= 0x00ff {
		n = n + 8
		x = x << 8
	}
	if x <= 0x0fff {
		n = n + 4
		x = x << 4
	}
	if x <= 0x3fff {
		n = n + 2
		x = x << 2
	}
	if x <= 0x7fff {
		n = n + 1
	}
	return n
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}

	n := 0
	if x <= 0x0000ffff {
		n = n + 16
		x = x << 16
	}
	if x <= 0x00ffffff {
		n = n + 8
		x = x << 8
	}
	if x <= 0x0fffffff {
		n = n + 4
		x = x << 4
	}
	if x <= 0x3fffffff {
		n = n + 2
		x = x << 2
	}
	if x <= 0x7fffffff {
		n = n + 1
	}
	return n
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}

	n := 0
	if x <= 0x00000000ffffffff {
		n = n + 32
		x = x << 32
	}
	if x <= 0x0000ffffffffffff {
		n = n + 16
		x = x << 16
	}
	if x <= 0x00ffffffffffffff {
		n = n + 8
		x = x << 8
	}
	if x <= 0x0fffffffffffffff {
		n = n + 4
		x = x << 4
	}
	if x <= 0x3fffffffffffffff {
		n = n + 2
		x = x << 2
	}
	if x <= 0x7fffffffffffffff {
		n = n + 1
	}
	return n
}
