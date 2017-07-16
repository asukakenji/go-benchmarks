package bitcount32

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

// Naive returns the number of 1-bits in x.
func Naive(x uint32) uint {
	count := uint(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// CallGCC returns the number of 1-bits in x.
func CallGCC(x uint32) uint {
	return uint(C.popcount(C.uint(x)))
}

// Pop0 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop0)
func Pop0(x uint32) uint {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return uint(x)
}

// Pop1 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1)
// Source: java.lang.Integer#bitCount
// Source: http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/8u40-b25/java/lang/Integer.java#Integer.bitCount%28int%29
func Pop1(x uint32) uint {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return uint(x & 0x0000003f)
}

// Pop1Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func Pop1Alt(x uint32) uint {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return uint((x * 0x01010101) >> 24)
}

// Pop2 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2)
func Pop2(x uint32) uint {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707 // 6-bit sums.
	return uint(x % 63)               // Add 6-bit sums.
}

// Pop2Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2 + alternative)
func Pop2Alt(x uint32) uint {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707                   // 6-bit sums.
	return uint(((x * 000404040404) >> 26) + (x >> 30)) // Add 6-bit sums.
}

// Pop3 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
func Pop3(x uint32) uint {
	n := (x >> 1) & 0x77777777 // Count bits in
	x = x - n                  // each 4-bit
	n = (n >> 1) & 0x77777777  // field.
	x = x - n
	n = (n >> 1) & 0x77777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f // Get byte sums.
	x = x * 0x01010101              // Add the bytes.
	return uint(x >> 24)
}

// Pop4 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func Pop4(x uint32) uint {
	n := uint(0)
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

// Pop5 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
func Pop5(x uint32) uint {
	sum := reinterpret.Uint32AsInt32(x)
	for i := 1; i <= 31; i++ {
		x = rotate32(x, 1)
		sum = sum + reinterpret.Uint32AsInt32(x)
	}
	return uint(reinterpret.Int32AsUint32(-sum))
}

// Pop5a returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func Pop5a(x uint32) uint {
	sum := reinterpret.Uint32AsInt32(x)
	for x != 0 {
		x = x >> 1
		sum = sum - reinterpret.Uint32AsInt32(x)
	}
	return uint(reinterpret.Int32AsUint32(sum))
}

var byteToBitCountTable = [...]uint{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

// Pop6 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop6)
func Pop6(x uint32) uint {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)]
}

// Hakmem returns the number of 1-bits in x.
// Source: https://stackoverflow.com/q/8590432/142239
func Hakmem(x uint32) uint {
	y := (x >> 1) & 033333333333
	y = x - y - ((y >> 1) & 033333333333)
	return uint(((y + (y >> 3)) & 030707070707) % 63)
}

// HakmemUnrolled returns the number of 1-bits in x.
// Source: https://stackoverflow.com/q/8590432/142239
func HakmemUnrolled(x uint32) uint {
	y := x - ((x >> 1) & 033333333333) - ((x >> 2) & 011111111111)
	return uint(((y + (y >> 3)) & 030707070707) % 63)
}
