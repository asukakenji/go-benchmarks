package benchmark

import "testing"

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

// BoolFuncInt benchmarks a function with the signature:
//     func(int) bool
// ID: B-1-2
func BoolFuncInt(b *testing.B, genIntFunc func() int, f func(int) bool) int {
	dummy := 0 // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		if f(genIntFunc()) {
			dummy++
		} else {
			dummy--
		}
	}
	return dummy
}

// UintFuncUint benchmarks a function with the signature:
//     func(uint) uint
// ID: B-7-7
func UintFuncUint(b *testing.B, genUintFunc func() uint, f func(uint) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUintFunc())
	}
	return dummy
}

// UintFuncUint32 benchmarks a function with the signature:
//     func(uint32) uint
// ID: B-7-10
func UintFuncUint32(b *testing.B, genUint32Func func() uint32, f func(uint32) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint32Func())
	}
	return dummy
}

// UintFuncUint64 benchmarks a function with the signature:
//     func(uint64) uint
// ID: B-7-11
func UintFuncUint64(b *testing.B, genUint64Func func() uint64, f func(uint64) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint64Func())
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

// BoolFuncIntWithRandom benchmarks a function with the signature:
//     func(int) bool
// with random numbers.
// ID: BR-1-2
func BoolFuncIntWithRandom(b *testing.B, f func(int) bool) int {
	dummy := 0 // Prevent the call from being optimized out
	genInt.Reset()
	for i, count := 0, b.N; i < count; i++ {
		if f(genInt.Next()) {
			dummy++
		} else {
			dummy--
		}
	}
	return dummy
}

// IntFuncUintWithRandom benchmarks a function with the signature:
//     func(uint) int
// with random numbers.
// ID: BR-2-7
func IntFuncUintWithRandom(b *testing.B, f func(uint) int) int {
	dummy := 0 // Prevent the call from being optimized out
	genUint.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint.Next())
	}
	return dummy
}

// IntFuncUint8WithRandom benchmarks a function with the signature:
//     func(uint8) int
// with random numbers.
// ID: BR-2-8
func IntFuncUint8WithRandom(b *testing.B, f func(uint8) int) int {
	dummy := 0 // Prevent the call from being optimized out
	genUint8.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint8.Next())
	}
	return dummy
}

// IntFuncUint16WithRandom benchmarks a function with the signature:
//     func(uint16) int
// with random numbers.
// ID: BR-2-9
func IntFuncUint16WithRandom(b *testing.B, f func(uint16) int) int {
	dummy := 0 // Prevent the call from being optimized out
	genUint16.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint16.Next())
	}
	return dummy
}

// IntFuncUint32WithRandom benchmarks a function with the signature:
//     func(uint32) int
// with random numbers.
// ID: BR-2-10
func IntFuncUint32WithRandom(b *testing.B, f func(uint32) int) int {
	dummy := 0 // Prevent the call from being optimized out
	genUint32.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint32.Next())
	}
	return dummy
}

// IntFuncUint64WithRandom benchmarks a function with the signature:
//     func(uint64) int
// with random numbers.
// ID: BR-2-11
func IntFuncUint64WithRandom(b *testing.B, f func(uint64) int) int {
	dummy := 0 // Prevent the call from being optimized out
	genUint64.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint64.Next())
	}
	return dummy
}

// UintFuncUintWithRandom benchmarks a function with the signature:
//     func(uint) uint
// with random numbers.
// ID: BR-7-7
func UintFuncUintWithRandom(b *testing.B, f func(uint) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint.Next())
	}
	return dummy
}

// UintFuncUint8WithRandom benchmarks a function with the signature:
//     func(uint8) uint
// with random numbers.
// ID: BR-7-8
func UintFuncUint8WithRandom(b *testing.B, f func(uint8) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint8.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint8.Next())
	}
	return dummy
}

// UintFuncUint16WithRandom benchmarks a function with the signature:
//     func(uint16) uint
// with random numbers.
// ID: BR-7-9
func UintFuncUint16WithRandom(b *testing.B, f func(uint16) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint16.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint16.Next())
	}
	return dummy
}

// UintFuncUint32WithRandom benchmarks a function with the signature:
//     func(uint32) uint
// with random numbers.
// ID: BR-7-10
func UintFuncUint32WithRandom(b *testing.B, f func(uint32) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint32.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint32.Next())
	}
	return dummy
}

// UintFuncUint64WithRandom benchmarks a function with the signature:
//     func(uint64) uint
// with random numbers.
// ID: BR-7-11
func UintFuncUint64WithRandom(b *testing.B, f func(uint64) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint64.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint64.Next())
	}
	return dummy
}
