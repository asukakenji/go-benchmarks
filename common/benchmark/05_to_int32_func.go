package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToInt32FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isIntToInt32FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToInt32FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToInt32FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToInt32FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToInt32FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isUintToInt32FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToInt32FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToInt32FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToInt32FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToInt32FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setBoolToInt32FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setIntToInt32FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToInt32FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToInt32FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToInt32FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToInt32FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setUintToInt32FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToInt32FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToInt32FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToInt32FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToInt32FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Int32, reflect.ValueOf(supplier).Pointer())
}

// BoolToInt32Func benchmarks a function with the signature:
//     func(bool) int32
// ID: B-5-1
func BoolToInt32Func(b *testing.B, supplier func() bool, toInt32Func func(bool) int32) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToInt32FuncCalibrated(supplier) {
		panic("BoolToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// IntToInt32Func benchmarks a function with the signature:
//     func(int) int32
// ID: B-5-2
func IntToInt32Func(b *testing.B, supplier func() int, toInt32Func func(int) int32) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToInt32FuncCalibrated(supplier) {
		panic("IntToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int8ToInt32Func benchmarks a function with the signature:
//     func(int8) int32
// ID: B-5-3
func Int8ToInt32Func(b *testing.B, supplier func() int8, toInt32Func func(int8) int32) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToInt32FuncCalibrated(supplier) {
		panic("Int8ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int16ToInt32Func benchmarks a function with the signature:
//     func(int16) int32
// ID: B-5-4
func Int16ToInt32Func(b *testing.B, supplier func() int16, toInt32Func func(int16) int32) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToInt32FuncCalibrated(supplier) {
		panic("Int16ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int32ToInt32Func benchmarks a function with the signature:
//     func(int32) int32
// ID: B-5-5
func Int32ToInt32Func(b *testing.B, supplier func() int32, toInt32Func func(int32) int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToInt32FuncCalibrated(supplier) {
		panic("Int32ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Int64ToInt32Func benchmarks a function with the signature:
//     func(int64) int32
// ID: B-5-6
func Int64ToInt32Func(b *testing.B, supplier func() int64, toInt32Func func(int64) int32) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToInt32FuncCalibrated(supplier) {
		panic("Int64ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// UintToInt32Func benchmarks a function with the signature:
//     func(uint) int32
// ID: B-5-7
func UintToInt32Func(b *testing.B, supplier func() uint, toInt32Func func(uint) int32) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToInt32FuncCalibrated(supplier) {
		panic("UintToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint8ToInt32Func benchmarks a function with the signature:
//     func(uint8) int32
// ID: B-5-8
func Uint8ToInt32Func(b *testing.B, supplier func() uint8, toInt32Func func(uint8) int32) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToInt32FuncCalibrated(supplier) {
		panic("Uint8ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint16ToInt32Func benchmarks a function with the signature:
//     func(uint16) int32
// ID: B-5-9
func Uint16ToInt32Func(b *testing.B, supplier func() uint16, toInt32Func func(uint16) int32) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToInt32FuncCalibrated(supplier) {
		panic("Uint16ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint32ToInt32Func benchmarks a function with the signature:
//     func(uint32) int32
// ID: B-5-10
func Uint32ToInt32Func(b *testing.B, supplier func() uint32, toInt32Func func(uint32) int32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToInt32FuncCalibrated(supplier) {
		panic("Uint32ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}

// Uint64ToInt32Func benchmarks a function with the signature:
//     func(int32) int32
// ID: B-5-11
func Uint64ToInt32Func(b *testing.B, supplier func() uint64, toInt32Func func(uint64) int32) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToInt32FuncCalibrated(supplier) {
		panic("Uint64ToInt32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toInt32Func(supplier())
	}
}
