package onescount

/*
References:
http://www.hackersdelight.org/hdcodetxt/pop.c.txt
http://www.dalkescientific.com/writings/diary/archive/2008/07/03/hakmem_and_other_popcounts.html
*/

import (
	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// OnesCountNaive returns the number of one bits ("population count") in x.
func OnesCountNaive(x uint) int {
	count := 0
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

// OnesCountPop1Alt returns the number of one bits ("population count") in x.
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
func OnesCountPop1Alt(x uint) int {
	x = x - ((x >> 1) & bitMask55)
	x = (x & bitMask33) + ((x >> 2) & bitMask33)
	x = (x + (x >> 4)) & bitMask0f
	return int((x * bitMask01) >> bitShift)
}

// OnesCountPop4 returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
func OnesCountPop4(x uint) int {
	n := 0
	for x != 0 {
		n = n + 1
		x = x & (x - 1)
	}
	return n
}

// OnesCountPop5a returns the number of one bits ("population count") in x.
// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
func OnesCountPop5a(x uint) int {
	sum := reinterpret.UintAsInt(x)
	for x != 0 {
		x = x >> 1
		sum = sum - reinterpret.UintAsInt(x)
	}
	return sum
}

// OnesCountPop1AltSwitch returns the number of one bits ("population count") in x.
func OnesCountPop1AltSwitch(x uint) int {
	switch common.SizeOfIntInBits {
	case 32:
		return OnesCount32Pop1Alt(uint32(x))
	case 64:
		return OnesCount64Pop1Alt(uint64(x))
	default:
		panic("uint is neither 32-bit nor 64-bit")
	}
}
