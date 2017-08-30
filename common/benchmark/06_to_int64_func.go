package benchmark

import "testing"

// BoolToInt64Func benchmarks a function with the signature:
//     func(bool) int64
// ID: B-6-1
func BoolToInt64Func(b *testing.B, supplier func() bool, toInt64Func func(bool) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// IntToInt64Func benchmarks a function with the signature:
//     func(int) int64
// ID: B-6-2
func IntToInt64Func(b *testing.B, supplier func() int, toInt64Func func(int) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int8ToInt64Func benchmarks a function with the signature:
//     func(int8) int64
// ID: B-6-3
func Int8ToInt64Func(b *testing.B, supplier func() int8, toInt64Func func(int8) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int16ToInt64Func benchmarks a function with the signature:
//     func(int16) int64
// ID: B-6-4
func Int16ToInt64Func(b *testing.B, supplier func() int16, toInt64Func func(int16) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int32ToInt64Func benchmarks a function with the signature:
//     func(int32) int64
// ID: B-6-5
func Int32ToInt64Func(b *testing.B, supplier func() int32, toInt64Func func(int32) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int64ToInt64Func benchmarks a function with the signature:
//     func(int64) int64
// ID: B-6-6
func Int64ToInt64Func(b *testing.B, supplier func() int64, toInt64Func func(int64) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// UintToInt64Func benchmarks a function with the signature:
//     func(uint) int64
// ID: B-6-7
func UintToInt64Func(b *testing.B, supplier func() uint, toInt64Func func(uint) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint8ToInt64Func benchmarks a function with the signature:
//     func(uint8) int64
// ID: B-6-8
func Uint8ToInt64Func(b *testing.B, supplier func() uint8, toInt64Func func(uint8) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint16ToInt64Func benchmarks a function with the signature:
//     func(uint16) int64
// ID: B-6-9
func Uint16ToInt64Func(b *testing.B, supplier func() uint16, toInt64Func func(uint16) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint32ToInt64Func benchmarks a function with the signature:
//     func(uint32) int64
// ID: B-6-10
func Uint32ToInt64Func(b *testing.B, supplier func() uint32, toInt64Func func(uint32) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint64ToInt64Func benchmarks a function with the signature:
//     func(uint64) int64
// ID: B-6-11
func Uint64ToInt64Func(b *testing.B, supplier func() uint64, toInt64Func func(uint64) int64) {
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}
