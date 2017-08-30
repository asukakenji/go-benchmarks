package benchmark

import "testing"

func dummyBoolToUint32Func(_ bool) uint32     { return 0 }
func dummyIntToUint32Func(_ int) uint32       { return 0 }
func dummyInt8ToUint32Func(_ int8) uint32     { return 0 }
func dummyInt16ToUint32Func(_ int16) uint32   { return 0 }
func dummyInt32ToUint32Func(_ int32) uint32   { return 0 }
func dummyInt64ToUint32Func(_ int64) uint32   { return 0 }
func dummyUintToUint32Func(_ uint) uint32     { return 0 }
func dummyUint8ToUint32Func(_ uint8) uint32   { return 0 }
func dummyUint16ToUint32Func(_ uint16) uint32 { return 0 }
func dummyUint32ToUint32Func(_ uint32) uint32 { return 0 }
func dummyUint64ToUint32Func(_ uint64) uint32 { return 0 }

// CalibrateBoolToUint32Func calibrates the BoolToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToUint32Func.
//
// ID: BC-10-1
func CalibrateBoolToUint32Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToUint32FuncCalibrated(supplier)
	BoolToUint32Func(b, supplier, dummyBoolToUint32Func)
}

// CalibrateIntToUint32Func calibrates the IntToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToUint32Func.
//
// ID: BC-10-2
func CalibrateIntToUint32Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToUint32FuncCalibrated(supplier)
	IntToUint32Func(b, supplier, dummyIntToUint32Func)
}

// CalibrateInt8ToUint32Func calibrates the Int8ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToUint32Func.
//
// ID: BC-10-3
func CalibrateInt8ToUint32Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToUint32FuncCalibrated(supplier)
	Int8ToUint32Func(b, supplier, dummyInt8ToUint32Func)
}

// CalibrateInt16ToUint32Func calibrates the Int16ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToUint32Func.
//
// ID: BC-10-4
func CalibrateInt16ToUint32Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToUint32FuncCalibrated(supplier)
	Int16ToUint32Func(b, supplier, dummyInt16ToUint32Func)
}

// CalibrateInt32ToUint32Func calibrates the Int32ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToUint32Func.
//
// ID: BC-10-5
func CalibrateInt32ToUint32Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToUint32FuncCalibrated(supplier)
	Int32ToUint32Func(b, supplier, dummyInt32ToUint32Func)
}

// CalibrateInt64ToUint32Func calibrates the Int64ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToUint32Func.
//
// ID: BC-10-6
func CalibrateInt64ToUint32Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToUint32FuncCalibrated(supplier)
	Int64ToUint32Func(b, supplier, dummyInt64ToUint32Func)
}

// CalibrateUintToUint32Func calibrates the UintToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToUint32Func.
//
// ID: BC-10-7
func CalibrateUintToUint32Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToUint32FuncCalibrated(supplier)
	UintToUint32Func(b, supplier, dummyUintToUint32Func)
}

// CalibrateUint8ToUint32Func calibrates the Uint8ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToUint32Func.
//
// ID: BC-10-8
func CalibrateUint8ToUint32Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToUint32FuncCalibrated(supplier)
	Uint8ToUint32Func(b, supplier, dummyUint8ToUint32Func)
}

// CalibrateUint16ToUint32Func calibrates the Uint16ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToUint32Func.
//
// ID: BC-10-9
func CalibrateUint16ToUint32Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToUint32FuncCalibrated(supplier)
	Uint16ToUint32Func(b, supplier, dummyUint16ToUint32Func)
}

// CalibrateUint32ToUint32Func calibrates the Uint32ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToUint32Func.
//
// ID: BC-10-10
func CalibrateUint32ToUint32Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToUint32FuncCalibrated(supplier)
	Uint32ToUint32Func(b, supplier, dummyUint32ToUint32Func)
}

// CalibrateUint64ToUint32Func calibrates the Uint64ToUint32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToUint32Func.
//
// ID: BC-10-11
func CalibrateUint64ToUint32Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToUint32FuncCalibrated(supplier)
	Uint64ToUint32Func(b, supplier, dummyUint64ToUint32Func)
}
