package table

import "github.com/asukakenji/go-benchmarks"

// LeadingZeros returns the number of leading zero bits in x; the result is the size of uint in bits for x == 0.
func LeadingZeros(x uint) int {
	if x == 0 {
		return benchmarks.SizeOfUintInBits
	}
	n := uint8(0)
	shift := benchmarks.SizeOfUintInBits - 8
	nForZeroX := uint8(8)
	for {
		n += nlz8tab[(x>>shift)&0xff]
		if n != nForZeroX {
			return int(n)
		}
		nForZeroX += 8
		shift -= 8
	}
}

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int {
	return int(nlz8tab[x])
}

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int {
	n := nlz8tab[x>>8]
	if n != 8 {
		return int(n)
	}
	return int(n + nlz8tab[x&0xff])
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	n := nlz8tab[x>>24]
	if n != 8 {
		return int(n)
	}
	n += nlz8tab[(x>>16)&0xff]
	if n != 16 {
		return int(n)
	}
	n += nlz8tab[(x>>8)&0xff]
	if n != 24 {
		return int(n)
	}
	return int(n + nlz8tab[x&0xff])
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int {
	n := nlz8tab[x>>56]
	if n != 8 {
		return int(n)
	}
	n += nlz8tab[(x>>48)&0xff]
	if n != 16 {
		return int(n)
	}
	n += nlz8tab[(x>>40)&0xff]
	if n != 24 {
		return int(n)
	}
	n += nlz8tab[(x>>32)&0xff]
	if n != 32 {
		return int(n)
	}
	n += nlz8tab[(x>>24)&0xff]
	if n != 40 {
		return int(n)
	}
	n += nlz8tab[(x>>16)&0xff]
	if n != 48 {
		return int(n)
	}
	n += nlz8tab[(x>>8)&0xff]
	if n != 56 {
		return int(n)
	}
	return int(n + nlz8tab[x&0xff])
}

// LeadingZerosPtr returns the number of leading zero bits in x; the result is the size of uintptr in bits for x == 0.
func LeadingZerosPtr(x uintptr) int {
	if x == 0 {
		return benchmarks.SizeOf[uintptr]()
	}
	n := uint8(0)
	shift := benchmarks.SizeOf[uintptr]() - 8
	nForZeroX := uint8(8)
	for {
		n += nlz8tab[(x>>shift)&0xff]
		if n != nForZeroX {
			return int(n)
		}
		nForZeroX += 8
		shift -= 8
	}
}
