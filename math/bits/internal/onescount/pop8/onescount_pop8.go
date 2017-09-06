package pop8

// OnesCount7 returns the number of one bits ("population count") in x.
// x must be less than 0x80 (128), or otherwise the result is undefined.
func OnesCount7(x uint8) int {
	_x := uint32(x)
	_x = _x * 0x02040810 // Make 4 copies, left-adjusted.
	_x = _x & 0x11111111 // Every 4th bit.
	_x = _x * 0x11111111 // Sum the digits (each 0 or 1).
	_x = _x >> 28        // Position the result.
	return int(_x)
}

// OnesCount7Unrolled returns the number of one bits ("population count") in x.
// x must be less than 0x80 (128), or otherwise the result is undefined.
func OnesCount7Unrolled(x uint8) int {
	return int((((uint32(x) * 0x02040810) & 0x11111111) * 0x11111111) >> 28)
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	return OnesCount7(x&0x7f) + int((x>>7)&0x1)
}

// OnesCount8Unrolled returns the number of one bits ("population count") in x.
func OnesCount8Unrolled(x uint8) int {
	return OnesCount7Unrolled(x&0x7f) + int((x>>7)&0x1)
}
