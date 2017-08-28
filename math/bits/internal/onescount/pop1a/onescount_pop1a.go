// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop1 + alternative)
// Source: https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L840-L859
// Source: https://gcc.gnu.org/bugzilla/show_bug.cgi?id=36041#c8
package pop1a

import "github.com/asukakenji/go-benchmarks/common"

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

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	x = x - ((x >> 1) & bitMask55)
	x = (x & bitMask33) + ((x >> 2) & bitMask33)
	x = (x + (x >> 4)) & bitMask0f
	return int((x * bitMask01) >> bitShift)
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	x = x - ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	x = (x + (x >> 4)) & 0x0f
	return int(x)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	x = x - ((x >> 1) & 0x5555)
	x = (x & 0x3333) + ((x >> 2) & 0x3333)
	x = (x + (x >> 4)) & 0x0f0f
	return int((x * 0x0101) >> 8)
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return int((x * 0x01010101) >> 24)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	return int((x * 0x0101010101010101) >> 56)
}
