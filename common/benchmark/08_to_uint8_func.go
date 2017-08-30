package benchmark

import "testing"

// BoolToUint8Func benchmarks a function with the signature:
//     func(bool) uint8
// ID: B-8-1
func BoolToUint8Func(b *testing.B, supplier func() bool, toUint8Func func(bool) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// IntToUint8Func benchmarks a function with the signature:
//     func(int) uint8
// ID: B-8-2
func IntToUint8Func(b *testing.B, supplier func() int, toUint8Func func(int) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int8ToUint8Func benchmarks a function with the signature:
//     func(int8) uint8
// ID: B-8-3
func Int8ToUint8Func(b *testing.B, supplier func() int8, toUint8Func func(int8) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int16ToUint8Func benchmarks a function with the signature:
//     func(int16) uint8
// ID: B-8-4
func Int16ToUint8Func(b *testing.B, supplier func() int16, toUint8Func func(int16) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int32ToUint8Func benchmarks a function with the signature:
//     func(int32) uint8
// ID: B-8-5
func Int32ToUint8Func(b *testing.B, supplier func() int32, toUint8Func func(int32) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int64ToUint8Func benchmarks a function with the signature:
//     func(int64) uint8
// ID: B-8-6
func Int64ToUint8Func(b *testing.B, supplier func() int64, toUint8Func func(int64) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// UintToUint8Func benchmarks a function with the signature:
//     func(uint) uint8
// ID: B-8-7
func UintToUint8Func(b *testing.B, supplier func() uint, toUint8Func func(uint) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint8ToUint8Func benchmarks a function with the signature:
//     func(uint8) uint8
// ID: B-8-8
func Uint8ToUint8Func(b *testing.B, supplier func() uint8, toUint8Func func(uint8) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint16ToUint8Func benchmarks a function with the signature:
//     func(uint16) uint8
// ID: B-8-9
func Uint16ToUint8Func(b *testing.B, supplier func() uint16, toUint8Func func(uint16) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint32ToUint8Func benchmarks a function with the signature:
//     func(uint32) uint8
// ID: B-8-10
func Uint32ToUint8Func(b *testing.B, supplier func() uint32, toUint8Func func(uint32) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint64ToUint8Func benchmarks a function with the signature:
//     func(uint64) uint8
// ID: B-8-11
func Uint64ToUint8Func(b *testing.B, supplier func() uint64, toUint8Func func(uint64) uint8) {
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}
