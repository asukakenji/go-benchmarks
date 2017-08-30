package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToUint16FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToUint16FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToUint16FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToUint16FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToUint16FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToUint16FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToUint16FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToUint16FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToUint16FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToUint16FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToUint16FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToUint16FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToUint16FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToUint16FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToUint16FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToUint16FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToUint16FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToUint16FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToUint16FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToUint16FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToUint16FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToUint16FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToUint16FuncCalibrated(supplier func() bool) {
	_isBoolToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToUint16FuncCalibrated(supplier func() int) {
	_isIntToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToUint16FuncCalibrated(supplier func() int8) {
	_isInt8ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToUint16FuncCalibrated(supplier func() int16) {
	_isInt16ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToUint16FuncCalibrated(supplier func() int32) {
	_isInt32ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToUint16FuncCalibrated(supplier func() int64) {
	_isInt64ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToUint16FuncCalibrated(supplier func() uint) {
	_isUintToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToUint16FuncCalibrated(supplier func() uint8) {
	_isUint8ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToUint16FuncCalibrated(supplier func() uint16) {
	_isUint16ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToUint16FuncCalibrated(supplier func() uint32) {
	_isUint32ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToUint16FuncCalibrated(supplier func() uint64) {
	_isUint64ToUint16FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
