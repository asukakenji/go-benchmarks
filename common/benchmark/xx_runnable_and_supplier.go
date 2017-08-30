package benchmark

import "testing"

// Runnable benchmarks a function with the signature:
//     func()
// ID: B-0-0
func Runnable(b *testing.B, runnable func()) {
	for i, count := 0, b.N; i < count; i++ {
		runnable()
	}
}

// BoolSupplier benchmarks a function with the signature:
//     func() bool
// ID: B-1-0
func BoolSupplier(b *testing.B, supplier func() bool) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// IntSupplier benchmarks a function with the signature:
//     func() int
// ID: B-2-0
func IntSupplier(b *testing.B, supplier func() int) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Int8Supplier benchmarks a function with the signature:
//     func() int8
// ID: B-3-0
func Int8Supplier(b *testing.B, supplier func() int8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Int16Supplier benchmarks a function with the signature:
//     func() int16
// ID: B-4-0
func Int16Supplier(b *testing.B, supplier func() int16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Int32Supplier benchmarks a function with the signature:
//     func() int32
// ID: B-5-0
func Int32Supplier(b *testing.B, supplier func() int32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Int64Supplier benchmarks a function with the signature:
//     func() int64
// ID: B-6-0
func Int64Supplier(b *testing.B, supplier func() int64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// UintSupplier benchmarks a function with the signature:
//     func() uint
// ID: B-7-0
func UintSupplier(b *testing.B, supplier func() uint) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Uint8Supplier benchmarks a function with the signature:
//     func() uint8
// ID: B-8-0
func Uint8Supplier(b *testing.B, supplier func() uint8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Uint16Supplier benchmarks a function with the signature:
//     func() uint16
// ID: B-9-0
func Uint16Supplier(b *testing.B, supplier func() uint16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Uint32Supplier benchmarks a function with the signature:
//     func() uint32
// ID: B-10-0
func Uint32Supplier(b *testing.B, supplier func() uint32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}

// Uint64Supplier benchmarks a function with the signature:
//     func() uint64
// ID: B-11-0
func Uint64Supplier(b *testing.B, supplier func() uint64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
}
