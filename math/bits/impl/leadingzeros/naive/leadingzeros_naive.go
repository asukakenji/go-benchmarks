package naive

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int {
	n := uintSize
	for x != 0 {
		x >>= 1
		n--
	}
	return n
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	if x == 0 {
		return 8
	}
	n := 0
	for x&0x80 == 0 {
		x <<= 1
		n++
	}
	return n
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}
	n := 0
	for x&0x8000 == 0 {
		x <<= 1
		n++
	}
	return n
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}
	n := 0
	for x&0x80000000 == 0 {
		x <<= 1
		n++
	}
	return n
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}
	n := 0
	for x&0x8000000000000000 == 0 {
		x <<= 1
		n++
	}
	return n
}
