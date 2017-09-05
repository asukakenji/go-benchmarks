// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2)
package pop2

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707 // 6-bit sums.
	return int(x % 63)                // Add 6-bit sums.
}

/* Does NOT work! */
func onesCount64(x uint64) int {
	n := (x >> 1) & 01333333333333333333333 // Count bits in
	x = x - n                               // each 3-bit
	n = (n >> 1) & 01333333333333333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 00707070707070707070707 // 6-bit sums.
	return int(x % 63)                           // Add 6-bit sums.
}
