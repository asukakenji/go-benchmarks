package benchmark

import "testing"

// BoolToInt16Func benchmarks a function with the signature:
//     func(bool) int16
// ID: B-4-1
func BoolToInt16Func(b *testing.B, supplier func() bool, toInt16Func func(bool) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// IntToInt16Func benchmarks a function with the signature:
//     func(int) int16
// ID: B-4-2
func IntToInt16Func(b *testing.B, supplier func() int, toInt16Func func(int) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int8ToInt16Func benchmarks a function with the signature:
//     func(int8) int16
// ID: B-4-3
func Int8ToInt16Func(b *testing.B, supplier func() int8, toInt16Func func(int8) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int16ToInt16Func benchmarks a function with the signature:
//     func(int16) int16
// ID: B-4-4
func Int16ToInt16Func(b *testing.B, supplier func() int16, toInt16Func func(int16) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int32ToInt16Func benchmarks a function with the signature:
//     func(int32) int16
// ID: B-4-5
func Int32ToInt16Func(b *testing.B, supplier func() int32, toInt16Func func(int32) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int64ToInt16Func benchmarks a function with the signature:
//     func(int64) int16
// ID: B-4-6
func Int64ToInt16Func(b *testing.B, supplier func() int64, toInt16Func func(int64) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// UintToInt16Func benchmarks a function with the signature:
//     func(uint) int16
// ID: B-4-7
func UintToInt16Func(b *testing.B, supplier func() uint, toInt16Func func(uint) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint8ToInt16Func benchmarks a function with the signature:
//     func(uint8) int16
// ID: B-4-8
func Uint8ToInt16Func(b *testing.B, supplier func() uint8, toInt16Func func(uint8) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint16ToInt16Func benchmarks a function with the signature:
//     func(uint16) int16
// ID: B-4-9
func Uint16ToInt16Func(b *testing.B, supplier func() uint16, toInt16Func func(uint16) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint32ToInt16Func benchmarks a function with the signature:
//     func(uint32) int16
// ID: B-4-10
func Uint32ToInt16Func(b *testing.B, supplier func() uint32, toInt16Func func(uint32) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint64ToInt16Func benchmarks a function with the signature:
//     func(uint64) int16
// ID: B-4-11
func Uint64ToInt16Func(b *testing.B, supplier func() uint64, toInt16Func func(uint64) int16) {
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}
