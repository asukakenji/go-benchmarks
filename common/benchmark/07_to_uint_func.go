package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToUintFuncCalibrated   = map[uintptr]struct{}{}
	_isIntToUintFuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToUintFuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToUintFuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToUintFuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToUintFuncCalibrated  = map[uintptr]struct{}{}
	_isUintToUintFuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToUintFuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToUintFuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToUintFuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToUintFuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToUintFuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToUintFuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToUintFuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToUintFuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToUintFuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToUintFuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToUintFuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToUintFuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToUintFuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToUintFuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToUintFuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToUintFuncCalibrated(supplier func() bool) {
	_isBoolToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToUintFuncCalibrated(supplier func() int) {
	_isIntToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToUintFuncCalibrated(supplier func() int8) {
	_isInt8ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToUintFuncCalibrated(supplier func() int16) {
	_isInt16ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToUintFuncCalibrated(supplier func() int32) {
	_isInt32ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToUintFuncCalibrated(supplier func() int64) {
	_isInt64ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToUintFuncCalibrated(supplier func() uint) {
	_isUintToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToUintFuncCalibrated(supplier func() uint8) {
	_isUint8ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToUintFuncCalibrated(supplier func() uint16) {
	_isUint16ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToUintFuncCalibrated(supplier func() uint32) {
	_isUint32ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToUintFuncCalibrated(supplier func() uint64) {
	_isUint64ToUintFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
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
