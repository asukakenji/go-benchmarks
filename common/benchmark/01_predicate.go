package benchmark

import (
	"reflect"
	"testing"
)

var (
	_isBoolPredicateCalibrated   = map[uintptr]struct{}{}
	_isIntPredicateCalibrated    = map[uintptr]struct{}{}
	_isInt8PredicateCalibrated   = map[uintptr]struct{}{}
	_isInt16PredicateCalibrated  = map[uintptr]struct{}{}
	_isInt32PredicateCalibrated  = map[uintptr]struct{}{}
	_isInt64PredicateCalibrated  = map[uintptr]struct{}{}
	_isUintPredicateCalibrated   = map[uintptr]struct{}{}
	_isUint8PredicateCalibrated  = map[uintptr]struct{}{}
	_isUint16PredicateCalibrated = map[uintptr]struct{}{}
	_isUint32PredicateCalibrated = map[uintptr]struct{}{}
	_isUint64PredicateCalibrated = map[uintptr]struct{}{}
)

func isBoolPredicateCalibrated(supplier func() bool) bool {
	_, ok := _isBoolPredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isIntPredicateCalibrated(supplier func() int) bool {
	_, ok := _isIntPredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt8PredicateCalibrated(supplier func() int8) bool {
	_, ok := _isInt8PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt16PredicateCalibrated(supplier func() int16) bool {
	_, ok := _isInt16PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt32PredicateCalibrated(supplier func() int32) bool {
	_, ok := _isInt32PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isInt64PredicateCalibrated(supplier func() int64) bool {
	_, ok := _isInt64PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUintPredicateCalibrated(supplier func() uint) bool {
	_, ok := _isUintPredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint8PredicateCalibrated(supplier func() uint8) bool {
	_, ok := _isUint8PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint16PredicateCalibrated(supplier func() uint16) bool {
	_, ok := _isUint16PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint32PredicateCalibrated(supplier func() uint32) bool {
	_, ok := _isUint32PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func isUint64PredicateCalibrated(supplier func() uint64) bool {
	_, ok := _isUint64PredicateCalibrated[reflect.ValueOf(supplier).Pointer()]
	return ok
}

func setBoolPredicateCalibrated(supplier func() bool) {
	_isBoolPredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setIntPredicateCalibrated(supplier func() int) {
	_isIntPredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt8PredicateCalibrated(supplier func() int8) {
	_isInt8PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt16PredicateCalibrated(supplier func() int16) {
	_isInt16PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt32PredicateCalibrated(supplier func() int32) {
	_isInt32PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setInt64PredicateCalibrated(supplier func() int64) {
	_isInt64PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUintPredicateCalibrated(supplier func() uint) {
	_isUintPredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint8PredicateCalibrated(supplier func() uint8) {
	_isUint8PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint16PredicateCalibrated(supplier func() uint16) {
	_isUint16PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint32PredicateCalibrated(supplier func() uint32) {
	_isUint32PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

func setUint64PredicateCalibrated(supplier func() uint64) {
	_isUint64PredicateCalibrated[reflect.ValueOf(supplier).Pointer()] = struct{}{}
}

// BoolPredicate benchmarks a function with the signature:
//     func(bool) bool
// ID: B-1-1
func BoolPredicate(b *testing.B, supplier func() bool, predicate func(bool) bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isBoolPredicateCalibrated(supplier) {
		panic("BoolPredicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// IntPredicate benchmarks a function with the signature:
//     func(int) bool
// ID: B-1-2
func IntPredicate(b *testing.B, supplier func() int, predicate func(int) bool) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isIntPredicateCalibrated(supplier) {
		panic("IntPredicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int8Predicate benchmarks a function with the signature:
//     func(int8) bool
// ID: B-1-3
func Int8Predicate(b *testing.B, supplier func() int8, predicate func(int8) bool) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt8PredicateCalibrated(supplier) {
		panic("Int8Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int16Predicate benchmarks a function with the signature:
//     func(int16) bool
// ID: B-1-4
func Int16Predicate(b *testing.B, supplier func() int16, predicate func(int16) bool) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt16PredicateCalibrated(supplier) {
		panic("Int16Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int32Predicate benchmarks a function with the signature:
//     func(int32) bool
// ID: B-1-5
func Int32Predicate(b *testing.B, supplier func() int32, predicate func(int32) bool) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt32PredicateCalibrated(supplier) {
		panic("Int32Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Int64Predicate benchmarks a function with the signature:
//     func(int64) bool
// ID: B-1-6
func Int64Predicate(b *testing.B, supplier func() int64, predicate func(int64) bool) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isInt64PredicateCalibrated(supplier) {
		panic("Int64Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// UintPredicate benchmarks a function with the signature:
//     func(uint) bool
// ID: B-1-7
func UintPredicate(b *testing.B, supplier func() uint, predicate func(uint) bool) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUintPredicateCalibrated(supplier) {
		panic("UintPredicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint8Predicate benchmarks a function with the signature:
//     func(uint8) bool
// ID: B-1-8
func Uint8Predicate(b *testing.B, supplier func() uint8, predicate func(uint8) bool) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint8PredicateCalibrated(supplier) {
		panic("Uint8Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint16Predicate benchmarks a function with the signature:
//     func(uint16) bool
// ID: B-1-9
func Uint16Predicate(b *testing.B, supplier func() uint16, predicate func(uint16) bool) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint16PredicateCalibrated(supplier) {
		panic("Uint16Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint32Predicate benchmarks a function with the signature:
//     func(uint32) bool
// ID: B-1-10
func Uint32Predicate(b *testing.B, supplier func() uint32, predicate func(uint32) bool) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint32PredicateCalibrated(supplier) {
		panic("Uint32Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}

// Uint64Predicate benchmarks a function with the signature:
//     func(uint64) bool
// ID: B-1-11
func Uint64Predicate(b *testing.B, supplier func() uint64, predicate func(uint64) bool) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	if !isUint64PredicateCalibrated(supplier) {
		panic("Uint64Predicate not calibrated with this supplier")
	}
	for i, count := 0, b.N; i < count; i++ {
		predicate(supplier())
	}
}
