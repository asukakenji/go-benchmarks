package benchmark

import "testing"

func dummyBoolToUint64Func(_ bool) uint64     { return 0 }
func dummyIntToUint64Func(_ int) uint64       { return 0 }
func dummyInt8ToUint64Func(_ int8) uint64     { return 0 }
func dummyInt16ToUint64Func(_ int16) uint64   { return 0 }
func dummyInt32ToUint64Func(_ int32) uint64   { return 0 }
func dummyInt64ToUint64Func(_ int64) uint64   { return 0 }
func dummyUintToUint64Func(_ uint) uint64     { return 0 }
func dummyUint8ToUint64Func(_ uint8) uint64   { return 0 }
func dummyUint16ToUint64Func(_ uint16) uint64 { return 0 }
func dummyUint32ToUint64Func(_ uint32) uint64 { return 0 }
func dummyUint64ToUint64Func(_ uint64) uint64 { return 0 }

// CalibrateBoolToUint64Func calibrates the BoolToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToUint64Func.
//
// ID: BC-11-1
func CalibrateBoolToUint64Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToUint64FuncCalibrated(supplier)
	BoolToUint64Func(b, supplier, dummyBoolToUint64Func)
}

// CalibrateIntToUint64Func calibrates the IntToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToUint64Func.
//
// ID: BC-11-2
func CalibrateIntToUint64Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToUint64FuncCalibrated(supplier)
	IntToUint64Func(b, supplier, dummyIntToUint64Func)
}

// CalibrateInt8ToUint64Func calibrates the Int8ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToUint64Func.
//
// ID: BC-11-3
func CalibrateInt8ToUint64Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToUint64FuncCalibrated(supplier)
	Int8ToUint64Func(b, supplier, dummyInt8ToUint64Func)
}

// CalibrateInt16ToUint64Func calibrates the Int16ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToUint64Func.
//
// ID: BC-11-4
func CalibrateInt16ToUint64Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToUint64FuncCalibrated(supplier)
	Int16ToUint64Func(b, supplier, dummyInt16ToUint64Func)
}

// CalibrateInt32ToUint64Func calibrates the Int32ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToUint64Func.
//
// ID: BC-11-5
func CalibrateInt32ToUint64Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToUint64FuncCalibrated(supplier)
	Int32ToUint64Func(b, supplier, dummyInt32ToUint64Func)
}

// CalibrateInt64ToUint64Func calibrates the Int64ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToUint64Func.
//
// ID: BC-11-6
func CalibrateInt64ToUint64Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToUint64FuncCalibrated(supplier)
	Int64ToUint64Func(b, supplier, dummyInt64ToUint64Func)
}

// CalibrateUintToUint64Func calibrates the UintToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToUint64Func.
//
// ID: BC-11-7
func CalibrateUintToUint64Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToUint64FuncCalibrated(supplier)
	UintToUint64Func(b, supplier, dummyUintToUint64Func)
}

// CalibrateUint8ToUint64Func calibrates the Uint8ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToUint64Func.
//
// ID: BC-11-8
func CalibrateUint8ToUint64Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToUint64FuncCalibrated(supplier)
	Uint8ToUint64Func(b, supplier, dummyUint8ToUint64Func)
}

// CalibrateUint16ToUint64Func calibrates the Uint16ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToUint64Func.
//
// ID: BC-11-9
func CalibrateUint16ToUint64Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToUint64FuncCalibrated(supplier)
	Uint16ToUint64Func(b, supplier, dummyUint16ToUint64Func)
}

// CalibrateUint32ToUint64Func calibrates the Uint32ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToUint64Func.
//
// ID: BC-11-10
func CalibrateUint32ToUint64Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToUint64FuncCalibrated(supplier)
	Uint32ToUint64Func(b, supplier, dummyUint32ToUint64Func)
}

// CalibrateUint64ToUint64Func calibrates the Uint64ToUint64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToUint64Func.
//
// ID: BC-11-11
func CalibrateUint64ToUint64Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToUint64FuncCalibrated(supplier)
	Uint64ToUint64Func(b, supplier, dummyUint64ToUint64Func)
}
