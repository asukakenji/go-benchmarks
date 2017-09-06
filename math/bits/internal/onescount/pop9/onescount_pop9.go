package pop9

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	y := uint64(x) * 0x0002000400080010
	y = y & 0x1111111111111111
	y = y * 0x1111111111111111
	y = y >> 60
	return int(y)
}

// OnesCount8Unrolled returns the number of one bits ("population count") in x.
func OnesCount8Unrolled(x uint8) int {
	return int((((uint64(x) * 0x0002000400080010) & 0x1111111111111111) * 0x1111111111111111) >> 60)
}

// OnesCount15 returns the number of one bits ("population count") in x.
// x must be less than 0x8000 (32768), or otherwise the result is undefined.
func OnesCount15(x uint16) int {
	y := uint64(x) * 0x0002000400080010
	y = y & 0x1111111111111111
	y = y * 0x1111111111111111
	y = y >> 60
	return int(y)
}

// OnesCount15Unrolled returns the number of one bits ("population count") in x.
// x must be less than 0x8000 (32768), or otherwise the result is undefined.
func OnesCount15Unrolled(x uint16) int {
	return int((((uint64(x) * 0x0002000400080010) & 0x1111111111111111) * 0x1111111111111111) >> 60)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	return OnesCount15(x&0x7fff) + int((x>>15)&0x1)
}

// OnesCount16Unrolled returns the number of one bits ("population count") in x.
func OnesCount16Unrolled(x uint16) int {
	return OnesCount15Unrolled(x&0x7fff) + int((x>>15)&0x1)
}
