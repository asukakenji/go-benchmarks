package benchmark

import (
	"reflect"
	"testing"
)

func isBoolConsumerCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isIntConsumerCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isInt8ConsumerCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isInt16ConsumerCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isInt32ConsumerCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isInt64ConsumerCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isUintConsumerCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isUint8ConsumerCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isUint16ConsumerCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isUint32ConsumerCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func isUint64ConsumerCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setBoolConsumerCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setIntConsumerCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setInt8ConsumerCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setInt16ConsumerCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setInt32ConsumerCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setInt64ConsumerCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setUintConsumerCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setUint8ConsumerCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setUint16ConsumerCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setUint32ConsumerCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

func setUint64ConsumerCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Invalid, reflect.ValueOf(supplier).Pointer())
}

// BoolConsumer benchmarks a function with the signature:
//     func(bool)
// ID: B-0-1
func BoolConsumer(b *testing.B, supplier func() bool, consumer func(bool)) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolConsumerCalibrated(supplier) {
		panic("BoolConsumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// IntConsumer benchmarks a function with the signature:
//     func(int)
// ID: B-0-2
func IntConsumer(b *testing.B, supplier func() int, consumer func(int)) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntConsumerCalibrated(supplier) {
		panic("IntConsumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int8Consumer benchmarks a function with the signature:
//     func(int8)
// ID: B-0-3
func Int8Consumer(b *testing.B, supplier func() int8, consumer func(int8)) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ConsumerCalibrated(supplier) {
		panic("Int8Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int16Consumer benchmarks a function with the signature:
//     func(int16)
// ID: B-0-4
func Int16Consumer(b *testing.B, supplier func() int16, consumer func(int16)) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ConsumerCalibrated(supplier) {
		panic("Int16Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int32Consumer benchmarks a function with the signature:
//     func(int32)
// ID: B-0-5
func Int32Consumer(b *testing.B, supplier func() int32, consumer func(int32)) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ConsumerCalibrated(supplier) {
		panic("Int32Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Int64Consumer benchmarks a function with the signature:
//     func(int64)
// ID: B-0-6
func Int64Consumer(b *testing.B, supplier func() int64, consumer func(int64)) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ConsumerCalibrated(supplier) {
		panic("Int64Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// UintConsumer benchmarks a function with the signature:
//     func(uint)
// ID: B-0-7
func UintConsumer(b *testing.B, supplier func() uint, consumer func(uint)) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintConsumerCalibrated(supplier) {
		panic("UintConsumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint8Consumer benchmarks a function with the signature:
//     func(uint8)
// ID: B-0-8
func Uint8Consumer(b *testing.B, supplier func() uint8, consumer func(uint8)) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ConsumerCalibrated(supplier) {
		panic("Uint8Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint16Consumer benchmarks a function with the signature:
//     func(uint16)
// ID: B-0-9
func Uint16Consumer(b *testing.B, supplier func() uint16, consumer func(uint16)) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ConsumerCalibrated(supplier) {
		panic("Uint16Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint32Consumer benchmarks a function with the signature:
//     func(uint32)
// ID: B-0-10
func Uint32Consumer(b *testing.B, supplier func() uint32, consumer func(uint32)) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ConsumerCalibrated(supplier) {
		panic("Uint32Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}

// Uint64Consumer benchmarks a function with the signature:
//     func(uint64)
// ID: B-0-11
func Uint64Consumer(b *testing.B, supplier func() uint64, consumer func(uint64)) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ConsumerCalibrated(supplier) {
		panic("Uint64Consumer not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		consumer(supplier())
	}
}
