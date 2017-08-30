package benchmark

import "testing"

func dummyBoolToInt64Func(_ bool) int64     { return 0 }
func dummyIntToInt64Func(_ int) int64       { return 0 }
func dummyInt8ToInt64Func(_ int8) int64     { return 0 }
func dummyInt16ToInt64Func(_ int16) int64   { return 0 }
func dummyInt32ToInt64Func(_ int32) int64   { return 0 }
func dummyInt64ToInt64Func(_ int64) int64   { return 0 }
func dummyUintToInt64Func(_ uint) int64     { return 0 }
func dummyUint8ToInt64Func(_ uint8) int64   { return 0 }
func dummyUint16ToInt64Func(_ uint16) int64 { return 0 }
func dummyUint32ToInt64Func(_ uint32) int64 { return 0 }
func dummyUint64ToInt64Func(_ uint64) int64 { return 0 }

// CalibrateBoolToInt64Func calibrates the BoolToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToInt64Func.
//
// ID: BC-6-1
func CalibrateBoolToInt64Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToInt64FuncCalibrated(supplier)
	BoolToInt64Func(b, supplier, dummyBoolToInt64Func)
}

// CalibrateIntToInt64Func calibrates the IntToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToInt64Func.
//
// ID: BC-6-2
func CalibrateIntToInt64Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToInt64FuncCalibrated(supplier)
	IntToInt64Func(b, supplier, dummyIntToInt64Func)
}

// CalibrateInt8ToInt64Func calibrates the Int8ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToInt64Func.
//
// ID: BC-6-3
func CalibrateInt8ToInt64Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToInt64FuncCalibrated(supplier)
	Int8ToInt64Func(b, supplier, dummyInt8ToInt64Func)
}

// CalibrateInt16ToInt64Func calibrates the Int16ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToInt64Func.
//
// ID: BC-6-4
func CalibrateInt16ToInt64Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToInt64FuncCalibrated(supplier)
	Int16ToInt64Func(b, supplier, dummyInt16ToInt64Func)
}

// CalibrateInt32ToInt64Func calibrates the Int32ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToInt64Func.
//
// ID: BC-6-5
func CalibrateInt32ToInt64Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToInt64FuncCalibrated(supplier)
	Int32ToInt64Func(b, supplier, dummyInt32ToInt64Func)
}

// CalibrateInt64ToInt64Func calibrates the Int64ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToInt64Func.
//
// ID: BC-6-6
func CalibrateInt64ToInt64Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToInt64FuncCalibrated(supplier)
	Int64ToInt64Func(b, supplier, dummyInt64ToInt64Func)
}

// CalibrateUintToInt64Func calibrates the UintToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToInt64Func.
//
// ID: BC-6-7
func CalibrateUintToInt64Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToInt64FuncCalibrated(supplier)
	UintToInt64Func(b, supplier, dummyUintToInt64Func)
}

// CalibrateUint8ToInt64Func calibrates the Uint8ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToInt64Func.
//
// ID: BC-6-8
func CalibrateUint8ToInt64Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToInt64FuncCalibrated(supplier)
	Uint8ToInt64Func(b, supplier, dummyUint8ToInt64Func)
}

// CalibrateUint16ToInt64Func calibrates the Uint16ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToInt64Func.
//
// ID: BC-6-9
func CalibrateUint16ToInt64Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToInt64FuncCalibrated(supplier)
	Uint16ToInt64Func(b, supplier, dummyUint16ToInt64Func)
}

// CalibrateUint32ToInt64Func calibrates the Uint32ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToInt64Func.
//
// ID: BC-6-10
func CalibrateUint32ToInt64Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToInt64FuncCalibrated(supplier)
	Uint32ToInt64Func(b, supplier, dummyUint32ToInt64Func)
}

// CalibrateUint64ToInt64Func calibrates the Uint64ToInt64Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToInt64Func.
//
// ID: BC-6-11
func CalibrateUint64ToInt64Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToInt64FuncCalibrated(supplier)
	Uint64ToInt64Func(b, supplier, dummyUint64ToInt64Func)
}
