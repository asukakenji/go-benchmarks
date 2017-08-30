package benchmark

import "testing"

func dummyBoolToUintFunc(_ bool) uint     { return 0 }
func dummyIntToUintFunc(_ int) uint       { return 0 }
func dummyInt8ToUintFunc(_ int8) uint     { return 0 }
func dummyInt16ToUintFunc(_ int16) uint   { return 0 }
func dummyInt32ToUintFunc(_ int32) uint   { return 0 }
func dummyInt64ToUintFunc(_ int64) uint   { return 0 }
func dummyUintToUintFunc(_ uint) uint     { return 0 }
func dummyUint8ToUintFunc(_ uint8) uint   { return 0 }
func dummyUint16ToUintFunc(_ uint16) uint { return 0 }
func dummyUint32ToUintFunc(_ uint32) uint { return 0 }
func dummyUint64ToUintFunc(_ uint64) uint { return 0 }

// CalibrateBoolToUintFunc calibrates the BoolToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToUintFunc.
//
// ID: BC-7-1
func CalibrateBoolToUintFunc(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToUintFuncCalibrated(supplier)
	BoolToUintFunc(b, supplier, dummyBoolToUintFunc)
}

// CalibrateIntToUintFunc calibrates the IntToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToUintFunc.
//
// ID: BC-7-2
func CalibrateIntToUintFunc(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToUintFuncCalibrated(supplier)
	IntToUintFunc(b, supplier, dummyIntToUintFunc)
}

// CalibrateInt8ToUintFunc calibrates the Int8ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToUintFunc.
//
// ID: BC-7-3
func CalibrateInt8ToUintFunc(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToUintFuncCalibrated(supplier)
	Int8ToUintFunc(b, supplier, dummyInt8ToUintFunc)
}

// CalibrateInt16ToUintFunc calibrates the Int16ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToUintFunc.
//
// ID: BC-7-4
func CalibrateInt16ToUintFunc(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToUintFuncCalibrated(supplier)
	Int16ToUintFunc(b, supplier, dummyInt16ToUintFunc)
}

// CalibrateInt32ToUintFunc calibrates the Int32ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToUintFunc.
//
// ID: BC-7-5
func CalibrateInt32ToUintFunc(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToUintFuncCalibrated(supplier)
	Int32ToUintFunc(b, supplier, dummyInt32ToUintFunc)
}

// CalibrateInt64ToUintFunc calibrates the Int64ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToUintFunc.
//
// ID: BC-7-6
func CalibrateInt64ToUintFunc(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToUintFuncCalibrated(supplier)
	Int64ToUintFunc(b, supplier, dummyInt64ToUintFunc)
}

// CalibrateUintToUintFunc calibrates the UintToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToUintFunc.
//
// ID: BC-7-7
func CalibrateUintToUintFunc(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToUintFuncCalibrated(supplier)
	UintToUintFunc(b, supplier, dummyUintToUintFunc)
}

// CalibrateUint8ToUintFunc calibrates the Uint8ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToUintFunc.
//
// ID: BC-7-8
func CalibrateUint8ToUintFunc(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToUintFuncCalibrated(supplier)
	Uint8ToUintFunc(b, supplier, dummyUint8ToUintFunc)
}

// CalibrateUint16ToUintFunc calibrates the Uint16ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToUintFunc.
//
// ID: BC-7-9
func CalibrateUint16ToUintFunc(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToUintFuncCalibrated(supplier)
	Uint16ToUintFunc(b, supplier, dummyUint16ToUintFunc)
}

// CalibrateUint32ToUintFunc calibrates the Uint32ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToUintFunc.
//
// ID: BC-7-10
func CalibrateUint32ToUintFunc(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToUintFuncCalibrated(supplier)
	Uint32ToUintFunc(b, supplier, dummyUint32ToUintFunc)
}

// CalibrateUint64ToUintFunc calibrates the Uint64ToUintFunc function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToUintFunc.
//
// ID: BC-7-11
func CalibrateUint64ToUintFunc(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToUintFuncCalibrated(supplier)
	Uint64ToUintFunc(b, supplier, dummyUint64ToUintFunc)
}
