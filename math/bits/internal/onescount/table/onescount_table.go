// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop6)
package table

var byteToBitCountTable = [...]int{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

// OnesCountConcept returns the number of one bits ("population count") in x.
func OnesCountConcept(x uint) int {
	sum := 0
	for x != 0 {
		sum += byteToBitCountTable[x&0xff]
		x >>= 8
	}
	return sum
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	return byteToBitCountTable[x&0xff]
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff]
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)]
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)&0xff] +
		byteToBitCountTable[(x>>32)&0xff] +
		byteToBitCountTable[(x>>40)&0xff] +
		byteToBitCountTable[(x>>48)&0xff] +
		byteToBitCountTable[(x>>56)]
}
