package benchmark

import "testing"

// BoolToInt8Func benchmarks a function with the signature:
//     func(bool) int8
// ID: B-3-1
func BoolToInt8Func(b *testing.B, supplier func() bool, toInt8Func func(bool) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// IntToInt8Func benchmarks a function with the signature:
//     func(int) int8
// ID: B-3-2
func IntToInt8Func(b *testing.B, supplier func() int, toInt8Func func(int) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int8ToInt8Func benchmarks a function with the signature:
//     func(int8) int8
// ID: B-3-3
func Int8ToInt8Func(b *testing.B, supplier func() int8, toInt8Func func(int8) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int16ToInt8Func benchmarks a function with the signature:
//     func(int16) int8
// ID: B-3-4
func Int16ToInt8Func(b *testing.B, supplier func() int16, toInt8Func func(int16) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int32ToInt8Func benchmarks a function with the signature:
//     func(int32) int8
// ID: B-3-5
func Int32ToInt8Func(b *testing.B, supplier func() int32, toInt8Func func(int32) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int64ToInt8Func benchmarks a function with the signature:
//     func(int64) int8
// ID: B-3-6
func Int64ToInt8Func(b *testing.B, supplier func() int64, toInt8Func func(int64) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// UintToInt8Func benchmarks a function with the signature:
//     func(uint) int8
// ID: B-3-7
func UintToInt8Func(b *testing.B, supplier func() uint, toInt8Func func(uint) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint8ToInt8Func benchmarks a function with the signature:
//     func(uint8) int8
// ID: B-3-8
func Uint8ToInt8Func(b *testing.B, supplier func() uint8, toInt8Func func(uint8) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint16ToInt8Func benchmarks a function with the signature:
//     func(uint16) int8
// ID: B-3-9
func Uint16ToInt8Func(b *testing.B, supplier func() uint16, toInt8Func func(uint16) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint32ToInt8Func benchmarks a function with the signature:
//     func(uint32) int8
// ID: B-3-10
func Uint32ToInt8Func(b *testing.B, supplier func() uint32, toInt8Func func(uint32) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint64ToInt8Func benchmarks a function with the signature:
//     func(uint64) int8
// ID: B-3-11
func Uint64ToInt8Func(b *testing.B, supplier func() uint64, toInt8Func func(uint64) int8) {
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}
