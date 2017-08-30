package benchmark

import "testing"

// CalibrateUintSliceGenerator calibrates the UintSliceGenerator function.
// It should be called once before all benchmarks if any benchmark depends on
// UintSliceGenerator.
//
// ID: BC-0-23(7)-19(0-23(7)-19(0-23(7)))
func CalibrateUintSliceGenerator(b *testing.B, genUintSliceFunc func() []uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	consumer := func(s []uint) {
		if len(s) != 0 {
			dummy += s[0]
		}
	}
	for i, count := 0, b.N; i < count; i++ {
		s := genUintSliceFunc()
		consumer(s)
	}
	return dummy
}

// CalibrateUintSliceGeneratorWithRandom calibrates the UintSliceGeneratorWithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// UintSliceGeneratorWithRandom.
//
// ID: BRC-0-23(7)-19(0-23(7)-19(0-23(7)))
func CalibrateUintSliceGeneratorWithRandom(b *testing.B) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	consumer := func(s []uint) {
		if len(s) != 0 {
			dummy += s[0]
		}
	}
	genUintSlice.Reset()
	for i, count := 0, b.N; i < count; i++ {
		s := genUintSlice.Next()
		consumer(s)
	}
	return dummy
}
