package benchmark

import (
	"reflect"
	"testing"
)

func isBoolToUintFuncCalibrated(supplier func() bool) bool {
	return isCalibrated(reflect.Bool, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isIntToUintFuncCalibrated(supplier func() int) bool {
	return isCalibrated(reflect.Int, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isInt8ToUintFuncCalibrated(supplier func() int8) bool {
	return isCalibrated(reflect.Int8, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isInt16ToUintFuncCalibrated(supplier func() int16) bool {
	return isCalibrated(reflect.Int16, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isInt32ToUintFuncCalibrated(supplier func() int32) bool {
	return isCalibrated(reflect.Int32, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isInt64ToUintFuncCalibrated(supplier func() int64) bool {
	return isCalibrated(reflect.Int64, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUintToUintFuncCalibrated(supplier func() uint) bool {
	return isCalibrated(reflect.Uint, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUint8ToUintFuncCalibrated(supplier func() uint8) bool {
	return isCalibrated(reflect.Uint8, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUint16ToUintFuncCalibrated(supplier func() uint16) bool {
	return isCalibrated(reflect.Uint16, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUint32ToUintFuncCalibrated(supplier func() uint32) bool {
	return isCalibrated(reflect.Uint32, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func isUint64ToUintFuncCalibrated(supplier func() uint64) bool {
	return isCalibrated(reflect.Uint64, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setBoolToUintFuncCalibrated(supplier func() bool) {
	setCalibrated(reflect.Bool, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setIntToUintFuncCalibrated(supplier func() int) {
	setCalibrated(reflect.Int, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setInt8ToUintFuncCalibrated(supplier func() int8) {
	setCalibrated(reflect.Int8, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setInt16ToUintFuncCalibrated(supplier func() int16) {
	setCalibrated(reflect.Int16, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setInt32ToUintFuncCalibrated(supplier func() int32) {
	setCalibrated(reflect.Int32, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setInt64ToUintFuncCalibrated(supplier func() int64) {
	setCalibrated(reflect.Int64, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setUintToUintFuncCalibrated(supplier func() uint) {
	setCalibrated(reflect.Uint, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setUint8ToUintFuncCalibrated(supplier func() uint8) {
	setCalibrated(reflect.Uint8, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setUint16ToUintFuncCalibrated(supplier func() uint16) {
	setCalibrated(reflect.Uint16, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setUint32ToUintFuncCalibrated(supplier func() uint32) {
	setCalibrated(reflect.Uint32, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

func setUint64ToUintFuncCalibrated(supplier func() uint64) {
	setCalibrated(reflect.Uint64, reflect.Uint, reflect.ValueOf(supplier).Pointer())
}

// BoolToUintFunc benchmarks a function with the signature:
//     func(bool) uint
// ID: B-7-1
func BoolToUintFunc(b *testing.B, supplier func() bool, toUintFunc func(bool) uint) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToUintFuncCalibrated(supplier) {
		panic("BoolToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// IntToUintFunc benchmarks a function with the signature:
//     func(int) uint
// ID: B-7-2
func IntToUintFunc(b *testing.B, supplier func() int, toUintFunc func(int) uint) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToUintFuncCalibrated(supplier) {
		panic("IntToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int8ToUintFunc benchmarks a function with the signature:
//     func(int8) uint
// ID: B-7-3
func Int8ToUintFunc(b *testing.B, supplier func() int8, toUintFunc func(int8) uint) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToUintFuncCalibrated(supplier) {
		panic("Int8ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int16ToUintFunc benchmarks a function with the signature:
//     func(int16) uint
// ID: B-7-4
func Int16ToUintFunc(b *testing.B, supplier func() int16, toUintFunc func(int16) uint) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToUintFuncCalibrated(supplier) {
		panic("Int16ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int32ToUintFunc benchmarks a function with the signature:
//     func(int32) uint
// ID: B-7-5
func Int32ToUintFunc(b *testing.B, supplier func() int32, toUintFunc func(int32) uint) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToUintFuncCalibrated(supplier) {
		panic("Int32ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Int64ToUintFunc benchmarks a function with the signature:
//     func(int64) uint
// ID: B-7-6
func Int64ToUintFunc(b *testing.B, supplier func() int64, toUintFunc func(int64) uint) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToUintFuncCalibrated(supplier) {
		panic("Int64ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// UintToUintFunc benchmarks a function with the signature:
//     func(uint) uint
// ID: B-7-7
func UintToUintFunc(b *testing.B, supplier func() uint, toUintFunc func(uint) uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToUintFuncCalibrated(supplier) {
		panic("UintToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint8ToUintFunc benchmarks a function with the signature:
//     func(uint8) uint
// ID: B-7-8
func Uint8ToUintFunc(b *testing.B, supplier func() uint8, toUintFunc func(uint8) uint) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToUintFuncCalibrated(supplier) {
		panic("Uint8ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint16ToUintFunc benchmarks a function with the signature:
//     func(uint16) uint
// ID: B-7-9
func Uint16ToUintFunc(b *testing.B, supplier func() uint16, toUintFunc func(uint16) uint) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToUintFuncCalibrated(supplier) {
		panic("Uint16ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint32ToUintFunc benchmarks a function with the signature:
//     func(uint32) uint
// ID: B-7-10
func Uint32ToUintFunc(b *testing.B, supplier func() uint32, toUintFunc func(uint32) uint) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToUintFuncCalibrated(supplier) {
		panic("Uint32ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}

// Uint64ToUintFunc benchmarks a function with the signature:
//     func(uint) uint
// ID: B-7-11
func Uint64ToUintFunc(b *testing.B, supplier func() uint64, toUintFunc func(uint64) uint) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToUintFuncCalibrated(supplier) {
		panic("Uint64ToUintFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUintFunc(supplier())
	}
}
