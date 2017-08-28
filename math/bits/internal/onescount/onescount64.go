package onescount

/*
References:
http://www.hackersdelight.org/hdcodetxt/pop.c.txt
http://www.dalkescientific.com/writings/diary/archive/2008/07/03/hakmem_and_other_popcounts.html
*/

// // These functions calls the built-ins in GCC / clang.
//
// static inline int popcount(unsigned int x) {
//     return __builtin_popcount(x);
// }
//
// static inline int popcountl(unsigned long x) {
//     return __builtin_popcountl(x);
// }
//
// static inline int popcountll(unsigned long long x) {
//     return __builtin_popcountll(x);
// }
import "C"

import "github.com/asukakenji/go-benchmarks/common/reinterpret"

// OnesCount64CallGCC returns the number of one bits ("population count") in x.
func OnesCount64CallGCC(x uint64) int {
	return int(C.popcountll(C.ulonglong(x)))
}

/* Does NOT work! */
func onesCount64Pop2(x uint64) int {
	n := (x >> 1) & 01333333333333333333333 // Count bits in
	x = x - n                               // each 3-bit
	n = (n >> 1) & 01333333333333333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 00707070707070707070707 // 6-bit sums.
	return int(x % 63)                           // Add 6-bit sums.
}

/* Does NOT work! */
func onesCount64Pop2Alt(x uint64) int {
	n := (x >> 1) & 01333333333333333333333 // Count bits in
	x = x - n                               // each 3-bit
	n = (n >> 1) & 01333333333333333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 00707070707070707070707             // 6-bit sums.
	return int(((x * 000404040404040404) >> 58) + (x >> 62)) // Add 6-bit sums.
}

// OnesCount64Pop3 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
func OnesCount64Pop3(x uint64) int {
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

func rotate64(x uint64, n uint) uint64 {
	if n > 127 {
		panic("rotate64, n out of range.")
	}
	return (x << n) | (x >> (64 - n))
}

// OnesCount64Pop5 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
func OnesCount64Pop5(x uint64) int {
	sum := reinterpret.Uint64AsInt64(x)
	for i := 1; i <= 63; i++ {
		x = rotate64(x, 1)
		sum = sum + reinterpret.Uint64AsInt64(x)
	}
	return int(-sum)
}

// OnesCount64Hakmem returns the number of one bits ("population count") in x.
// Source: http://www.stmintz.com/ccc/index.php?id=94570
func OnesCount64Hakmem(x uint64) int {
	y := (x >> 1) & 0x7777777777777777
	z := (y >> 1) & 0x7777777777777777
	z = x - y - z - ((z >> 1) & 0x7777777777777777)
	return int(((z + (z >> 4)) & 0x0f0f0f0f0f0f0f0f) % 255)
}
