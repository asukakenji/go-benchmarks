package onescount

/*
References:
http://www.hackersdelight.org/hdcodetxt/pop.c.txt
http://www.dalkescientific.com/writings/diary/archive/2008/07/03/hakmem_and_other_popcounts.html
*/

import "github.com/asukakenji/go-benchmarks/common/reinterpret"

// OnesCount32Pop2 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2)
func OnesCount32Pop2(x uint32) int {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707 // 6-bit sums.
	return int(x % 63)                // Add 6-bit sums.
}

// OnesCount32Pop2Alt returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2 + alternative)
func OnesCount32Pop2Alt(x uint32) int {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707                  // 6-bit sums.
	return int(((x * 000404040404) >> 26) + (x >> 30)) // Add 6-bit sums.
}

// OnesCount32Pop3 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
func OnesCount32Pop3(x uint32) int {
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

func rotate32(x uint32, n uint) uint32 {
	if n > 63 {
		panic("rotate32, n out of range.")
	}
	return (x << n) | (x >> (32 - n))
}

// OnesCount32Pop5 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
func OnesCount32Pop5(x uint32) int {
	sum := reinterpret.Uint32AsInt32(x)
	for i := 1; i <= 31; i++ {
		x = rotate32(x, 1)
		sum = sum + reinterpret.Uint32AsInt32(x)
	}
	return int(-sum)
}

// OnesCount32Hakmem returns the number of one bits ("population count") in x.
// Source: https://stackoverflow.com/q/8590432/142239
func OnesCount32Hakmem(x uint32) int {
	y := (x >> 1) & 033333333333
	y = x - y - ((y >> 1) & 033333333333)
	return int(((y + (y >> 3)) & 030707070707) % 63)
}

// OnesCount32HakmemUnrolled returns the number of one bits ("population count") in x.
// Source: https://stackoverflow.com/q/8590432/142239
func OnesCount32HakmemUnrolled(x uint32) int {
	y := x - ((x >> 1) & 033333333333) - ((x >> 2) & 011111111111)
	return int(((y + (y >> 3)) & 030707070707) % 63)
}
