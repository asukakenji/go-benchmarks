package nlz2a

import "github.com/asukakenji/go-benchmarks"

// LeadingZeros returns the number of leading zero bits in x; the result is the size of uint in bits for x == 0.
func LeadingZeros(x uint) int {
	n := uint(benchmarks.SizeOfUintInBits)
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

// LeadingZerosPtr returns the number of leading zero bits in x; the result is the size of uintptr in bits for x == 0.
func LeadingZerosPtr(x uintptr) int {
	n := uintptr(benchmarks.SizeOfUintInBits)
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
