package table

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

var nlz8tab = [256]uint8{
	8, 7, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,

	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,

	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

// LeadingZerosConcept returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZerosConcept(x uint) int {
	if x == 0 {
		return uintSize
	}
	n := uint8(0)
	shift := uint(uintSize - 8)
	nWhenAllZeros := uint8(8)
	for {
		n += nlz8tab[(x>>shift)&0xff]
		if n != nWhenAllZeros {
			return int(n)
		}
		nWhenAllZeros += 8
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
