// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
package pop3

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	n := (x >> 1) & 0x77777777 // Count bits in
	x = x - n                  // each 4-bit
	n = (n >> 1) & 0x77777777  // field.
	x = x - n
	n = (n >> 1) & 0x77777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f // Get byte sums.
	x = x * 0x01010101              // Add the bytes.
	return int(x >> 24)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	n := (x >> 1) & 0x7777777777777777 // Count bits in
	x = x - n                          // each 4-bit
	n = (n >> 1) & 0x7777777777777777  // field.
	x = x - n
	n = (n >> 1) & 0x7777777777777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f // Get byte sums.
	x = x * 0x0101010101010101              // Add the bytes.
	return int(x >> 56)
}
