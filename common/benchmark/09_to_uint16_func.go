package benchmark

import "testing"

// BoolToUint16Func benchmarks a function with the signature:
//     func(bool) uint16
// ID: B-9-1
func BoolToUint16Func(b *testing.B, supplier func() bool, toUint16Func func(bool) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// IntToUint16Func benchmarks a function with the signature:
//     func(int) uint16
// ID: B-9-2
func IntToUint16Func(b *testing.B, supplier func() int, toUint16Func func(int) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int8ToUint16Func benchmarks a function with the signature:
//     func(int8) uint16
// ID: B-9-3
func Int8ToUint16Func(b *testing.B, supplier func() int8, toUint16Func func(int8) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int16ToUint16Func benchmarks a function with the signature:
//     func(int16) uint16
// ID: B-9-4
func Int16ToUint16Func(b *testing.B, supplier func() int16, toUint16Func func(int16) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int32ToUint16Func benchmarks a function with the signature:
//     func(int32) uint16
// ID: B-9-5
func Int32ToUint16Func(b *testing.B, supplier func() int32, toUint16Func func(int32) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int64ToUint16Func benchmarks a function with the signature:
//     func(int64) uint16
// ID: B-9-6
func Int64ToUint16Func(b *testing.B, supplier func() int64, toUint16Func func(int64) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// UintToUint16Func benchmarks a function with the signature:
//     func(uint) uint16
// ID: B-9-7
func UintToUint16Func(b *testing.B, supplier func() uint, toUint16Func func(uint) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint8ToUint16Func benchmarks a function with the signature:
//     func(uint8) uint16
// ID: B-9-8
func Uint8ToUint16Func(b *testing.B, supplier func() uint8, toUint16Func func(uint8) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint16ToUint16Func benchmarks a function with the signature:
//     func(uint16) uint16
// ID: B-9-9
func Uint16ToUint16Func(b *testing.B, supplier func() uint16, toUint16Func func(uint16) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint32ToUint16Func benchmarks a function with the signature:
//     func(uint32) uint16
// ID: B-9-10
func Uint32ToUint16Func(b *testing.B, supplier func() uint32, toUint16Func func(uint32) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint64ToUint16Func benchmarks a function with the signature:
//     func(uint64) uint16
// ID: B-9-11
func Uint64ToUint16Func(b *testing.B, supplier func() uint64, toUint16Func func(uint64) uint16) {
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}
