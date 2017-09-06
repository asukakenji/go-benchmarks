package asm

import "github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/asm/internal/asm"

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	return int(asm.OnesCount64Asm(x))
}
