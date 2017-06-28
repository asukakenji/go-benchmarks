package math

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

// BitCountUInt32Naive return the number of 1-bits in n.
func BitCountUInt32Naive(n uint32) int {
	count := 0
	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}

// BitCountUInt64Naive return the number of 1-bits in n.
func BitCountUInt64Naive(n uint64) int {
	count := 0
	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}

// BitCountUInt32Java return the number of 1-bits in n.
func BitCountUInt32Java(n uint32) int {
	n = n - ((n >> 1) & uint32(0x55555555))
	n = (n & uint32(0x33333333)) + ((n >> 2) & uint32(0x33333333))
	n = (n + (n >> 4)) & uint32(0x0f0f0f0f)
	n = n + (n >> 8)
	n = n + (n >> 16)
	return int(n & uint32(0x3f))
}

// BitCountUInt64Java return the number of 1-bits in n.
func BitCountUInt64Java(n uint64) int {
	n = n - ((n >> 1) & uint64(0x5555555555555555))
	n = (n & uint64(0x3333333333333333)) + ((n >> 2) & uint64(0x3333333333333333))
	n = (n + (n >> 4)) & uint64(0x0f0f0f0f0f0f0f0f)
	n = n + (n >> 8)
	n = n + (n >> 16)
	n = n + (n >> 32)
	return int(n & uint64(0x7f))
}

// BitCountUInt32CallGCC return the number of 1-bits in n.
func BitCountUInt32CallGCC(n uint32) int {
	return int(C.popcount(C.uint(n)))
}

// BitCountUInt64CallGCC return the number of 1-bits in n.
func BitCountUInt64CallGCC(n uint64) int {
	return int(C.popcountll(C.ulonglong(n)))
}

// BitCountUInt32GCCImpl return the number of 1-bits in n.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUInt32GCCImpl(n uint32) int {
	n = n - ((n >> 1) & uint32(0x55555555))
	n = (n & uint32(0x33333333)) + ((n >> 2) & uint32(0x33333333))
	n = (n + (n >> 4)) & uint32(0x0f0f0f0f)
	return int((n * uint32(0x01010101)) >> 24)
}

// BitCountUInt64GCCImpl return the number of 1-bits in n.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUInt64GCCImpl(n uint64) int {
	n = n - ((n >> 1) & uint64(0x5555555555555555))
	n = (n & uint64(0x3333333333333333)) + ((n >> 2) & uint64(0x3333333333333333))
	n = (n + (n >> 4)) & uint64(0x0f0f0f0f0f0f0f0f)
	return int((n * uint64(0x0101010101010101)) >> 56)
}
