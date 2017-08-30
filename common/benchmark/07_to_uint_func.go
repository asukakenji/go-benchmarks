package benchmark

import "testing"

// BoolToUintFunc benchmarks a function with the signature:
//     func(bool) uint
// ID: B-7-1
func BoolToUintFunc(b *testing.B, supplier func() bool, toUintFunc func(bool) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// IntToUintFunc benchmarks a function with the signature:
//     func(int) uint
// ID: B-7-2
func IntToUintFunc(b *testing.B, supplier func() int, toUintFunc func(int) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int8ToUintFunc benchmarks a function with the signature:
//     func(int8) uint
// ID: B-7-3
func Int8ToUintFunc(b *testing.B, supplier func() int8, toUintFunc func(int8) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int16ToUintFunc benchmarks a function with the signature:
//     func(int16) uint
// ID: B-7-4
func Int16ToUintFunc(b *testing.B, supplier func() int16, toUintFunc func(int16) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int32ToUintFunc benchmarks a function with the signature:
//     func(int32) uint
// ID: B-7-5
func Int32ToUintFunc(b *testing.B, supplier func() int32, toUintFunc func(int32) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int64ToUintFunc benchmarks a function with the signature:
//     func(int64) uint
// ID: B-7-6
func Int64ToUintFunc(b *testing.B, supplier func() int64, toUintFunc func(int64) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// UintToUintFunc benchmarks a function with the signature:
//     func(uint) uint
// ID: B-7-7
func UintToUintFunc(b *testing.B, supplier func() uint, toUintFunc func(uint) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint8ToUintFunc benchmarks a function with the signature:
//     func(uint8) uint
// ID: B-7-8
func Uint8ToUintFunc(b *testing.B, supplier func() uint8, toUintFunc func(uint8) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint16ToUintFunc benchmarks a function with the signature:
//     func(uint16) uint
// ID: B-7-9
func Uint16ToUintFunc(b *testing.B, supplier func() uint16, toUintFunc func(uint16) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint32ToUintFunc benchmarks a function with the signature:
//     func(uint32) uint
// ID: B-7-10
func Uint32ToUintFunc(b *testing.B, supplier func() uint32, toUintFunc func(uint32) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint64ToUintFunc benchmarks a function with the signature:
//     func(uint64) uint
// ID: B-7-11
func Uint64ToUintFunc(b *testing.B, supplier func() uint64, toUintFunc func(uint64) uint) {
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}
