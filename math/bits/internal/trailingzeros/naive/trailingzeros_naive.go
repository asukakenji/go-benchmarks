package naive

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
func TrailingZeros(x uint) int {
	if x == 0 {
		return uintSize
	}

	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func TrailingZeros8(x uint8) int {
	if x == 0 {
		return 8
	}

	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func TrailingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}

	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}

	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}

	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}
