package benchmark

import "testing"

func dummyBoolToInt32Func(_ bool) int32     { return 0 }
func dummyIntToInt32Func(_ int) int32       { return 0 }
func dummyInt8ToInt32Func(_ int8) int32     { return 0 }
func dummyInt16ToInt32Func(_ int16) int32   { return 0 }
func dummyInt32ToInt32Func(_ int32) int32   { return 0 }
func dummyInt64ToInt32Func(_ int64) int32   { return 0 }
func dummyUintToInt32Func(_ uint) int32     { return 0 }
func dummyUint8ToInt32Func(_ uint8) int32   { return 0 }
func dummyUint16ToInt32Func(_ uint16) int32 { return 0 }
func dummyUint32ToInt32Func(_ uint32) int32 { return 0 }
func dummyUint64ToInt32Func(_ uint64) int32 { return 0 }

// CalibrateBoolToInt32Func calibrates the BoolToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToInt32Func.
//
// ID: BC-5-1
func CalibrateBoolToInt32Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToInt32FuncCalibrated(supplier)
	BoolToInt32Func(b, supplier, dummyBoolToInt32Func)
}

// CalibrateIntToInt32Func calibrates the IntToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToInt32Func.
//
// ID: BC-5-2
func CalibrateIntToInt32Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToInt32FuncCalibrated(supplier)
	IntToInt32Func(b, supplier, dummyIntToInt32Func)
}

// CalibrateInt8ToInt32Func calibrates the Int8ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToInt32Func.
//
// ID: BC-5-3
func CalibrateInt8ToInt32Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToInt32FuncCalibrated(supplier)
	Int8ToInt32Func(b, supplier, dummyInt8ToInt32Func)
}

// CalibrateInt16ToInt32Func calibrates the Int16ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToInt32Func.
//
// ID: BC-5-4
func CalibrateInt16ToInt32Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToInt32FuncCalibrated(supplier)
	Int16ToInt32Func(b, supplier, dummyInt16ToInt32Func)
}

// CalibrateInt32ToInt32Func calibrates the Int32ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToInt32Func.
//
// ID: BC-5-5
func CalibrateInt32ToInt32Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToInt32FuncCalibrated(supplier)
	Int32ToInt32Func(b, supplier, dummyInt32ToInt32Func)
}

// CalibrateInt64ToInt32Func calibrates the Int64ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToInt32Func.
//
// ID: BC-5-6
func CalibrateInt64ToInt32Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToInt32FuncCalibrated(supplier)
	Int64ToInt32Func(b, supplier, dummyInt64ToInt32Func)
}

// CalibrateUintToInt32Func calibrates the UintToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToInt32Func.
//
// ID: BC-5-7
func CalibrateUintToInt32Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToInt32FuncCalibrated(supplier)
	UintToInt32Func(b, supplier, dummyUintToInt32Func)
}

// CalibrateUint8ToInt32Func calibrates the Uint8ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToInt32Func.
//
// ID: BC-5-8
func CalibrateUint8ToInt32Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToInt32FuncCalibrated(supplier)
	Uint8ToInt32Func(b, supplier, dummyUint8ToInt32Func)
}

// CalibrateUint16ToInt32Func calibrates the Uint16ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToInt32Func.
//
// ID: BC-5-9
func CalibrateUint16ToInt32Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToInt32FuncCalibrated(supplier)
	Uint16ToInt32Func(b, supplier, dummyUint16ToInt32Func)
}

// CalibrateUint32ToInt32Func calibrates the Uint32ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToInt32Func.
//
// ID: BC-5-10
func CalibrateUint32ToInt32Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToInt32FuncCalibrated(supplier)
	Uint32ToInt32Func(b, supplier, dummyUint32ToInt32Func)
}

// CalibrateUint64ToInt32Func calibrates the Uint64ToInt32Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToInt32Func.
//
// ID: BC-5-11
func CalibrateUint64ToInt32Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToInt32FuncCalibrated(supplier)
	Uint64ToInt32Func(b, supplier, dummyUint64ToInt32Func)
}
