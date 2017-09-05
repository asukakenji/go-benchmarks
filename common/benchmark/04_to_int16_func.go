package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToInt16FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isIntToInt16FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToInt16FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToInt16FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToInt16FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToInt16FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isUintToInt16FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToInt16FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToInt16FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToInt16FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToInt16FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setBoolToInt16FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setIntToInt16FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToInt16FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToInt16FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToInt16FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToInt16FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setUintToInt16FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToInt16FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToInt16FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToInt16FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToInt16FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Int16, reflect.ValueOf(supplier).Pointer())
}

// BoolToInt16Func benchmarks a function with the signature:
//     func(bool) int16
// ID: B-4-1
func BoolToInt16Func(b *testing.B, supplier func() bool, toInt16Func func(bool) int16) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToInt16FuncCalibrated(supplier) {
		panic("BoolToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// IntToInt16Func benchmarks a function with the signature:
//     func(int) int16
// ID: B-4-2
func IntToInt16Func(b *testing.B, supplier func() int, toInt16Func func(int) int16) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToInt16FuncCalibrated(supplier) {
		panic("IntToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int8ToInt16Func benchmarks a function with the signature:
//     func(int8) int16
// ID: B-4-3
func Int8ToInt16Func(b *testing.B, supplier func() int8, toInt16Func func(int8) int16) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToInt16FuncCalibrated(supplier) {
		panic("Int8ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int16ToInt16Func benchmarks a function with the signature:
//     func(int16) int16
// ID: B-4-4
func Int16ToInt16Func(b *testing.B, supplier func() int16, toInt16Func func(int16) int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToInt16FuncCalibrated(supplier) {
		panic("Int16ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int32ToInt16Func benchmarks a function with the signature:
//     func(int32) int16
// ID: B-4-5
func Int32ToInt16Func(b *testing.B, supplier func() int32, toInt16Func func(int32) int16) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToInt16FuncCalibrated(supplier) {
		panic("Int32ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Int64ToInt16Func benchmarks a function with the signature:
//     func(int64) int16
// ID: B-4-6
func Int64ToInt16Func(b *testing.B, supplier func() int64, toInt16Func func(int64) int16) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToInt16FuncCalibrated(supplier) {
		panic("Int64ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// UintToInt16Func benchmarks a function with the signature:
//     func(uint) int16
// ID: B-4-7
func UintToInt16Func(b *testing.B, supplier func() uint, toInt16Func func(uint) int16) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToInt16FuncCalibrated(supplier) {
		panic("UintToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint8ToInt16Func benchmarks a function with the signature:
//     func(uint8) int16
// ID: B-4-8
func Uint8ToInt16Func(b *testing.B, supplier func() uint8, toInt16Func func(uint8) int16) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToInt16FuncCalibrated(supplier) {
		panic("Uint8ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint16ToInt16Func benchmarks a function with the signature:
//     func(uint16) int16
// ID: B-4-9
func Uint16ToInt16Func(b *testing.B, supplier func() uint16, toInt16Func func(uint16) int16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToInt16FuncCalibrated(supplier) {
		panic("Uint16ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint32ToInt16Func benchmarks a function with the signature:
//     func(uint32) int16
// ID: B-4-10
func Uint32ToInt16Func(b *testing.B, supplier func() uint32, toInt16Func func(uint32) int16) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToInt16FuncCalibrated(supplier) {
		panic("Uint32ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}

// Uint64ToInt16Func benchmarks a function with the signature:
//     func(int16) int16
// ID: B-4-11
func Uint64ToInt16Func(b *testing.B, supplier func() uint64, toInt16Func func(uint64) int16) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToInt16FuncCalibrated(supplier) {
		panic("Uint64ToInt16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt16Func(supplier())
	}
}
