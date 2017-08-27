package onescount

import "github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/internal/asm"

// OnesCount64Asm returns the number of one bits ("population count") in x.
func OnesCount64Asm(x uint64) int {
	return int(asm.OnesCount64Asm(x))
}
