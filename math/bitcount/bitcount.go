package bitcount

/*
References:
http://www.hackersdelight.org/hdcodetxt/pop.c.txt
http://www.dalkescientific.com/writings/diary/archive/2008/07/03/hakmem_and_other_popcounts.html
*/

import (
	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
	"github.com/asukakenji/go-benchmarks/math/bitcount32"
	"github.com/asukakenji/go-benchmarks/math/bitcount64"
)

// Naive returns the number of 1-bits in x.
func Naive(x uint) uint {
	count := uint(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

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
		for i := uint(0); i < common.LogSizeOfIntInBytes; i++ {
			bitMask |= (bitMask << shift)
			shift <<= 1
		}
		return bitMask
	}
	bitMask55 = genBitMask(0x55)
	bitMask33 = genBitMask(0x33)
	bitMask0f = genBitMask(0x0f)
	bitMask01 = genBitMask(0x01)
	bitShift = common.SizeOfIntInBits - 8
}

// Pop1Alt returns the number of 1-bits in x.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func Pop1Alt(x uint) uint {
	x = x - ((x >> 1) & bitMask55)
	x = (x & bitMask33) + ((x >> 2) & bitMask33)
	x = (x + (x >> 4)) & bitMask0f
	return (x * bitMask01) >> bitShift
}

// Pop4 returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func Pop4(x uint) uint {
	n := uint(0)
	for x != 0 {
		n = n + 1
		x = x & (x - 1)
	}
	return n
}

// Pop5a returns the number of 1-bits in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func Pop5a(x uint) uint {
	sum := reinterpret.UintAsInt(x)
	for x != 0 {
		x = x >> 1
		sum = sum - reinterpret.UintAsInt(x)
	}
	return uint(reinterpret.IntAsUint(sum))
}

// Pop1AltSwitch returns the number of 1-bits in x.
func Pop1AltSwitch(x uint) uint {
	switch common.SizeOfIntInBits {
	case 32:
		return bitcount32.Pop1Alt(uint32(x))
	case 64:
		return bitcount64.Pop1Alt(uint64(x))
	default:
		panic("uint is neither 32-bit nor 64-bit")
	}
}
