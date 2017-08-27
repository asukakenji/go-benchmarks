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

// OnesCount32Naive returns the number of one bits ("population count") in x.
func OnesCount32Naive(x uint32) int {
	count := int(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// OnesCount32CallGCC returns the number of one bits ("population count") in x.
func OnesCount32CallGCC(x uint32) int {
	return int(C.popcount(C.uint(x)))
}

// OnesCount32Pop0 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop0)
func OnesCount32Pop0(x uint32) int {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return int(x)
}

// OnesCount32Pop1 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1)
// Source: java.lang.Integer#bitCount
// Source: http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/8u40-b25/java/lang/Integer.java#Integer.bitCount%28int%29
func OnesCount32Pop1(x uint32) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return int(x & 0x0000003f)
}

// OnesCount32Pop1Alt returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func OnesCount32Pop1Alt(x uint32) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return int((x * 0x01010101) >> 24)
}

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

// OnesCount32Pop4 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func OnesCount32Pop4(x uint32) int {
	n := 0
	for x != 0 {
		n = n + 1
		x = x & (x - 1)
	}
	return n
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

// OnesCount32Pop5a returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func OnesCount32Pop5a(x uint32) int {
	sum := reinterpret.Uint32AsInt32(x)
	for x != 0 {
		x = x >> 1
		sum = sum - reinterpret.Uint32AsInt32(x)
	}
	return int(sum)
}

// OnesCount32Pop6 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop6)
func OnesCount32Pop6(x uint32) int {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)]
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
