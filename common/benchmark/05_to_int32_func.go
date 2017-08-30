package benchmark

import "testing"

// BoolToInt32Func benchmarks a function with the signature:
//     func(bool) int32
// ID: B-5-1
func BoolToInt32Func(b *testing.B, supplier func() bool, toInt32Func func(bool) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// IntToInt32Func benchmarks a function with the signature:
//     func(int) int32
// ID: B-5-2
func IntToInt32Func(b *testing.B, supplier func() int, toInt32Func func(int) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int8ToInt32Func benchmarks a function with the signature:
//     func(int8) int32
// ID: B-5-3
func Int8ToInt32Func(b *testing.B, supplier func() int8, toInt32Func func(int8) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int16ToInt32Func benchmarks a function with the signature:
//     func(int16) int32
// ID: B-5-4
func Int16ToInt32Func(b *testing.B, supplier func() int16, toInt32Func func(int16) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int32ToInt32Func benchmarks a function with the signature:
//     func(int32) int32
// ID: B-5-5
func Int32ToInt32Func(b *testing.B, supplier func() int32, toInt32Func func(int32) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int64ToInt32Func benchmarks a function with the signature:
//     func(int64) int32
// ID: B-5-6
func Int64ToInt32Func(b *testing.B, supplier func() int64, toInt32Func func(int64) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// UintToInt32Func benchmarks a function with the signature:
//     func(uint) int32
// ID: B-5-7
func UintToInt32Func(b *testing.B, supplier func() uint, toInt32Func func(uint) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint8ToInt32Func benchmarks a function with the signature:
//     func(uint8) int32
// ID: B-5-8
func Uint8ToInt32Func(b *testing.B, supplier func() uint8, toInt32Func func(uint8) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint16ToInt32Func benchmarks a function with the signature:
//     func(uint16) int32
// ID: B-5-9
func Uint16ToInt32Func(b *testing.B, supplier func() uint16, toInt32Func func(uint16) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint32ToInt32Func benchmarks a function with the signature:
//     func(uint32) int32
// ID: B-5-10
func Uint32ToInt32Func(b *testing.B, supplier func() uint32, toInt32Func func(uint32) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint64ToInt32Func benchmarks a function with the signature:
//     func(uint64) int32
// ID: B-5-11
func Uint64ToInt32Func(b *testing.B, supplier func() uint64, toInt32Func func(uint64) int32) {
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}
