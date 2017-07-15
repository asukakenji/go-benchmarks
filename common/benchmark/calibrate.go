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

// CalibrateUintFuncUint calibrates the UintFuncUint function.
// It should be called once before all benchmarks if any benchmark depends on
// UintFuncUint.
//
// ID: BC-7-7
func CalibrateUintFuncUint(b *testing.B, genUintFunc func() uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUintFunc()
	}
	return dummy
}

// CalibrateUintFuncUint32 calibrates the UintFuncUint32 function.
// It should be called once before all benchmarks if any benchmark depends on
// UintFuncUint32.
//
// ID: BC-7-10
func CalibrateUintFuncUint32(b *testing.B, genUint32Func func() uint32) uint32 {
	dummy := uint32(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint32Func()
	}
	return dummy
}

// CalibrateUintFuncUint64 calibrates the UintFuncUint64 function.
// It should be called once before all benchmarks if any benchmark depends on
// UintFuncUint64.
//
// ID: BC-7-11
func CalibrateUintFuncUint64(b *testing.B, genUint64Func func() uint64) uint64 {
	dummy := uint64(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint64Func()
	}
	return dummy
}

// CalibrateRandomUintSliceGenerator calibrates the RandomUintSliceGenerator function.
// It should be called once before all benchmarks if any benchmark depends on
// RandomUintSliceGenerator.
//
// ID: BRC-0-23(7)-19(0-23(7)-19(0-23(7)))
func CalibrateRandomUintSliceGenerator(b *testing.B, genUintSliceFunc func() []uint) uint {
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

// CalibrateRandomUintFuncUint calibrates the RandomUintFuncUint function.
// It should be called once before all benchmarks if any benchmark depends on
// RandomUintFuncUint.
//
// ID: BRC-7-7
func CalibrateRandomUintFuncUint(b *testing.B) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint.Next()
	}
	return dummy
}

// CalibrateRandomUintFuncUint32 calibrates the RandomUintFuncUint32 function.
// It should be called once before all benchmarks if any benchmark depends on
// RandomUintFuncUint32.
//
// ID: BRC-7-10
func CalibrateRandomUintFuncUint32(b *testing.B) uint32 {
	dummy := uint32(0) // Prevent the call from being optimized out
	genUint32.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint32.Next()
	}
	return dummy
}

// CalibrateRandomUintFuncUint64 calibrates the RandomUintFuncUint64 function.
// It should be called once before all benchmarks if any benchmark depends on
// RandomUintFuncUint64.
//
// ID: BRC-7-11
func CalibrateRandomUintFuncUint64(b *testing.B) uint64 {
	dummy := uint64(0) // Prevent the call from being optimized out
	genUint64.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint64.Next()
	}
	return dummy
}
