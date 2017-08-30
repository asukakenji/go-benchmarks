package benchmark

import "testing"

// BoolToUint64Func benchmarks a function with the signature:
//     func(bool) uint64
// ID: B-11-1
func BoolToUint64Func(b *testing.B, supplier func() bool, toUint64Func func(bool) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// IntToUint64Func benchmarks a function with the signature:
//     func(int) uint64
// ID: B-11-2
func IntToUint64Func(b *testing.B, supplier func() int, toUint64Func func(int) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int8ToUint64Func benchmarks a function with the signature:
//     func(int8) uint64
// ID: B-11-3
func Int8ToUint64Func(b *testing.B, supplier func() int8, toUint64Func func(int8) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int16ToUint64Func benchmarks a function with the signature:
//     func(int16) uint64
// ID: B-11-4
func Int16ToUint64Func(b *testing.B, supplier func() int16, toUint64Func func(int16) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int32ToUint64Func benchmarks a function with the signature:
//     func(int32) uint64
// ID: B-11-5
func Int32ToUint64Func(b *testing.B, supplier func() int32, toUint64Func func(int32) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int64ToUint64Func benchmarks a function with the signature:
//     func(int64) uint64
// ID: B-11-6
func Int64ToUint64Func(b *testing.B, supplier func() int64, toUint64Func func(int64) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// UintToUint64Func benchmarks a function with the signature:
//     func(uint) uint64
// ID: B-11-7
func UintToUint64Func(b *testing.B, supplier func() uint, toUint64Func func(uint) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint8ToUint64Func benchmarks a function with the signature:
//     func(uint8) uint64
// ID: B-11-8
func Uint8ToUint64Func(b *testing.B, supplier func() uint8, toUint64Func func(uint8) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint16ToUint64Func benchmarks a function with the signature:
//     func(uint16) uint64
// ID: B-11-9
func Uint16ToUint64Func(b *testing.B, supplier func() uint16, toUint64Func func(uint16) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint32ToUint64Func benchmarks a function with the signature:
//     func(uint32) uint64
// ID: B-11-10
func Uint32ToUint64Func(b *testing.B, supplier func() uint32, toUint64Func func(uint32) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint64ToUint64Func benchmarks a function with the signature:
//     func(uint64) uint64
// ID: B-11-11
func Uint64ToUint64Func(b *testing.B, supplier func() uint64, toUint64Func func(uint64) uint64) {
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}
