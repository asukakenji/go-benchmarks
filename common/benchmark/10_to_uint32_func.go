package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToUint32FuncCalibrated   = map[uintptr]struct{}{}
	_isIntToUint32FuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToUint32FuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToUint32FuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToUint32FuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToUint32FuncCalibrated  = map[uintptr]struct{}{}
	_isUintToUint32FuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToUint32FuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToUint32FuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToUint32FuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToUint32FuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToUint32FuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToUint32FuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToUint32FuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToUint32FuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToUint32FuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToUint32FuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToUint32FuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToUint32FuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToUint32FuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToUint32FuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToUint32FuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToUint32FuncCalibrated(supplier func() bool) {
	_isBoolToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToUint32FuncCalibrated(supplier func() int) {
	_isIntToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToUint32FuncCalibrated(supplier func() int8) {
	_isInt8ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToUint32FuncCalibrated(supplier func() int16) {
	_isInt16ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToUint32FuncCalibrated(supplier func() int32) {
	_isInt32ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToUint32FuncCalibrated(supplier func() int64) {
	_isInt64ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToUint32FuncCalibrated(supplier func() uint) {
	_isUintToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToUint32FuncCalibrated(supplier func() uint8) {
	_isUint8ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToUint32FuncCalibrated(supplier func() uint16) {
	_isUint16ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToUint32FuncCalibrated(supplier func() uint32) {
	_isUint32ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToUint32FuncCalibrated(supplier func() uint64) {
	_isUint64ToUint32FuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// BoolToUint32Func benchmarks a function with the signature:
//     func(bool) uint32
// ID: B-10-1
func BoolToUint32Func(b *testing.B, supplier func() bool, toUint32Func func(bool) uint32) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToUint32FuncCalibrated(supplier) {
		panic("BoolToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// IntToUint32Func benchmarks a function with the signature:
//     func(int) uint32
// ID: B-10-2
func IntToUint32Func(b *testing.B, supplier func() int, toUint32Func func(int) uint32) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToUint32FuncCalibrated(supplier) {
		panic("IntToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int8ToUint32Func benchmarks a function with the signature:
//     func(int8) uint32
// ID: B-10-3
func Int8ToUint32Func(b *testing.B, supplier func() int8, toUint32Func func(int8) uint32) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToUint32FuncCalibrated(supplier) {
		panic("Int8ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int16ToUint32Func benchmarks a function with the signature:
//     func(int16) uint32
// ID: B-10-4
func Int16ToUint32Func(b *testing.B, supplier func() int16, toUint32Func func(int16) uint32) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToUint32FuncCalibrated(supplier) {
		panic("Int16ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int32ToUint32Func benchmarks a function with the signature:
//     func(int32) uint32
// ID: B-10-5
func Int32ToUint32Func(b *testing.B, supplier func() int32, toUint32Func func(int32) uint32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToUint32FuncCalibrated(supplier) {
		panic("Int32ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Int64ToUint32Func benchmarks a function with the signature:
//     func(int64) uint32
// ID: B-10-6
func Int64ToUint32Func(b *testing.B, supplier func() int64, toUint32Func func(int64) uint32) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToUint32FuncCalibrated(supplier) {
		panic("Int64ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// UintToUint32Func benchmarks a function with the signature:
//     func(uint) uint32
// ID: B-10-7
func UintToUint32Func(b *testing.B, supplier func() uint, toUint32Func func(uint) uint32) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToUint32FuncCalibrated(supplier) {
		panic("UintToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint8ToUint32Func benchmarks a function with the signature:
//     func(uint8) uint32
// ID: B-10-8
func Uint8ToUint32Func(b *testing.B, supplier func() uint8, toUint32Func func(uint8) uint32) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToUint32FuncCalibrated(supplier) {
		panic("Uint8ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint16ToUint32Func benchmarks a function with the signature:
//     func(uint16) uint32
// ID: B-10-9
func Uint16ToUint32Func(b *testing.B, supplier func() uint16, toUint32Func func(uint16) uint32) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToUint32FuncCalibrated(supplier) {
		panic("Uint16ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint32ToUint32Func benchmarks a function with the signature:
//     func(uint32) uint32
// ID: B-10-10
func Uint32ToUint32Func(b *testing.B, supplier func() uint32, toUint32Func func(uint32) uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToUint32FuncCalibrated(supplier) {
		panic("Uint32ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}

// Uint64ToUint32Func benchmarks a function with the signature:
//     func(uint32) uint32
// ID: B-10-11
func Uint64ToUint32Func(b *testing.B, supplier func() uint64, toUint32Func func(uint64) uint32) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToUint32FuncCalibrated(supplier) {
		panic("Uint64ToUint32Func not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toUint32Func(supplier())
	}
}
