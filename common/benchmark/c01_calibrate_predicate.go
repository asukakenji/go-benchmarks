package benchmark

import "testing"

func dummyBoolPredicate(_ bool) bool     { return false }
func dummyIntPredicate(_ int) bool       { return false }
func dummyInt8Predicate(_ int8) bool     { return false }
func dummyInt16Predicate(_ int16) bool   { return false }
func dummyInt32Predicate(_ int32) bool   { return false }
func dummyInt64Predicate(_ int64) bool   { return false }
func dummyUintPredicate(_ uint) bool     { return false }
func dummyUint8Predicate(_ uint8) bool   { return false }
func dummyUint16Predicate(_ uint16) bool { return false }
func dummyUint32Predicate(_ uint32) bool { return false }
func dummyUint64Predicate(_ uint64) bool { return false }

// CalibrateBoolPredicate calibrates the BoolPredicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolPredicate.
//
// ID: BC-1-1
func CalibrateBoolPredicate(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolPredicateCalibrated(supplier)
	BoolPredicate(b, supplier, dummyBoolPredicate)
}

// CalibrateIntPredicate calibrates the IntPredicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntPredicate.
//
// ID: BC-1-2
func CalibrateIntPredicate(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntPredicateCalibrated(supplier)
	IntPredicate(b, supplier, dummyIntPredicate)
}

// CalibrateInt8Predicate calibrates the Int8Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8Predicate.
//
// ID: BC-1-3
func CalibrateInt8Predicate(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8PredicateCalibrated(supplier)
	Int8Predicate(b, supplier, dummyInt8Predicate)
}

// CalibrateInt16Predicate calibrates the Int16Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16Predicate.
//
// ID: BC-1-4
func CalibrateInt16Predicate(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16PredicateCalibrated(supplier)
	Int16Predicate(b, supplier, dummyInt16Predicate)
}

// CalibrateInt32Predicate calibrates the Int32Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32Predicate.
//
// ID: BC-1-5
func CalibrateInt32Predicate(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32PredicateCalibrated(supplier)
	Int32Predicate(b, supplier, dummyInt32Predicate)
}

// CalibrateInt64Predicate calibrates the Int64Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64Predicate.
//
// ID: BC-1-6
func CalibrateInt64Predicate(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64PredicateCalibrated(supplier)
	Int64Predicate(b, supplier, dummyInt64Predicate)
}

// CalibrateUintPredicate calibrates the UintPredicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintPredicate.
//
// ID: BC-1-7
func CalibrateUintPredicate(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintPredicateCalibrated(supplier)
	UintPredicate(b, supplier, dummyUintPredicate)
}

// CalibrateUint8Predicate calibrates the Uint8Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8Predicate.
//
// ID: BC-1-8
func CalibrateUint8Predicate(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8PredicateCalibrated(supplier)
	Uint8Predicate(b, supplier, dummyUint8Predicate)
}

// CalibrateUint16Predicate calibrates the Uint16Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16Predicate.
//
// ID: BC-1-9
func CalibrateUint16Predicate(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16PredicateCalibrated(supplier)
	Uint16Predicate(b, supplier, dummyUint16Predicate)
}

// CalibrateUint32Predicate calibrates the Uint32Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32Predicate.
//
// ID: BC-1-10
func CalibrateUint32Predicate(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32PredicateCalibrated(supplier)
	Uint32Predicate(b, supplier, dummyUint32Predicate)
}

// CalibrateUint64Predicate calibrates the Uint64Predicate function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64Predicate.
//
// ID: BC-1-11
func CalibrateUint64Predicate(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64PredicateCalibrated(supplier)
	Uint64Predicate(b, supplier, dummyUint64Predicate)
}
