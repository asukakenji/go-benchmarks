package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToInt8FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToInt8FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToInt8FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToInt8FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToInt8FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToInt8FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToInt8FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToInt8FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToInt8FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToInt8FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToInt8FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToInt8FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToInt8FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToInt8FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToInt8FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToInt8FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToInt8FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToInt8FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToInt8FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToInt8FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToInt8FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToInt8FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToInt8FuncCalibrated(supplier func() bool) {
	_isBoolToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToInt8FuncCalibrated(supplier func() int) {
	_isIntToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToInt8FuncCalibrated(supplier func() int8) {
	_isInt8ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToInt8FuncCalibrated(supplier func() int16) {
	_isInt16ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToInt8FuncCalibrated(supplier func() int32) {
	_isInt32ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToInt8FuncCalibrated(supplier func() int64) {
	_isInt64ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToInt8FuncCalibrated(supplier func() uint) {
	_isUintToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToInt8FuncCalibrated(supplier func() uint8) {
	_isUint8ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToInt8FuncCalibrated(supplier func() uint16) {
	_isUint16ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToInt8FuncCalibrated(supplier func() uint32) {
	_isUint32ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToInt8FuncCalibrated(supplier func() uint64) {
	_isUint64ToInt8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
