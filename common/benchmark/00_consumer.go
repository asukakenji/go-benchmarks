package benchmark

import "testing"

// BoolConsumer benchmarks a function with the signature:
//     func(bool)
// ID: B-0-1
func BoolConsumer(b *testing.B, supplier func() bool, consumer func(bool)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// IntConsumer benchmarks a function with the signature:
//     func(int)
// ID: B-0-2
func IntConsumer(b *testing.B, supplier func() int, consumer func(int)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int8Consumer benchmarks a function with the signature:
//     func(int8)
// ID: B-0-3
func Int8Consumer(b *testing.B, supplier func() int8, consumer func(int8)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int16Consumer benchmarks a function with the signature:
//     func(int16)
// ID: B-0-4
func Int16Consumer(b *testing.B, supplier func() int16, consumer func(int16)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int32Consumer benchmarks a function with the signature:
//     func(int32)
// ID: B-0-5
func Int32Consumer(b *testing.B, supplier func() int32, consumer func(int32)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int64Consumer benchmarks a function with the signature:
//     func(int64)
// ID: B-0-6
func Int64Consumer(b *testing.B, supplier func() int64, consumer func(int64)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// UintConsumer benchmarks a function with the signature:
//     func(uint)
// ID: B-0-7
func UintConsumer(b *testing.B, supplier func() uint, consumer func(uint)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint8Consumer benchmarks a function with the signature:
//     func(uint8)
// ID: B-0-8
func Uint8Consumer(b *testing.B, supplier func() uint8, consumer func(uint8)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint16Consumer benchmarks a function with the signature:
//     func(uint16)
// ID: B-0-9
func Uint16Consumer(b *testing.B, supplier func() uint16, consumer func(uint16)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint32Consumer benchmarks a function with the signature:
//     func(uint32)
// ID: B-0-10
func Uint32Consumer(b *testing.B, supplier func() uint32, consumer func(uint32)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint64Consumer benchmarks a function with the signature:
//     func(uint64)
// ID: B-0-11
func Uint64Consumer(b *testing.B, supplier func() uint64, consumer func(uint64)) {
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}
