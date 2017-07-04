package math

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
import "unsafe"

// BitCountUintNaive returns the number of 1-bits in n.
func BitCountUintNaive(n uint) uint {
	count := uint(0)
	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}

// BitCountUint32Naive returns the number of 1-bits in n.
func BitCountUint32Naive(n uint32) uint {
	count := uint(0)
	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}

// BitCountUint64Naive returns the number of 1-bits in n.
func BitCountUint64Naive(n uint64) uint {
	count := uint(0)
	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}

// BitCountUint32CallGCC returns the number of 1-bits in n.
func BitCountUint32CallGCC(n uint32) uint {
	return uint(C.popcount(C.uint(n)))
}

// BitCountUint64CallGCC returns the number of 1-bits in n.
func BitCountUint64CallGCC(n uint64) uint {
	return uint(C.popcountll(C.ulonglong(n)))
}

const (
	maxUint              = ^uint(0)
	logSizeOfUintInBytes = maxUint>>8&1 + maxUint>>16&1 + maxUint>>32&1
	sizeOfUintInBytes    = 1 << logSizeOfUintInBytes
	sizeOfUintInBits     = sizeOfUintInBytes << 3
)

var (
	bitMask55 uint
	bitMask33 uint
	bitMask0f uint
	bitMask01 uint
	bitShift  uint
)

func init() {
	genBitMask := func(x uint) uint {
		bitMask := x
		shift := uint(8)
		for i := uint(0); i < logSizeOfUintInBytes; i++ {
			bitMask |= (bitMask << shift)
			shift <<= 1
		}
		return bitMask
	}
	bitMask55 = genBitMask(0x55)
	bitMask33 = genBitMask(0x33)
	bitMask0f = genBitMask(0x0f)
	bitMask01 = genBitMask(0x01)
	bitShift = sizeOfUintInBits - 8
}

// BitCountUintGCCImpl returns the number of 1-bits in n.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUintGCCImpl(n uint) uint {
	n = n - ((n >> 1) & bitMask55)
	n = (n & bitMask33) + ((n >> 2) & bitMask33)
	n = (n + (n >> 4)) & bitMask0f
	return (n * bitMask01) >> bitShift
}

// BitCountUintGCCImplSwitch returns the number of 1-bits in n.
func BitCountUintGCCImplSwitch(n uint) uint {
	switch sizeOfUintInBits {
	case 32:
		return BitCountUint32Pop1Alt(uint32(n))
	case 64:
		return BitCountUint64Pop1Alt(uint64(n))
	default:
		panic("uint is neither 32-bit nor 64-bit")
	}
}

// BitCountUint32Pop0 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop0)
func BitCountUint32Pop0(x uint32) uint {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return uint(x)
}

// BitCountUint32Pop1 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1)
// Source: java.lang.Integer#bitCount
// Source: http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/8u40-b25/java/lang/Integer.java#Integer.bitCount%28int%29
func BitCountUint32Pop1(x uint32) uint {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return uint(x & 0x0000003f)
}

// BitCountUint64Pop1 returns the number of 1-bits in x.
// Source: java.lang.Long#bitCount
// Source: http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/8u40-b25/java/lang/Long.java#Long.bitCount%28long%29
func BitCountUint64Pop1(x uint64) uint {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return uint(x & 0x000000000000007f)
}

// BitCountUint32Pop1Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUint32Pop1Alt(x uint32) uint {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return uint((x * 0x01010101) >> 24)
}

// BitCountUint64Pop1Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUint64Pop1Alt(x uint64) uint {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	return uint((x * 0x0101010101010101) >> 56)
}

// BitCountUint32Pop2 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2)
func BitCountUint32Pop2(x uint32) uint {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707 // 6-bit sums.
	return uint(x % 63)               // Add 6-bit sums.
}

// BitCountUint32Pop2Alt returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop2 + alternative)
func BitCountUint32Pop2Alt(x uint32) uint {
	n := (x >> 1) & 033333333333 // Count bits in
	x = x - n                    // each 3-bit
	n = (n >> 1) & 033333333333  // field.
	x = x - n
	x = (x + (x >> 3)) & 030707070707                 // 6-bit sums.
	return uint(((x * 0404040404) >> 26) + (x >> 30)) // Add 6-bit sums.
}

// BitCountUint32Pop3 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop3)
func BitCountUint32Pop3(x uint32) uint {
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

// BitCountUint32Pop4 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func BitCountUint32Pop4(x uint32) uint {
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

func uint32ToInt32(x uint32) int32 {
	return *(*int32)(unsafe.Pointer(&x))
}

func int32ToUint32(x int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&x))
}

// BitCountUint32Pop5 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
func BitCountUint32Pop5(x uint32) uint {
	sum := uint32ToInt32(x)
	for i := 1; i <= 31; i++ {
		x = rotate32(x, 1)
		sum = sum + uint32ToInt32(x)
	}
	return uint(int32ToUint32(-sum))
}

// BitCountUint32Pop5a returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func BitCountUint32Pop5a(x uint32) uint {
	sum := uint32ToInt32(x)
	for x != 0 {
		x = x >> 1
		sum = sum - uint32ToInt32(x)
	}
	return uint(int32ToUint32(sum))
}

var table = [...]uint{
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

// BitCountUint32Pop6 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func BitCountUint32Pop6(x uint32) uint {
	return table[x&0xff] +
		table[(x>>8)&0xff] +
		table[(x>>16)&0xff] +
		table[(x>>24)]
}
