package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToInt64FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isIntToInt64FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToInt64FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToInt64FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToInt64FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToInt64FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUintToInt64FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToInt64FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToInt64FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToInt64FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToInt64FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setBoolToInt64FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setIntToInt64FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToInt64FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToInt64FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToInt64FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToInt64FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setUintToInt64FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToInt64FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToInt64FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToInt64FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToInt64FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Int64, reflect.ValueOf(supplier).Pointer())
}

// BoolToInt64Func benchmarks a function with the signature:
//     func(bool) int64
// ID: B-6-1
func BoolToInt64Func(b *testing.B, supplier func() bool, toInt64Func func(bool) int64) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToInt64FuncCalibrated(supplier) {
		panic("BoolToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// IntToInt64Func benchmarks a function with the signature:
//     func(int) int64
// ID: B-6-2
func IntToInt64Func(b *testing.B, supplier func() int, toInt64Func func(int) int64) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToInt64FuncCalibrated(supplier) {
		panic("IntToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int8ToInt64Func benchmarks a function with the signature:
//     func(int8) int64
// ID: B-6-3
func Int8ToInt64Func(b *testing.B, supplier func() int8, toInt64Func func(int8) int64) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToInt64FuncCalibrated(supplier) {
		panic("Int8ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int16ToInt64Func benchmarks a function with the signature:
//     func(int16) int64
// ID: B-6-4
func Int16ToInt64Func(b *testing.B, supplier func() int16, toInt64Func func(int16) int64) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToInt64FuncCalibrated(supplier) {
		panic("Int16ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int32ToInt64Func benchmarks a function with the signature:
//     func(int32) int64
// ID: B-6-5
func Int32ToInt64Func(b *testing.B, supplier func() int32, toInt64Func func(int32) int64) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToInt64FuncCalibrated(supplier) {
		panic("Int32ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Int64ToInt64Func benchmarks a function with the signature:
//     func(int64) int64
// ID: B-6-6
func Int64ToInt64Func(b *testing.B, supplier func() int64, toInt64Func func(int64) int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToInt64FuncCalibrated(supplier) {
		panic("Int64ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// UintToInt64Func benchmarks a function with the signature:
//     func(uint) int64
// ID: B-6-7
func UintToInt64Func(b *testing.B, supplier func() uint, toInt64Func func(uint) int64) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToInt64FuncCalibrated(supplier) {
		panic("UintToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint8ToInt64Func benchmarks a function with the signature:
//     func(uint8) int64
// ID: B-6-8
func Uint8ToInt64Func(b *testing.B, supplier func() uint8, toInt64Func func(uint8) int64) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToInt64FuncCalibrated(supplier) {
		panic("Uint8ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint16ToInt64Func benchmarks a function with the signature:
//     func(uint16) int64
// ID: B-6-9
func Uint16ToInt64Func(b *testing.B, supplier func() uint16, toInt64Func func(uint16) int64) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToInt64FuncCalibrated(supplier) {
		panic("Uint16ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint32ToInt64Func benchmarks a function with the signature:
//     func(uint32) int64
// ID: B-6-10
func Uint32ToInt64Func(b *testing.B, supplier func() uint32, toInt64Func func(uint32) int64) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToInt64FuncCalibrated(supplier) {
		panic("Uint32ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}

// Uint64ToInt64Func benchmarks a function with the signature:
//     func(int64) int64
// ID: B-6-11
func Uint64ToInt64Func(b *testing.B, supplier func() uint64, toInt64Func func(uint64) int64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToInt64FuncCalibrated(supplier) {
		panic("Uint64ToInt64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt64Func(supplier())
	}
}
