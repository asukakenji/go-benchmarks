// Source: https://stackoverflow.com/q/8590432/142239
// Source: http://www.stmintz.com/ccc/index.php?id=94570
package hakmem

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	y := (x >> 1) & 033333333333
	y = x - y - ((y >> 1) & 033333333333)
	return int(((y + (y >> 3)) & 030707070707) % 63)
}

// OnesCount32Unrolled returns the number of one bits ("population count") in x.
func OnesCount32Unrolled(x uint32) int {
	y := x - ((x >> 1) & 033333333333) - ((x >> 2) & 011111111111)
	return int(((y + (y >> 3)) & 030707070707) % 63)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	y := (x >> 1) & 0x7777777777777777
	z := (y >> 1) & 0x7777777777777777
	z = x - y - z - ((z >> 1) & 0x7777777777777777)
	return int(((z + (z >> 4)) & 0x0f0f0f0f0f0f0f0f) % 255)
}
