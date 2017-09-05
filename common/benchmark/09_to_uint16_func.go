package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToUint16FuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isIntToUint16FuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToUint16FuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToUint16FuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToUint16FuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToUint16FuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUintToUint16FuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToUint16FuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToUint16FuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToUint16FuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToUint16FuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setBoolToUint16FuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setIntToUint16FuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToUint16FuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToUint16FuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToUint16FuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToUint16FuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setUintToUint16FuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToUint16FuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToUint16FuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToUint16FuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToUint16FuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Uint16, reflect.ValueOf(supplier).Pointer())
}

// BoolToUint16Func benchmarks a function with the signature:
//     func(bool) uint16
// ID: B-9-1
func BoolToUint16Func(b *testing.B, supplier func() bool, toUint16Func func(bool) uint16) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToUint16FuncCalibrated(supplier) {
		panic("BoolToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// IntToUint16Func benchmarks a function with the signature:
//     func(int) uint16
// ID: B-9-2
func IntToUint16Func(b *testing.B, supplier func() int, toUint16Func func(int) uint16) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToUint16FuncCalibrated(supplier) {
		panic("IntToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int8ToUint16Func benchmarks a function with the signature:
//     func(int8) uint16
// ID: B-9-3
func Int8ToUint16Func(b *testing.B, supplier func() int8, toUint16Func func(int8) uint16) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToUint16FuncCalibrated(supplier) {
		panic("Int8ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int16ToUint16Func benchmarks a function with the signature:
//     func(int16) uint16
// ID: B-9-4
func Int16ToUint16Func(b *testing.B, supplier func() int16, toUint16Func func(int16) uint16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToUint16FuncCalibrated(supplier) {
		panic("Int16ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int32ToUint16Func benchmarks a function with the signature:
//     func(int32) uint16
// ID: B-9-5
func Int32ToUint16Func(b *testing.B, supplier func() int32, toUint16Func func(int32) uint16) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToUint16FuncCalibrated(supplier) {
		panic("Int32ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Int64ToUint16Func benchmarks a function with the signature:
//     func(int64) uint16
// ID: B-9-6
func Int64ToUint16Func(b *testing.B, supplier func() int64, toUint16Func func(int64) uint16) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToUint16FuncCalibrated(supplier) {
		panic("Int64ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// UintToUint16Func benchmarks a function with the signature:
//     func(uint) uint16
// ID: B-9-7
func UintToUint16Func(b *testing.B, supplier func() uint, toUint16Func func(uint) uint16) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToUint16FuncCalibrated(supplier) {
		panic("UintToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint8ToUint16Func benchmarks a function with the signature:
//     func(uint8) uint16
// ID: B-9-8
func Uint8ToUint16Func(b *testing.B, supplier func() uint8, toUint16Func func(uint8) uint16) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToUint16FuncCalibrated(supplier) {
		panic("Uint8ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint16ToUint16Func benchmarks a function with the signature:
//     func(uint16) uint16
// ID: B-9-9
func Uint16ToUint16Func(b *testing.B, supplier func() uint16, toUint16Func func(uint16) uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToUint16FuncCalibrated(supplier) {
		panic("Uint16ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint32ToUint16Func benchmarks a function with the signature:
//     func(uint32) uint16
// ID: B-9-10
func Uint32ToUint16Func(b *testing.B, supplier func() uint32, toUint16Func func(uint32) uint16) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToUint16FuncCalibrated(supplier) {
		panic("Uint32ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}

// Uint64ToUint16Func benchmarks a function with the signature:
//     func(uint16) uint16
// ID: B-9-11
func Uint64ToUint16Func(b *testing.B, supplier func() uint64, toUint16Func func(uint64) uint16) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToUint16FuncCalibrated(supplier) {
		panic("Uint64ToUint16Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint16Func(supplier())
	}
}
