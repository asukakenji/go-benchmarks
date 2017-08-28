// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop0)
package pop0

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	x = (x & 0x55) + ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	x = (x & 0x0f) + ((x >> 4) & 0x0f)
	return int(x)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	x = (x & 0x5555) + ((x >> 1) & 0x5555)
	x = (x & 0x3333) + ((x >> 2) & 0x3333)
	x = (x & 0x0f0f) + ((x >> 4) & 0x0f0f)
	x = (x & 0x00ff) + ((x >> 8) & 0x00ff)
	return int(x)
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return int(x)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0f0f0f0f0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
	x = (x & 0x00ff00ff00ff00ff) + ((x >> 8) & 0x00ff00ff00ff00ff)
	x = (x & 0x0000ffff0000ffff) + ((x >> 16) & 0x0000ffff0000ffff)
	x = (x & 0x00000000ffffffff) + ((x >> 32) & 0x00000000ffffffff)
	return int(x)
}
