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

// CalibrateBoolFuncInt calibrates the BoolFuncInt function.
// It should be called once before all benchmarks if any benchmark depends on
// BoolFuncInt.
//
// ID: BC-1-2
func CalibrateBoolFuncInt(b *testing.B, genIntFunc func() int) int {
	dummy := 0 // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += genIntFunc()
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

// CalibrateBoolFuncIntWithRandom calibrates the BoolFuncIntWithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// BoolFuncIntWithRandom.
//
// ID: BRC-1-2
func CalibrateBoolFuncIntWithRandom(b *testing.B) int {
	dummy := 0 // Prevent the call from being optimized out
	genInt.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genInt.Next()
	}
	return dummy
}

// CalibrateAnyFuncUintWithRandom calibrates the XxxFuncUintWithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// XxxFuncUintWithRandom, where Xxx is a type name.
//
// ID: BRC-?-7
func CalibrateAnyFuncUintWithRandom(b *testing.B) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint.Next()
	}
	return dummy
}

// CalibrateAnyFuncUint8WithRandom calibrates the XxxFuncUint8WithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// XxxFuncUint8WithRandom, where Xxx is a type name.
//
// ID: BRC-?-8
func CalibrateAnyFuncUint8WithRandom(b *testing.B) uint8 {
	dummy := uint8(0) // Prevent the call from being optimized out
	genUint8.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint8.Next()
	}
	return dummy
}

// CalibrateAnyFuncUint16WithRandom calibrates the XxxFuncUint16WithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// XxxFuncUint16WithRandom, where Xxx is a type name.
//
// ID: BRC-?-10
func CalibrateAnyFuncUint16WithRandom(b *testing.B) uint16 {
	dummy := uint16(0) // Prevent the call from being optimized out
	genUint16.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint16.Next()
	}
	return dummy
}

// CalibrateAnyFuncUint32WithRandom calibrates the XxxFuncUint32WithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// XxxFuncUint32WithRandom, where Xxx is a type name.
//
// ID: BRC-?-10
func CalibrateAnyFuncUint32WithRandom(b *testing.B) uint32 {
	dummy := uint32(0) // Prevent the call from being optimized out
	genUint32.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint32.Next()
	}
	return dummy
}

// CalibrateAnyFuncUint64WithRandom calibrates the XxxFuncUint64WithRandom function.
// It should be called once before all benchmarks if any benchmark depends on
// XxxFuncUint64WithRandom, where Xxx is a type name.
//
// ID: BRC-?-11
func CalibrateAnyFuncUint64WithRandom(b *testing.B) uint64 {
	dummy := uint64(0) // Prevent the call from being optimized out
	genUint64.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += genUint64.Next()
	}
	return dummy
}
