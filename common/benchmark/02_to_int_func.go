package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToIntFuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isIntToIntFuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToIntFuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToIntFuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToIntFuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToIntFuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isUintToIntFuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToIntFuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToIntFuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToIntFuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToIntFuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setBoolToIntFuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setIntToIntFuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToIntFuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToIntFuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToIntFuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToIntFuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setUintToIntFuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToIntFuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToIntFuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToIntFuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToIntFuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Int, reflect.ValueOf(supplier).Pointer())
}

// BoolToIntFunc benchmarks a function with the signature:
//     func(bool) int
// ID: B-2-1
func BoolToIntFunc(b *testing.B, supplier func() bool, toIntFunc func(bool) int) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToIntFuncCalibrated(supplier) {
		panic("BoolToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// IntToIntFunc benchmarks a function with the signature:
//     func(int) int
// ID: B-2-2
func IntToIntFunc(b *testing.B, supplier func() int, toIntFunc func(int) int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToIntFuncCalibrated(supplier) {
		panic("IntToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int8ToIntFunc benchmarks a function with the signature:
//     func(int8) int
// ID: B-2-3
func Int8ToIntFunc(b *testing.B, supplier func() int8, toIntFunc func(int8) int) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToIntFuncCalibrated(supplier) {
		panic("Int8ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int16ToIntFunc benchmarks a function with the signature:
//     func(int16) int
// ID: B-2-4
func Int16ToIntFunc(b *testing.B, supplier func() int16, toIntFunc func(int16) int) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToIntFuncCalibrated(supplier) {
		panic("Int16ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int32ToIntFunc benchmarks a function with the signature:
//     func(int32) int
// ID: B-2-5
func Int32ToIntFunc(b *testing.B, supplier func() int32, toIntFunc func(int32) int) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToIntFuncCalibrated(supplier) {
		panic("Int32ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int64ToIntFunc benchmarks a function with the signature:
//     func(int64) int
// ID: B-2-6
func Int64ToIntFunc(b *testing.B, supplier func() int64, toIntFunc func(int64) int) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToIntFuncCalibrated(supplier) {
		panic("Int64ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// UintToIntFunc benchmarks a function with the signature:
//     func(uint) int
// ID: B-2-7
func UintToIntFunc(b *testing.B, supplier func() uint, toIntFunc func(uint) int) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToIntFuncCalibrated(supplier) {
		panic("UintToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint8ToIntFunc benchmarks a function with the signature:
//     func(uint8) int
// ID: B-2-8
func Uint8ToIntFunc(b *testing.B, supplier func() uint8, toIntFunc func(uint8) int) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToIntFuncCalibrated(supplier) {
		panic("Uint8ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint16ToIntFunc benchmarks a function with the signature:
//     func(uint16) int
// ID: B-2-9
func Uint16ToIntFunc(b *testing.B, supplier func() uint16, toIntFunc func(uint16) int) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToIntFuncCalibrated(supplier) {
		panic("Uint16ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint32ToIntFunc benchmarks a function with the signature:
//     func(uint32) int
// ID: B-2-10
func Uint32ToIntFunc(b *testing.B, supplier func() uint32, toIntFunc func(uint32) int) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToIntFuncCalibrated(supplier) {
		panic("Uint32ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint64ToIntFunc benchmarks a function with the signature:
//     func(int) int
// ID: B-2-11
func Uint64ToIntFunc(b *testing.B, supplier func() uint64, toIntFunc func(uint64) int) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToIntFuncCalibrated(supplier) {
		panic("Uint64ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}
