package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToUint8FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToUint8FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToUint8FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToUint8FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToUint8FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToUint8FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToUint8FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToUint8FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToUint8FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToUint8FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToUint8FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToUint8FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToUint8FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToUint8FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToUint8FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToUint8FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToUint8FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToUint8FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToUint8FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToUint8FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToUint8FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToUint8FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToUint8FuncCalibrated(supplier func() bool) {
	_isBoolToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToUint8FuncCalibrated(supplier func() int) {
	_isIntToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToUint8FuncCalibrated(supplier func() int8) {
	_isInt8ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToUint8FuncCalibrated(supplier func() int16) {
	_isInt16ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToUint8FuncCalibrated(supplier func() int32) {
	_isInt32ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToUint8FuncCalibrated(supplier func() int64) {
	_isInt64ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToUint8FuncCalibrated(supplier func() uint) {
	_isUintToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToUint8FuncCalibrated(supplier func() uint8) {
	_isUint8ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToUint8FuncCalibrated(supplier func() uint16) {
	_isUint16ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToUint8FuncCalibrated(supplier func() uint32) {
	_isUint32ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToUint8FuncCalibrated(supplier func() uint64) {
	_isUint64ToUint8FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// BoolToUint8Func benchmarks a function with the signature:
//     func(bool) uint8
// ID: B-8-1
func BoolToUint8Func(b *testing.B, supplier func() bool, toUint8Func func(bool) uint8) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToUint8FuncCalibrated(supplier) {
		panic("BoolToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// IntToUint8Func benchmarks a function with the signature:
//     func(int) uint8
// ID: B-8-2
func IntToUint8Func(b *testing.B, supplier func() int, toUint8Func func(int) uint8) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToUint8FuncCalibrated(supplier) {
		panic("IntToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int8ToUint8Func benchmarks a function with the signature:
//     func(int8) uint8
// ID: B-8-3
func Int8ToUint8Func(b *testing.B, supplier func() int8, toUint8Func func(int8) uint8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToUint8FuncCalibrated(supplier) {
		panic("Int8ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int16ToUint8Func benchmarks a function with the signature:
//     func(int16) uint8
// ID: B-8-4
func Int16ToUint8Func(b *testing.B, supplier func() int16, toUint8Func func(int16) uint8) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToUint8FuncCalibrated(supplier) {
		panic("Int16ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int32ToUint8Func benchmarks a function with the signature:
//     func(int32) uint8
// ID: B-8-5
func Int32ToUint8Func(b *testing.B, supplier func() int32, toUint8Func func(int32) uint8) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToUint8FuncCalibrated(supplier) {
		panic("Int32ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Int64ToUint8Func benchmarks a function with the signature:
//     func(int64) uint8
// ID: B-8-6
func Int64ToUint8Func(b *testing.B, supplier func() int64, toUint8Func func(int64) uint8) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToUint8FuncCalibrated(supplier) {
		panic("Int64ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// UintToUint8Func benchmarks a function with the signature:
//     func(uint) uint8
// ID: B-8-7
func UintToUint8Func(b *testing.B, supplier func() uint, toUint8Func func(uint) uint8) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToUint8FuncCalibrated(supplier) {
		panic("UintToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint8ToUint8Func benchmarks a function with the signature:
//     func(uint8) uint8
// ID: B-8-8
func Uint8ToUint8Func(b *testing.B, supplier func() uint8, toUint8Func func(uint8) uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToUint8FuncCalibrated(supplier) {
		panic("Uint8ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint16ToUint8Func benchmarks a function with the signature:
//     func(uint16) uint8
// ID: B-8-9
func Uint16ToUint8Func(b *testing.B, supplier func() uint16, toUint8Func func(uint16) uint8) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToUint8FuncCalibrated(supplier) {
		panic("Uint16ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint32ToUint8Func benchmarks a function with the signature:
//     func(uint32) uint8
// ID: B-8-10
func Uint32ToUint8Func(b *testing.B, supplier func() uint32, toUint8Func func(uint32) uint8) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToUint8FuncCalibrated(supplier) {
		panic("Uint32ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}

// Uint64ToUint8Func benchmarks a function with the signature:
//     func(uint8) uint8
// ID: B-8-11
func Uint64ToUint8Func(b *testing.B, supplier func() uint64, toUint8Func func(uint64) uint8) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToUint8FuncCalibrated(supplier) {
		panic("Uint64ToUint8Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint8Func(supplier())
	}
}
