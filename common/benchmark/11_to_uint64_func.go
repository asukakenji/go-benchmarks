package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToUint64FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isIntToUint64FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToUint64FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToUint64FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToUint64FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToUint64FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isUintToUint64FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToUint64FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToUint64FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToUint64FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToUint64FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setBoolToUint64FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setIntToUint64FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToUint64FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToUint64FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToUint64FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToUint64FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setUintToUint64FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToUint64FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToUint64FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToUint64FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToUint64FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Uint64, reflect.ValueOf(supplier).Pointer())
}

// BoolToUint64Func benchmarks a function with the signature:
//     func(bool) uint64
// ID: B-11-1
func BoolToUint64Func(b *testing.B, supplier func() bool, toUint64Func func(bool) uint64) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToUint64FuncCalibrated(supplier) {
		panic("BoolToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// IntToUint64Func benchmarks a function with the signature:
//     func(int) uint64
// ID: B-11-2
func IntToUint64Func(b *testing.B, supplier func() int, toUint64Func func(int) uint64) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToUint64FuncCalibrated(supplier) {
		panic("IntToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int8ToUint64Func benchmarks a function with the signature:
//     func(int8) uint64
// ID: B-11-3
func Int8ToUint64Func(b *testing.B, supplier func() int8, toUint64Func func(int8) uint64) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToUint64FuncCalibrated(supplier) {
		panic("Int8ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int16ToUint64Func benchmarks a function with the signature:
//     func(int16) uint64
// ID: B-11-4
func Int16ToUint64Func(b *testing.B, supplier func() int16, toUint64Func func(int16) uint64) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToUint64FuncCalibrated(supplier) {
		panic("Int16ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int32ToUint64Func benchmarks a function with the signature:
//     func(int32) uint64
// ID: B-11-5
func Int32ToUint64Func(b *testing.B, supplier func() int32, toUint64Func func(int32) uint64) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToUint64FuncCalibrated(supplier) {
		panic("Int32ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Int64ToUint64Func benchmarks a function with the signature:
//     func(int64) uint64
// ID: B-11-6
func Int64ToUint64Func(b *testing.B, supplier func() int64, toUint64Func func(int64) uint64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToUint64FuncCalibrated(supplier) {
		panic("Int64ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// UintToUint64Func benchmarks a function with the signature:
//     func(uint) uint64
// ID: B-11-7
func UintToUint64Func(b *testing.B, supplier func() uint, toUint64Func func(uint) uint64) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToUint64FuncCalibrated(supplier) {
		panic("UintToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint8ToUint64Func benchmarks a function with the signature:
//     func(uint8) uint64
// ID: B-11-8
func Uint8ToUint64Func(b *testing.B, supplier func() uint8, toUint64Func func(uint8) uint64) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToUint64FuncCalibrated(supplier) {
		panic("Uint8ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint16ToUint64Func benchmarks a function with the signature:
//     func(uint16) uint64
// ID: B-11-9
func Uint16ToUint64Func(b *testing.B, supplier func() uint16, toUint64Func func(uint16) uint64) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToUint64FuncCalibrated(supplier) {
		panic("Uint16ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint32ToUint64Func benchmarks a function with the signature:
//     func(uint32) uint64
// ID: B-11-10
func Uint32ToUint64Func(b *testing.B, supplier func() uint32, toUint64Func func(uint32) uint64) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToUint64FuncCalibrated(supplier) {
		panic("Uint32ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}

// Uint64ToUint64Func benchmarks a function with the signature:
//     func(uint64) uint64
// ID: B-11-11
func Uint64ToUint64Func(b *testing.B, supplier func() uint64, toUint64Func func(uint64) uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToUint64FuncCalibrated(supplier) {
		panic("Uint64ToUint64Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint64Func(supplier())
	}
}
