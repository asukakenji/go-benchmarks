package bitcount64

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
func Naive(x uint64) uint {
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
func CallGCC(x uint64) uint {
	return uint(C.popcountll(C.ulonglong(x)))
}

// Pop0 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop0)
func Pop0(x uint64) uint {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0f0f0f0f0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
	x = (x & 0x00ff00ff00ff00ff) + ((x >> 8) & 0x00ff00ff00ff00ff)
	x = (x & 0x0000ffff0000ffff) + ((x >> 16) & 0x0000ffff0000ffff)
	x = (x & 0x00000000ffffffff) + ((x >> 32) & 0x00000000ffffffff)
	return uint(x)
}

// Pop1 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1)
// Source: java.lang.Long#bitCount
// Source: http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/8u40-b25/java/lang/Long.java#Long.bitCount%28long%29
func Pop1(x uint64) uint {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return uint(x & 0x000000000000007f)
}

// Pop1Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func Pop1Alt(x uint64) uint {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	return uint((x * 0x0101010101010101) >> 56)
}

/* Does NOT work! */
func pop2(x uint64) uint {
	n := (x >> 1) & 01333333333333333333333 // Count bits in
	x = x - n                               // each 3-bit
	n = (n >> 1) & 01333333333333333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 00707070707070707070707 // 6-bit sums.
	return uint(x % 63)                          // Add 6-bit sums.
}

/* Does NOT work! */
func pop2Alt(x uint64) uint {
	n := (x >> 1) & 01333333333333333333333 // Count bits in
	x = x - n                               // each 3-bit
	n = (n >> 1) & 01333333333333333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 00707070707070707070707              // 6-bit sums.
	return uint(((x * 000404040404040404) >> 58) + (x >> 62)) // Add 6-bit sums.
}

// Pop3 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
func Pop3(x uint64) uint {
	n := (x >> 1) & 0x7777777777777777 // Count bits in
	x = x - n                          // each 4-bit
	n = (n >> 1) & 0x7777777777777777  // field.
	x = x - n
	n = (n >> 1) & 0x7777777777777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f // Get byte sums.
	x = x * 0x0101010101010101              // Add the bytes.
	return uint(x >> 56)
}

// Pop4 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func Pop4(x uint64) uint {
	n := uint(0)
	for x != 0 {
		n = n + 1
		x = x & (x - 1)
	}
	return n
}

func rotate64(x uint64, n uint) uint64 {
	if n > 127 {
		panic("rotate64, n out of range.")
	}
	return (x << n) | (x >> (64 - n))
}

// Pop5 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
func Pop5(x uint64) uint {
	sum := reinterpret.Uint64AsInt64(x)
	for i := 1; i <= 63; i++ {
		x = rotate64(x, 1)
		sum = sum + reinterpret.Uint64AsInt64(x)
	}
	return uint(reinterpret.Int64AsUint64(-sum))
}

// Pop5a returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func Pop5a(x uint64) uint {
	sum := reinterpret.Uint64AsInt64(x)
	for x != 0 {
		x = x >> 1
		sum = sum - reinterpret.Uint64AsInt64(x)
	}
	return uint(reinterpret.Int64AsUint64(sum))
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
func Pop6(x uint64) uint {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)&0xff] +
		byteToBitCountTable[(x>>32)&0xff] +
		byteToBitCountTable[(x>>40)&0xff] +
		byteToBitCountTable[(x>>48)&0xff] +
		byteToBitCountTable[(x>>56)]
}

// Hakmem returns the number of 1-bits in x.
// Source: http://www.stmintz.com/ccc/index.php?id=94570
func Hakmem(x uint64) uint {
	y := (x >> 1) & 0x7777777777777777
	z := (y >> 1) & 0x7777777777777777
	z = x - y - z - ((z >> 1) & 0x7777777777777777)
	return uint(((z + (z >> 4)) & 0x0f0f0f0f0f0f0f0f) % 255)
}
