package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolConsumerCalibrated   = map[uintptr]struct{}{}
	_isIntConsumerCalibrated    = map[uintptr]struct{}{}
	_isInt8ConsumerCalibrated   = map[uintptr]struct{}{}
	_isInt16ConsumerCalibrated  = map[uintptr]struct{}{}
	_isInt32ConsumerCalibrated  = map[uintptr]struct{}{}
	_isInt64ConsumerCalibrated  = map[uintptr]struct{}{}
	_isUintConsumerCalibrated   = map[uintptr]struct{}{}
	_isUint8ConsumerCalibrated  = map[uintptr]struct{}{}
	_isUint16ConsumerCalibrated = map[uintptr]struct{}{}
	_isUint32ConsumerCalibrated = map[uintptr]struct{}{}
	_isUint64ConsumerCalibrated = map[uintptr]struct{}{}
)

func isBoolConsumerCalibrated(supplier func() bool) bool {
	_, ok := _isBoolConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntConsumerCalibrated(supplier func() int) bool {
	_, ok := _isIntConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ConsumerCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ConsumerCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ConsumerCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ConsumerCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintConsumerCalibrated(supplier func() uint) bool {
	_, ok := _isUintConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ConsumerCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ConsumerCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ConsumerCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ConsumerCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolConsumerCalibrated(supplier func() bool) {
	_isBoolConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntConsumerCalibrated(supplier func() int) {
	_isIntConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ConsumerCalibrated(supplier func() int8) {
	_isInt8ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ConsumerCalibrated(supplier func() int16) {
	_isInt16ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ConsumerCalibrated(supplier func() int32) {
	_isInt32ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ConsumerCalibrated(supplier func() int64) {
	_isInt64ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintConsumerCalibrated(supplier func() uint) {
	_isUintConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ConsumerCalibrated(supplier func() uint8) {
	_isUint8ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ConsumerCalibrated(supplier func() uint16) {
	_isUint16ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ConsumerCalibrated(supplier func() uint32) {
	_isUint32ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ConsumerCalibrated(supplier func() uint64) {
	_isUint64ConsumerCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
