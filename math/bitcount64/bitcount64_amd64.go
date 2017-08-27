package bitcount64

import (
	"github.com/asukakenji/go-benchmarks/math/bitcount64/internal/asm"
)

func Asm(x uint64) uint {
	return uint(OnesCount64Asm(x))
}

func OnesCount64Asm(x uint64) int {
	return int(asm.OnesCount64Asm(x))
}
