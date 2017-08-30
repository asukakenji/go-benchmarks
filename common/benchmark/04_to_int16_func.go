package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToInt16FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToInt16FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToInt16FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToInt16FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToInt16FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToInt16FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToInt16FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToInt16FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToInt16FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToInt16FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToInt16FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToInt16FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToInt16FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToInt16FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToInt16FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToInt16FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToInt16FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToInt16FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToInt16FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToInt16FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToInt16FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToInt16FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToInt16FuncCalibrated(supplier func() bool) {
	_isBoolToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToInt16FuncCalibrated(supplier func() int) {
	_isIntToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToInt16FuncCalibrated(supplier func() int8) {
	_isInt8ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToInt16FuncCalibrated(supplier func() int16) {
	_isInt16ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToInt16FuncCalibrated(supplier func() int32) {
	_isInt32ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToInt16FuncCalibrated(supplier func() int64) {
	_isInt64ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToInt16FuncCalibrated(supplier func() uint) {
	_isUintToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToInt16FuncCalibrated(supplier func() uint8) {
	_isUint8ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToInt16FuncCalibrated(supplier func() uint16) {
	_isUint16ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToInt16FuncCalibrated(supplier func() uint32) {
	_isUint32ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToInt16FuncCalibrated(supplier func() uint64) {
	_isUint64ToInt16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
