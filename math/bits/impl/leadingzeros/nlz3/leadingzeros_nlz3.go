package nlz3

import "github.com/asukakenji/go-benchmarks"

// LeadingZeros returns the number of leading zero bits in x; the result is the size of uint in bits for x == 0.
func LeadingZeros(x uint) int {
	_x := int(x)
	n := int(0)
	y := _x
L:
	if _x < 0 {
		return n
	}
	if y == 0 {
		return benchmarks.SizeOfUintInBits - n
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	_x := int8(x)
	n := int8(0)
	y := _x
L:
	if _x < 0 {
		return int(n)
	}
	if y == 0 {
		return int(8 - n)
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	_x := int16(x)
	n := int16(0)
	y := _x
L:
	if _x < 0 {
		return int(n)
	}
	if y == 0 {
		return int(16 - n)
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	_x := int32(x)
	n := int32(0)
	y := _x
L:
	if _x < 0 {
		return int(n)
	}
	if y == 0 {
		return int(32 - n)
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	_x := int64(x)
	n := int64(0)
	y := _x
L:
	if _x < 0 {
		return int(n)
	}
	if y == 0 {
		return int(64 - n)
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}

// LeadingZerosPtr returns the number of leading zero bits in x; the result is the size of uintptr in bits for x == 0.
func LeadingZerosPtr(x uintptr) int {
	_x := int(x)
	n := int(0)
	y := _x
L:
	if _x < 0 {
		return n
	}
	if y == 0 {
		return benchmarks.SizeOfUintInBits - n
	}
	n = n + 1
	x = x << 1
	y = y >> 1
	goto L
}
