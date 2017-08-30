package benchmark

import "testing"

// BoolConsumer benchmarks a function with the signature:
//     func(bool)
// ID: B-0-1
func BoolConsumer(b *testing.B, genFunc func() bool, f func(bool)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// IntConsumer benchmarks a function with the signature:
//     func(int)
// ID: B-0-2
func IntConsumer(b *testing.B, genFunc func() int, f func(int)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Int8Consumer benchmarks a function with the signature:
//     func(int8)
// ID: B-0-3
func Int8Consumer(b *testing.B, genFunc func() int8, f func(int8)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Int16Consumer benchmarks a function with the signature:
//     func(int16)
// ID: B-0-4
func Int16Consumer(b *testing.B, genFunc func() int16, f func(int16)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Int32Consumer benchmarks a function with the signature:
//     func(int32)
// ID: B-0-5
func Int32Consumer(b *testing.B, genFunc func() int32, f func(int32)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Int64Consumer benchmarks a function with the signature:
//     func(int64)
// ID: B-0-6
func Int64Consumer(b *testing.B, genFunc func() int64, f func(int64)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// UintConsumer benchmarks a function with the signature:
//     func(uint)
// ID: B-0-7
func UintConsumer(b *testing.B, genFunc func() uint, f func(uint)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Uint8Consumer benchmarks a function with the signature:
//     func(uint8)
// ID: B-0-8
func Uint8Consumer(b *testing.B, genFunc func() uint8, f func(uint8)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Uint16Consumer benchmarks a function with the signature:
//     func(uint16)
// ID: B-0-9
func Uint16Consumer(b *testing.B, genFunc func() uint16, f func(uint16)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Uint32Consumer benchmarks a function with the signature:
//     func(uint32)
// ID: B-0-10
func Uint32Consumer(b *testing.B, genFunc func() uint32, f func(uint32)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}

// Uint64Consumer benchmarks a function with the signature:
//     func(uint64)
// ID: B-0-11
func Uint64Consumer(b *testing.B, genFunc func() uint64, f func(uint64)) {
	for i, count := 0, b.N; i < count; i++ {
		f(genFunc())
	}
}
