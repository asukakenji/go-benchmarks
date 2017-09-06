package forward

import "github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/pop1a"

const (
	uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	is32Bit  = (^uint(0) >> 32 & 1) == 0
)

var (
	onesCount func(uint) int
)

func init() {
	if is32Bit {
		onesCount = onesCount32
	} else {
		onesCount = onesCount64
	}
}

func onesCount32(x uint) int {
	return pop1a.OnesCount32(uint32(x))
}

func onesCount64(x uint) int {
	return pop1a.OnesCount64(uint64(x))
}

// OnesCountIfConstBool returns the number of one bits ("population count") in x.
func OnesCountIfConstBool(x uint) int {
	if is32Bit {
		return pop1a.OnesCount32(uint32(x))
	}
	return pop1a.OnesCount64(uint64(x))
}

// OnesCountIfConstUint returns the number of one bits ("population count") in x.
func OnesCountIfConstUint(x uint) int {
	if uintSize == 32 {
		return pop1a.OnesCount32(uint32(x))
	}
	return pop1a.OnesCount64(uint64(x))
}

// OnesCountSwitchConstUint returns the number of one bits ("population count") in x.
func OnesCountSwitchConstUint(x uint) int {
	switch uintSize {
	case 32:
		return pop1a.OnesCount32(uint32(x))
	case 64:
		return pop1a.OnesCount64(uint64(x))
	default:
		panic("uint is neither 32-bit nor 64-bit")
	}
}

// OnesCountFuncPointer returns the number of one bits ("population count") in x.
func OnesCountFuncPointer(x uint) int {
	return onesCount(x)
}
