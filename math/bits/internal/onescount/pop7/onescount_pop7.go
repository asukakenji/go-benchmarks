package pop7

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	_x := uint32(x)
	_x = _x * 0x08040201 // Make 4 copies.
	_x = _x >> 3         // So next step hits proper bits.
	_x = _x & 0x11111111 // Every 4th bit.
	_x = _x * 0x11111111 // Sum the digits (each 0 or 1).
	_x = _x >> 28        // Position the result.
	return int(_x)
}

// OnesCount8Unrolled returns the number of one bits ("population count") in x.
func OnesCount8Unrolled(x uint8) int {
	return int(((((uint32(x) * 0x08040201) >> 3) & 0x11111111) * 0x11111111) >> 28)
}
