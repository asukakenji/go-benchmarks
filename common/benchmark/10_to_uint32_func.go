package benchmark

import "testing"

// BoolToUint32Func benchmarks a function with the signature:
//     func(bool) uint32
// ID: B-10-1
func BoolToUint32Func(b *testing.B, supplier func() bool, toUint32Func func(bool) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// IntToUint32Func benchmarks a function with the signature:
//     func(int) uint32
// ID: B-10-2
func IntToUint32Func(b *testing.B, supplier func() int, toUint32Func func(int) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int8ToUint32Func benchmarks a function with the signature:
//     func(int8) uint32
// ID: B-10-3
func Int8ToUint32Func(b *testing.B, supplier func() int8, toUint32Func func(int8) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int16ToUint32Func benchmarks a function with the signature:
//     func(int16) uint32
// ID: B-10-4
func Int16ToUint32Func(b *testing.B, supplier func() int16, toUint32Func func(int16) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int32ToUint32Func benchmarks a function with the signature:
//     func(int32) uint32
// ID: B-10-5
func Int32ToUint32Func(b *testing.B, supplier func() int32, toUint32Func func(int32) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int64ToUint32Func benchmarks a function with the signature:
//     func(int64) uint32
// ID: B-10-6
func Int64ToUint32Func(b *testing.B, supplier func() int64, toUint32Func func(int64) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// UintToUint32Func benchmarks a function with the signature:
//     func(uint) uint32
// ID: B-10-7
func UintToUint32Func(b *testing.B, supplier func() uint, toUint32Func func(uint) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint8ToUint32Func benchmarks a function with the signature:
//     func(uint8) uint32
// ID: B-10-8
func Uint8ToUint32Func(b *testing.B, supplier func() uint8, toUint32Func func(uint8) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint16ToUint32Func benchmarks a function with the signature:
//     func(uint16) uint32
// ID: B-10-9
func Uint16ToUint32Func(b *testing.B, supplier func() uint16, toUint32Func func(uint16) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint32ToUint32Func benchmarks a function with the signature:
//     func(uint32) uint32
// ID: B-10-10
func Uint32ToUint32Func(b *testing.B, supplier func() uint32, toUint32Func func(uint32) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint64ToUint32Func benchmarks a function with the signature:
//     func(uint64) uint32
// ID: B-10-11
func Uint64ToUint32Func(b *testing.B, supplier func() uint64, toUint32Func func(uint64) uint32) {
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}
