package naive

import (
	"constraints"

	"github.com/asukakenji/go-benchmarks"
)

// LeadingZerosGeneric returns the number of leading zero bits in x; the result is the size of T in bits for x == 0.
func LeadingZerosGeneric[T constraints.Unsigned](x T) int {
	n := benchmarks.SizeOf[T]()
	for x != 0 {
		x >>= 1
		n--
	}
	return n
}

// LeadingZeros returns the number of leading zero bits in x; the result is the size of uint in bits for x == 0.
func LeadingZeros(x uint) int {
	n := benchmarks.SizeOfUintInBits
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

// LeadingZerosPtr returns the number of leading zero bits in x; the result is the size of uintptr in bits for x == 0.
func LeadingZerosPtr(x uintptr) int {
	n := benchmarks.SizeOf[uintptr]()
	for x != 0 {
		x >>= 1
		n--
	}
	return n
}
