package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolToIntFuncCalibrated   = map[uintptr]struct{}{}
	_isIntToIntFuncCalibrated    = map[uintptr]struct{}{}
	_isInt8ToIntFuncCalibrated   = map[uintptr]struct{}{}
	_isInt16ToIntFuncCalibrated  = map[uintptr]struct{}{}
	_isInt32ToIntFuncCalibrated  = map[uintptr]struct{}{}
	_isInt64ToIntFuncCalibrated  = map[uintptr]struct{}{}
	_isUintToIntFuncCalibrated   = map[uintptr]struct{}{}
	_isUint8ToIntFuncCalibrated  = map[uintptr]struct{}{}
	_isUint16ToIntFuncCalibrated = map[uintptr]struct{}{}
	_isUint32ToIntFuncCalibrated = map[uintptr]struct{}{}
	_isUint64ToIntFuncCalibrated = map[uintptr]struct{}{}
)

func isBoolToIntFuncCalibrated(supplier func() bool) bool {
	_, ok := _isBoolToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntToIntFuncCalibrated(supplier func() int) bool {
	_, ok := _isIntToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8ToIntFuncCalibrated(supplier func() int8) bool {
	_, ok := _isInt8ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16ToIntFuncCalibrated(supplier func() int16) bool {
	_, ok := _isInt16ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32ToIntFuncCalibrated(supplier func() int32) bool {
	_, ok := _isInt32ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64ToIntFuncCalibrated(supplier func() int64) bool {
	_, ok := _isInt64ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintToIntFuncCalibrated(supplier func() uint) bool {
	_, ok := _isUintToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8ToIntFuncCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16ToIntFuncCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32ToIntFuncCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64ToIntFuncCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolToIntFuncCalibrated(supplier func() bool) {
	_isBoolToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntToIntFuncCalibrated(supplier func() int) {
	_isIntToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8ToIntFuncCalibrated(supplier func() int8) {
	_isInt8ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16ToIntFuncCalibrated(supplier func() int16) {
	_isInt16ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32ToIntFuncCalibrated(supplier func() int32) {
	_isInt32ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64ToIntFuncCalibrated(supplier func() int64) {
	_isInt64ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintToIntFuncCalibrated(supplier func() uint) {
	_isUintToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8ToIntFuncCalibrated(supplier func() uint8) {
	_isUint8ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16ToIntFuncCalibrated(supplier func() uint16) {
	_isUint16ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32ToIntFuncCalibrated(supplier func() uint32) {
	_isUint32ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64ToIntFuncCalibrated(supplier func() uint64) {
	_isUint64ToIntFuncCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// BoolToIntFunc benchmarks a function with the signature:
//     func(bool) int
// ID: B-2-1
func BoolToIntFunc(b *testing.B, supplier func() bool, toIntFunc func(bool) int) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolToIntFuncCalibrated(supplier) {
		panic("BoolToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// IntToIntFunc benchmarks a function with the signature:
//     func(int) int
// ID: B-2-2
func IntToIntFunc(b *testing.B, supplier func() int, toIntFunc func(int) int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntToIntFuncCalibrated(supplier) {
		panic("IntToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int8ToIntFunc benchmarks a function with the signature:
//     func(int8) int
// ID: B-2-3
func Int8ToIntFunc(b *testing.B, supplier func() int8, toIntFunc func(int8) int) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8ToIntFuncCalibrated(supplier) {
		panic("Int8ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int16ToIntFunc benchmarks a function with the signature:
//     func(int16) int
// ID: B-2-4
func Int16ToIntFunc(b *testing.B, supplier func() int16, toIntFunc func(int16) int) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16ToIntFuncCalibrated(supplier) {
		panic("Int16ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int32ToIntFunc benchmarks a function with the signature:
//     func(int32) int
// ID: B-2-5
func Int32ToIntFunc(b *testing.B, supplier func() int32, toIntFunc func(int32) int) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32ToIntFuncCalibrated(supplier) {
		panic("Int32ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Int64ToIntFunc benchmarks a function with the signature:
//     func(int64) int
// ID: B-2-6
func Int64ToIntFunc(b *testing.B, supplier func() int64, toIntFunc func(int64) int) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64ToIntFuncCalibrated(supplier) {
		panic("Int64ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// UintToIntFunc benchmarks a function with the signature:
//     func(uint) int
// ID: B-2-7
func UintToIntFunc(b *testing.B, supplier func() uint, toIntFunc func(uint) int) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintToIntFuncCalibrated(supplier) {
		panic("UintToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint8ToIntFunc benchmarks a function with the signature:
//     func(uint8) int
// ID: B-2-8
func Uint8ToIntFunc(b *testing.B, supplier func() uint8, toIntFunc func(uint8) int) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8ToIntFuncCalibrated(supplier) {
		panic("Uint8ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint16ToIntFunc benchmarks a function with the signature:
//     func(uint16) int
// ID: B-2-9
func Uint16ToIntFunc(b *testing.B, supplier func() uint16, toIntFunc func(uint16) int) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16ToIntFuncCalibrated(supplier) {
		panic("Uint16ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint32ToIntFunc benchmarks a function with the signature:
//     func(uint32) int
// ID: B-2-10
func Uint32ToIntFunc(b *testing.B, supplier func() uint32, toIntFunc func(uint32) int) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32ToIntFuncCalibrated(supplier) {
		panic("Uint32ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}

// Uint64ToIntFunc benchmarks a function with the signature:
//     func(int) int
// ID: B-2-11
func Uint64ToIntFunc(b *testing.B, supplier func() uint64, toIntFunc func(uint64) int) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64ToIntFuncCalibrated(supplier) {
		panic("Uint64ToIntFunc not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		toIntFunc(supplier())
	}
}
