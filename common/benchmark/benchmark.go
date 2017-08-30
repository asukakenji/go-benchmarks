package benchmark

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
)

var genUintSlice = randomsupplier.NewUintSlice()

// UintSliceGenerator benchmarks a function with the signature:
//     func([]uint, func([]uint))
// ID: B-0-23(7)-19(0-23(7)-19(0-23(7)))
func UintSliceGenerator(b *testing.B, genUintSliceFunc func() []uint, f func([]uint, func([]uint))) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	consumer := func(s []uint) {
		if len(s) != 0 {
			dummy += s[0]
		}
	}
	for i, count := 0, b.N; i < count; i++ {
		f(genUintSliceFunc(), consumer)
	}
	return dummy
}

// UintSliceGeneratorWithRandom benchmarks a function with the signature:
//     func([]uint, func([]uint))
// with random numbers.
// ID: B-0-23(7)-19(0-23(7)-19(0-23(7)))
func UintSliceGeneratorWithRandom(b *testing.B, f func([]uint, func([]uint))) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	consumer := func(s []uint) {
		if len(s) != 0 {
			dummy += s[0]
		}
	}
	genUintSlice.Reset()
	for i, count := 0, b.N; i < count; i++ {
		f(genUintSlice.Next(), consumer)
	}
	return dummy
}
