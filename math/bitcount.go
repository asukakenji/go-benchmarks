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

// BitCountUint32Java return the number of 1-bits in n.
func BitCountUint32Java(n uint32) uint {
	n = n - ((n >> 1) & uint32(0x55555555))
	n = (n & uint32(0x33333333)) + ((n >> 2) & uint32(0x33333333))
	n = (n + (n >> 4)) & uint32(0x0f0f0f0f)
	n = n + (n >> 8)
	n = n + (n >> 16)
	return uint(n & uint32(0x3f))
}

// BitCountUint64Java return the number of 1-bits in n.
func BitCountUint64Java(n uint64) uint {
	n = n - ((n >> 1) & uint64(0x5555555555555555))
	n = (n & uint64(0x3333333333333333)) + ((n >> 2) & uint64(0x3333333333333333))
	n = (n + (n >> 4)) & uint64(0x0f0f0f0f0f0f0f0f)
	n = n + (n >> 8)
	n = n + (n >> 16)
	n = n + (n >> 32)
	return uint(n & uint64(0x7f))
}

// BitCountUint32CallGCC return the number of 1-bits in n.
func BitCountUint32CallGCC(n uint32) uint {
	return uint(C.popcount(C.uint(n)))
}

// BitCountUint64CallGCC return the number of 1-bits in n.
func BitCountUint64CallGCC(n uint64) uint {
	return uint(C.popcountll(C.ulonglong(n)))
}

// BitCountUint32GCCImpl return the number of 1-bits in n.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUint32GCCImpl(n uint32) uint {
	n = n - ((n >> 1) & uint32(0x55555555))
	n = (n & uint32(0x33333333)) + ((n >> 2) & uint32(0x33333333))
	n = (n + (n >> 4)) & uint32(0x0f0f0f0f)
	return uint((n * uint32(0x01010101)) >> 24)
}

// BitCountUint64GCCImpl return the number of 1-bits in n.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func BitCountUint64GCCImpl(n uint64) uint {
	n = n - ((n >> 1) & uint64(0x5555555555555555))
	n = (n & uint64(0x3333333333333333)) + ((n >> 2) & uint64(0x3333333333333333))
	n = (n + (n >> 4)) & uint64(0x0f0f0f0f0f0f0f0f)
	return uint((n * uint64(0x0101010101010101)) >> 56)
}
