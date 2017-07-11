package benchmark

import "testing"

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

// RandomUintFuncUint benchmarks a function with the signature with random numbers:
//     func(uint) uint
// ID: BR-7-7
func RandomUintFuncUint(b *testing.B, f func(uint) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint.Next())
	}
	return dummy
}

// RandomUintFuncUint32 benchmarks a function with the signature with random numbers:
//     func(uint32) uint
// ID: BR-7-10
func RandomUintFuncUint32(b *testing.B, f func(uint32) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint32.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint32.Next())
	}
	return dummy
}

// RandomUintFuncUint64 benchmarks a function with the signature with random numbers:
//     func(uint64) uint
// ID: BR-7-11
func RandomUintFuncUint64(b *testing.B, f func(uint64) uint) uint {
	dummy := uint(0) // Prevent the call from being optimized out
	genUint64.Reset()
	for i, count := 0, b.N; i < count; i++ {
		dummy += f(genUint64.Next())
	}
	return dummy
}
