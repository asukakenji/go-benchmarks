package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToInt8FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isIntToInt8FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToInt8FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToInt8FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToInt8FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToInt8FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isUintToInt8FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToInt8FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToInt8FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToInt8FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToInt8FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setBoolToInt8FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setIntToInt8FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToInt8FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToInt8FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToInt8FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToInt8FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setUintToInt8FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToInt8FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToInt8FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToInt8FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToInt8FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Int8, reflect.ValueOf(supplier).Pointer())
}

// BoolToInt8Func benchmarks a function with the signature:
//     func(bool) int8
// ID: B-3-1
func BoolToInt8Func(b *testing.B, supplier func() bool, toInt8Func func(bool) int8) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToInt8FuncCalibrated(supplier) {
		panic("BoolToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// IntToInt8Func benchmarks a function with the signature:
//     func(int) int8
// ID: B-3-2
func IntToInt8Func(b *testing.B, supplier func() int, toInt8Func func(int) int8) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToInt8FuncCalibrated(supplier) {
		panic("IntToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int8ToInt8Func benchmarks a function with the signature:
//     func(int8) int8
// ID: B-3-3
func Int8ToInt8Func(b *testing.B, supplier func() int8, toInt8Func func(int8) int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToInt8FuncCalibrated(supplier) {
		panic("Int8ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int16ToInt8Func benchmarks a function with the signature:
//     func(int16) int8
// ID: B-3-4
func Int16ToInt8Func(b *testing.B, supplier func() int16, toInt8Func func(int16) int8) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToInt8FuncCalibrated(supplier) {
		panic("Int16ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int32ToInt8Func benchmarks a function with the signature:
//     func(int32) int8
// ID: B-3-5
func Int32ToInt8Func(b *testing.B, supplier func() int32, toInt8Func func(int32) int8) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToInt8FuncCalibrated(supplier) {
		panic("Int32ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Int64ToInt8Func benchmarks a function with the signature:
//     func(int64) int8
// ID: B-3-6
func Int64ToInt8Func(b *testing.B, supplier func() int64, toInt8Func func(int64) int8) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToInt8FuncCalibrated(supplier) {
		panic("Int64ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// UintToInt8Func benchmarks a function with the signature:
//     func(uint) int8
// ID: B-3-7
func UintToInt8Func(b *testing.B, supplier func() uint, toInt8Func func(uint) int8) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToInt8FuncCalibrated(supplier) {
		panic("UintToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint8ToInt8Func benchmarks a function with the signature:
//     func(uint8) int8
// ID: B-3-8
func Uint8ToInt8Func(b *testing.B, supplier func() uint8, toInt8Func func(uint8) int8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToInt8FuncCalibrated(supplier) {
		panic("Uint8ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint16ToInt8Func benchmarks a function with the signature:
//     func(uint16) int8
// ID: B-3-9
func Uint16ToInt8Func(b *testing.B, supplier func() uint16, toInt8Func func(uint16) int8) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToInt8FuncCalibrated(supplier) {
		panic("Uint16ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint32ToInt8Func benchmarks a function with the signature:
//     func(uint32) int8
// ID: B-3-10
func Uint32ToInt8Func(b *testing.B, supplier func() uint32, toInt8Func func(uint32) int8) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToInt8FuncCalibrated(supplier) {
		panic("Uint32ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}

// Uint64ToInt8Func benchmarks a function with the signature:
//     func(int8) int8
// ID: B-3-11
func Uint64ToInt8Func(b *testing.B, supplier func() uint64, toInt8Func func(uint64) int8) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToInt8FuncCalibrated(supplier) {
		panic("Uint64ToInt8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt8Func(supplier())
	}
}
