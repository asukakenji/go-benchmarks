package benchmark

import "testing"

// BoolToIntFunc benchmarks a function with the signature:
//     func(bool) int
// ID: B-2-1
func BoolToIntFunc(b *testing.B, supplier func() bool, toIntFunc func(bool) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// IntToIntFunc benchmarks a function with the signature:
//     func(int) int
// ID: B-2-2
func IntToIntFunc(b *testing.B, supplier func() int, toIntFunc func(int) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int8ToIntFunc benchmarks a function with the signature:
//     func(int8) int
// ID: B-2-3
func Int8ToIntFunc(b *testing.B, supplier func() int8, toIntFunc func(int8) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int16ToIntFunc benchmarks a function with the signature:
//     func(int16) int
// ID: B-2-4
func Int16ToIntFunc(b *testing.B, supplier func() int16, toIntFunc func(int16) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int32ToIntFunc benchmarks a function with the signature:
//     func(int32) int
// ID: B-2-5
func Int32ToIntFunc(b *testing.B, supplier func() int32, toIntFunc func(int32) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int64ToIntFunc benchmarks a function with the signature:
//     func(int64) int
// ID: B-2-6
func Int64ToIntFunc(b *testing.B, supplier func() int64, toIntFunc func(int64) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// UintToIntFunc benchmarks a function with the signature:
//     func(uint) int
// ID: B-2-7
func UintToIntFunc(b *testing.B, supplier func() uint, toIntFunc func(uint) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint8ToIntFunc benchmarks a function with the signature:
//     func(uint8) int
// ID: B-2-8
func Uint8ToIntFunc(b *testing.B, supplier func() uint8, toIntFunc func(uint8) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint16ToIntFunc benchmarks a function with the signature:
//     func(uint16) int
// ID: B-2-9
func Uint16ToIntFunc(b *testing.B, supplier func() uint16, toIntFunc func(uint16) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint32ToIntFunc benchmarks a function with the signature:
//     func(uint32) int
// ID: B-2-10
func Uint32ToIntFunc(b *testing.B, supplier func() uint32, toIntFunc func(uint32) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint64ToIntFunc benchmarks a function with the signature:
//     func(uint64) int
// ID: B-2-11
func Uint64ToIntFunc(b *testing.B, supplier func() uint64, toIntFunc func(uint64) int) {
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}
