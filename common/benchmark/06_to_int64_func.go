package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToInt64FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToInt64FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToInt64FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToInt64FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToInt64FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToInt64FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToInt64FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToInt64FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToInt64FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToInt64FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToInt64FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToInt64FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToInt64FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToInt64FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToInt64FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToInt64FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToInt64FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToInt64FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToInt64FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToInt64FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToInt64FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToInt64FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToInt64FuncCalibrated(supplier func() bool) {
	_isBoolToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToInt64FuncCalibrated(supplier func() int) {
	_isIntToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToInt64FuncCalibrated(supplier func() int8) {
	_isInt8ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToInt64FuncCalibrated(supplier func() int16) {
	_isInt16ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToInt64FuncCalibrated(supplier func() int32) {
	_isInt32ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToInt64FuncCalibrated(supplier func() int64) {
	_isInt64ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToInt64FuncCalibrated(supplier func() uint) {
	_isUintToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToInt64FuncCalibrated(supplier func() uint8) {
	_isUint8ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToInt64FuncCalibrated(supplier func() uint16) {
	_isUint16ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToInt64FuncCalibrated(supplier func() uint32) {
	_isUint32ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToInt64FuncCalibrated(supplier func() uint64) {
	_isUint64ToInt64FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
