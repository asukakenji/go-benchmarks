package nlz1a

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	if int8(x) <= 0 {
		return int((^x >> 4) & 8)
	}
	n := uint8(1)
	if (x >> 4) == 0 {
		n = n + 4
		x = x << 4
	}
	if (x >> 6) == 0 {
		n = n + 2
		x = x << 2
	}
	n = n - (x >> 7)
	return int(n)
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	if int16(x) <= 0 {
		return int((^x >> 11) & 16)
	}
	n := uint16(1)
	if (x >> 8) == 0 {
		n = n + 8
		x = x << 8
	}
	if (x >> 12) == 0 {
		n = n + 4
		x = x << 4
	}
	if (x >> 14) == 0 {
		n = n + 2
		x = x << 2
	}
	n = n - (x >> 15)
	return int(n)
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	if int32(x) <= 0 {
		return int((^x >> 26) & 32)
	}
	n := uint32(1)
	if (x >> 16) == 0 {
		n = n + 16
		x = x << 16
	}
	if (x >> 24) == 0 {
		n = n + 8
		x = x << 8
	}
	if (x >> 28) == 0 {
		n = n + 4
		x = x << 4
	}
	if (x >> 30) == 0 {
		n = n + 2
		x = x << 2
	}
	n = n - (x >> 31)
	return int(n)
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	if int64(x) <= 0 {
		return int((^x >> 57) & 64)
	}
	n := uint64(1)
	if (x >> 32) == 0 {
		n = n + 32
		x = x << 32
	}
	if (x >> 48) == 0 {
		n = n + 16
		x = x << 16
	}
	if (x >> 56) == 0 {
		n = n + 8
		x = x << 8
	}
	if (x >> 60) == 0 {
		n = n + 4
		x = x << 4
	}
	if (x >> 62) == 0 {
		n = n + 2
		x = x << 2
	}
	n = n - (x >> 63)
	return int(n)
}
