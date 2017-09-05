package benchmark

import (
	"reflect"
	"testing"
)

func isRunnableCalibrated(runnable func()) bool {
	return isCalibrated(reflect.Invalid, reflect.Invalid, reflect.ValueOf(runnable).Pointer())
}

func isBoolSupplierCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Invalid, reflect.Bool, reflect.ValueOf(supplier).Pointer())
}

func isIntSupplierCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Invalid, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isInt8SupplierCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Invalid, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isInt16SupplierCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Invalid, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isInt32SupplierCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Invalid, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isInt64SupplierCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Invalid, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUintSupplierCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Invalid, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUint8SupplierCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Invalid, reflect.Uint8, reflect.ValueOf(supplier).Pointer())
}

func isUint16SupplierCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Invalid, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUint32SupplierCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Invalid, reflect.Uint32, reflect.ValueOf(supplier).Pointer())
}

func isUint64SupplierCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Invalid, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

// Runnable benchmarks a function with the signature:
//     func()
// ID: B-0-0
func Runnable(b *testing.B, runnable func()) {
	for i, count := 0, b.N; i < count; i++ {
		runnable()
	}
	setCalibrated(reflect.Invalid, reflect.Invalid, reflect.ValueOf(runnable).Pointer())
}

// BoolSupplier benchmarks a function with the signature:
//     func() bool
// ID: B-1-0
func BoolSupplier(b *testing.B, supplier func() bool) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Bool, reflect.ValueOf(supplier).Pointer())
}

// IntSupplier benchmarks a function with the signature:
//     func() int
// ID: B-2-0
func IntSupplier(b *testing.B, supplier func() int) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

// Int8Supplier benchmarks a function with the signature:
//     func() int8
// ID: B-3-0
func Int8Supplier(b *testing.B, supplier func() int8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

// Int16Supplier benchmarks a function with the signature:
//     func() int16
// ID: B-4-0
func Int16Supplier(b *testing.B, supplier func() int16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

// Int32Supplier benchmarks a function with the signature:
//     func() int32
// ID: B-5-0
func Int32Supplier(b *testing.B, supplier func() int32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

// Int64Supplier benchmarks a function with the signature:
//     func() int64
// ID: B-6-0
func Int64Supplier(b *testing.B, supplier func() int64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

// UintSupplier benchmarks a function with the signature:
//     func() uint
// ID: B-7-0
func UintSupplier(b *testing.B, supplier func() uint) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

// Uint8Supplier benchmarks a function with the signature:
//     func() uint8
// ID: B-8-0
func Uint8Supplier(b *testing.B, supplier func() uint8) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Uint8, reflect.ValueOf(supplier).Pointer())
}

// Uint16Supplier benchmarks a function with the signature:
//     func() uint16
// ID: B-9-0
func Uint16Supplier(b *testing.B, supplier func() uint16) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

// Uint32Supplier benchmarks a function with the signature:
//     func() uint32
// ID: B-10-0
func Uint32Supplier(b *testing.B, supplier func() uint32) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Uint32, reflect.ValueOf(supplier).Pointer())
}

// Uint64Supplier benchmarks a function with the signature:
//     func() uint64
// ID: B-11-0
func Uint64Supplier(b *testing.B, supplier func() uint64) {
	for i, count := 0, b.N; i < count; i++ {
		supplier()
	}
	setCalibrated(reflect.Invalid, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}
