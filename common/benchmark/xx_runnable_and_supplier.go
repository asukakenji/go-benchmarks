package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isRunnableCalibrated       = map[uintptr]struct{}{}
	_isBoolSupplierCalibrated   = map[uintptr]struct{}{}
	_isIntSupplierCalibrated    = map[uintptr]struct{}{}
	_isInt8SupplierCalibrated   = map[uintptr]struct{}{}
	_isInt16SupplierCalibrated  = map[uintptr]struct{}{}
	_isInt32SupplierCalibrated  = map[uintptr]struct{}{}
	_isInt64SupplierCalibrated  = map[uintptr]struct{}{}
	_isUintSupplierCalibrated   = map[uintptr]struct{}{}
	_isUint8SupplierCalibrated  = map[uintptr]struct{}{}
	_isUint16SupplierCalibrated = map[uintptr]struct{}{}
	_isUint32SupplierCalibrated = map[uintptr]struct{}{}
	_isUint64SupplierCalibrated = map[uintptr]struct{}{}
)

func isRunnableCalibrated(runnable func()) bool {
	_, ok := _isRunnableCalibrated[reflect.ValueOf(runnable).Pointer()]
	return ok
}

func isBoolSupplierCalibrated(supplier func() bool) bool {
	_, ok := _isBoolSupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntSupplierCalibrated(supplier func() int) bool {
	_, ok := _isIntSupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8SupplierCalibrated(supplier func() int8) bool {
	_, ok := _isInt8SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16SupplierCalibrated(supplier func() int16) bool {
	_, ok := _isInt16SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32SupplierCalibrated(supplier func() int32) bool {
	_, ok := _isInt32SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64SupplierCalibrated(supplier func() int64) bool {
	_, ok := _isInt64SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintSupplierCalibrated(supplier func() uint) bool {
	_, ok := _isUintSupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8SupplierCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16SupplierCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32SupplierCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64SupplierCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64SupplierCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

// Runnable benchmarks a function with the signature:
//     func()
// ID: B-0-0
func Runnable(b *testing.B, runnable func()) {
	for i, count := 0, b.N; i < count; i++ {
		runnable()
	}
	_isRunnableCalibrated[reflect.ValueOf(runnable).Pointer()] = struct{}{}
}

// BoolSupplier benchmarks a function with the signature:
//     func() bool
// ID: B-1-0
func BoolSupplier(b *testing.B, supplier func() bool) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isBoolSupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// IntSupplier benchmarks a function with the signature:
//     func() int
// ID: B-2-0
func IntSupplier(b *testing.B, supplier func() int) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isIntSupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Int8Supplier benchmarks a function with the signature:
//     func() int8
// ID: B-3-0
func Int8Supplier(b *testing.B, supplier func() int8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isInt8SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Int16Supplier benchmarks a function with the signature:
//     func() int16
// ID: B-4-0
func Int16Supplier(b *testing.B, supplier func() int16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isInt16SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Int32Supplier benchmarks a function with the signature:
//     func() int32
// ID: B-5-0
func Int32Supplier(b *testing.B, supplier func() int32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isInt32SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Int64Supplier benchmarks a function with the signature:
//     func() int64
// ID: B-6-0
func Int64Supplier(b *testing.B, supplier func() int64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isInt64SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// UintSupplier benchmarks a function with the signature:
//     func() uint
// ID: B-7-0
func UintSupplier(b *testing.B, supplier func() uint) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isUintSupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Uint8Supplier benchmarks a function with the signature:
//     func() uint8
// ID: B-8-0
func Uint8Supplier(b *testing.B, supplier func() uint8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isUint8SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Uint16Supplier benchmarks a function with the signature:
//     func() uint16
// ID: B-9-0
func Uint16Supplier(b *testing.B, supplier func() uint16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isUint16SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Uint32Supplier benchmarks a function with the signature:
//     func() uint32
// ID: B-10-0
func Uint32Supplier(b *testing.B, supplier func() uint32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isUint32SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// Uint64Supplier benchmarks a function with the signature:
//     func() uint64
// ID: B-11-0
func Uint64Supplier(b *testing.B, supplier func() uint64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	_isUint64SupplierCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}
