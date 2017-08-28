package onescount

/*
References:
http://www.hackersdelight.org/hdcodetxt/pop.c.txt
http://www.dalkescientific.com/writings/diary/archive/2008/07/03/hakmem_and_other_popcounts.html
*/

import (
	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/pop1a"
)

// OnesCountPop1AltSwitch returns the number of one bits ("population count") in x.
func OnesCountPop1AltSwitch(x uint) int {
	switch common.SizeOfIntInBits {
	case 32:
		return pop1a.OnesCount32(uint32(x))
	case 64:
		return pop1a.OnesCount64(uint64(x))
	default:
		panic("uint is neither 32-bit nor 64-bit")
	}
}
