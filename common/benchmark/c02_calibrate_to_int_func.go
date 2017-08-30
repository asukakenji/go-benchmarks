package benchmark

import "testing"

func dummyBoolToIntFunc(_ bool) int     { return 0 }
func dummyIntToIntFunc(_ int) int       { return 0 }
func dummyInt8ToIntFunc(_ int8) int     { return 0 }
func dummyInt16ToIntFunc(_ int16) int   { return 0 }
func dummyInt32ToIntFunc(_ int32) int   { return 0 }
func dummyInt64ToIntFunc(_ int64) int   { return 0 }
func dummyUintToIntFunc(_ uint) int     { return 0 }
func dummyUint8ToIntFunc(_ uint8) int   { return 0 }
func dummyUint16ToIntFunc(_ uint16) int { return 0 }
func dummyUint32ToIntFunc(_ uint32) int { return 0 }
func dummyUint64ToIntFunc(_ uint64) int { return 0 }

// CalibrateBoolToIntFunc calibrates the BoolToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToIntFunc.
//
// ID: BC-2-1
func CalibrateBoolToIntFunc(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToIntFuncCalibrated(supplier)
	BoolToIntFunc(b, supplier, dummyBoolToIntFunc)
}

// CalibrateIntToIntFunc calibrates the IntToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToIntFunc.
//
// ID: BC-2-2
func CalibrateIntToIntFunc(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToIntFuncCalibrated(supplier)
	IntToIntFunc(b, supplier, dummyIntToIntFunc)
}

// CalibrateInt8ToIntFunc calibrates the Int8ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToIntFunc.
//
// ID: BC-2-3
func CalibrateInt8ToIntFunc(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToIntFuncCalibrated(supplier)
	Int8ToIntFunc(b, supplier, dummyInt8ToIntFunc)
}

// CalibrateInt16ToIntFunc calibrates the Int16ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToIntFunc.
//
// ID: BC-2-4
func CalibrateInt16ToIntFunc(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToIntFuncCalibrated(supplier)
	Int16ToIntFunc(b, supplier, dummyInt16ToIntFunc)
}

// CalibrateInt32ToIntFunc calibrates the Int32ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToIntFunc.
//
// ID: BC-2-5
func CalibrateInt32ToIntFunc(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToIntFuncCalibrated(supplier)
	Int32ToIntFunc(b, supplier, dummyInt32ToIntFunc)
}

// CalibrateInt64ToIntFunc calibrates the Int64ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToIntFunc.
//
// ID: BC-2-6
func CalibrateInt64ToIntFunc(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToIntFuncCalibrated(supplier)
	Int64ToIntFunc(b, supplier, dummyInt64ToIntFunc)
}

// CalibrateUintToIntFunc calibrates the UintToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToIntFunc.
//
// ID: BC-2-7
func CalibrateUintToIntFunc(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToIntFuncCalibrated(supplier)
	UintToIntFunc(b, supplier, dummyUintToIntFunc)
}

// CalibrateUint8ToIntFunc calibrates the Uint8ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToIntFunc.
//
// ID: BC-2-8
func CalibrateUint8ToIntFunc(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToIntFuncCalibrated(supplier)
	Uint8ToIntFunc(b, supplier, dummyUint8ToIntFunc)
}

// CalibrateUint16ToIntFunc calibrates the Uint16ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToIntFunc.
//
// ID: BC-2-9
func CalibrateUint16ToIntFunc(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToIntFuncCalibrated(supplier)
	Uint16ToIntFunc(b, supplier, dummyUint16ToIntFunc)
}

// CalibrateUint32ToIntFunc calibrates the Uint32ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToIntFunc.
//
// ID: BC-2-10
func CalibrateUint32ToIntFunc(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToIntFuncCalibrated(supplier)
	Uint32ToIntFunc(b, supplier, dummyUint32ToIntFunc)
}

// CalibrateUint64ToIntFunc calibrates the Uint64ToIntFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToIntFunc.
//
// ID: BC-2-11
func CalibrateUint64ToIntFunc(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToIntFuncCalibrated(supplier)
	Uint64ToIntFunc(b, supplier, dummyUint64ToIntFunc)
}
