package benchmark

import "testing"

// BoolPredicate benchmarks a function with the signature:
//     func(bool) bool
// ID: B-1-1
func BoolPredicate(b *testing.B, supplier func() bool, predicate func(bool) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// IntPredicate benchmarks a function with the signature:
//     func(int) bool
// ID: B-1-2
func IntPredicate(b *testing.B, supplier func() int, predicate func(int) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int8Predicate benchmarks a function with the signature:
//     func(int8) bool
// ID: B-1-3
func Int8Predicate(b *testing.B, supplier func() int8, predicate func(int8) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int16Predicate benchmarks a function with the signature:
//     func(int16) bool
// ID: B-1-4
func Int16Predicate(b *testing.B, supplier func() int16, predicate func(int16) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int32Predicate benchmarks a function with the signature:
//     func(int32) bool
// ID: B-1-5
func Int32Predicate(b *testing.B, supplier func() int32, predicate func(int32) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int64Predicate benchmarks a function with the signature:
//     func(int64) bool
// ID: B-1-6
func Int64Predicate(b *testing.B, supplier func() int64, predicate func(int64) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// UintPredicate benchmarks a function with the signature:
//     func(uint) bool
// ID: B-1-7
func UintPredicate(b *testing.B, supplier func() uint, predicate func(uint) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint8Predicate benchmarks a function with the signature:
//     func(uint8) bool
// ID: B-1-8
func Uint8Predicate(b *testing.B, supplier func() uint8, predicate func(uint8) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint16Predicate benchmarks a function with the signature:
//     func(uint16) bool
// ID: B-1-9
func Uint16Predicate(b *testing.B, supplier func() uint16, predicate func(uint16) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint32Predicate benchmarks a function with the signature:
//     func(uint32) bool
// ID: B-1-10
func Uint32Predicate(b *testing.B, supplier func() uint32, predicate func(uint32) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint64Predicate benchmarks a function with the signature:
//     func(uint64) bool
// ID: B-1-11
func Uint64Predicate(b *testing.B, supplier func() uint64, predicate func(uint64) bool) {
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}
