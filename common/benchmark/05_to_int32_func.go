package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToInt32FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToInt32FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToInt32FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToInt32FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToInt32FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToInt32FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToInt32FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToInt32FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToInt32FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToInt32FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToInt32FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToInt32FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToInt32FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToInt32FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToInt32FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToInt32FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToInt32FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToInt32FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToInt32FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToInt32FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToInt32FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToInt32FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToInt32FuncCalibrated(supplier func() bool) {
	_isBoolToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToInt32FuncCalibrated(supplier func() int) {
	_isIntToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToInt32FuncCalibrated(supplier func() int8) {
	_isInt8ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToInt32FuncCalibrated(supplier func() int16) {
	_isInt16ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToInt32FuncCalibrated(supplier func() int32) {
	_isInt32ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToInt32FuncCalibrated(supplier func() int64) {
	_isInt64ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToInt32FuncCalibrated(supplier func() uint) {
	_isUintToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToInt32FuncCalibrated(supplier func() uint8) {
	_isUint8ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToInt32FuncCalibrated(supplier func() uint16) {
	_isUint16ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToInt32FuncCalibrated(supplier func() uint32) {
	_isUint32ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToInt32FuncCalibrated(supplier func() uint64) {
	_isUint64ToInt32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
