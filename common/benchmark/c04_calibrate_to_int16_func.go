package benchmark

import "testing"

func dummyBoolToInt16Func(_ bool) int16     { return 0 }
func dummyIntToInt16Func(_ int) int16       { return 0 }
func dummyInt8ToInt16Func(_ int8) int16     { return 0 }
func dummyInt16ToInt16Func(_ int16) int16   { return 0 }
func dummyInt32ToInt16Func(_ int32) int16   { return 0 }
func dummyInt64ToInt16Func(_ int64) int16   { return 0 }
func dummyUintToInt16Func(_ uint) int16     { return 0 }
func dummyUint8ToInt16Func(_ uint8) int16   { return 0 }
func dummyUint16ToInt16Func(_ uint16) int16 { return 0 }
func dummyUint32ToInt16Func(_ uint32) int16 { return 0 }
func dummyUint64ToInt16Func(_ uint64) int16 { return 0 }

// CalibrateBoolToInt16Func calibrates the BoolToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToInt16Func.
//
// ID: BC-4-1
func CalibrateBoolToInt16Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToInt16FuncCalibrated(supplier)
	BoolToInt16Func(b, supplier, dummyBoolToInt16Func)
}

// CalibrateIntToInt16Func calibrates the IntToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToInt16Func.
//
// ID: BC-4-2
func CalibrateIntToInt16Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToInt16FuncCalibrated(supplier)
	IntToInt16Func(b, supplier, dummyIntToInt16Func)
}

// CalibrateInt8ToInt16Func calibrates the Int8ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToInt16Func.
//
// ID: BC-4-3
func CalibrateInt8ToInt16Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToInt16FuncCalibrated(supplier)
	Int8ToInt16Func(b, supplier, dummyInt8ToInt16Func)
}

// CalibrateInt16ToInt16Func calibrates the Int16ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToInt16Func.
//
// ID: BC-4-4
func CalibrateInt16ToInt16Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToInt16FuncCalibrated(supplier)
	Int16ToInt16Func(b, supplier, dummyInt16ToInt16Func)
}

// CalibrateInt32ToInt16Func calibrates the Int32ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToInt16Func.
//
// ID: BC-4-5
func CalibrateInt32ToInt16Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToInt16FuncCalibrated(supplier)
	Int32ToInt16Func(b, supplier, dummyInt32ToInt16Func)
}

// CalibrateInt64ToInt16Func calibrates the Int64ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToInt16Func.
//
// ID: BC-4-6
func CalibrateInt64ToInt16Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToInt16FuncCalibrated(supplier)
	Int64ToInt16Func(b, supplier, dummyInt64ToInt16Func)
}

// CalibrateUintToInt16Func calibrates the UintToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToInt16Func.
//
// ID: BC-4-7
func CalibrateUintToInt16Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToInt16FuncCalibrated(supplier)
	UintToInt16Func(b, supplier, dummyUintToInt16Func)
}

// CalibrateUint8ToInt16Func calibrates the Uint8ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToInt16Func.
//
// ID: BC-4-8
func CalibrateUint8ToInt16Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToInt16FuncCalibrated(supplier)
	Uint8ToInt16Func(b, supplier, dummyUint8ToInt16Func)
}

// CalibrateUint16ToInt16Func calibrates the Uint16ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToInt16Func.
//
// ID: BC-4-9
func CalibrateUint16ToInt16Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToInt16FuncCalibrated(supplier)
	Uint16ToInt16Func(b, supplier, dummyUint16ToInt16Func)
}

// CalibrateUint32ToInt16Func calibrates the Uint32ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToInt16Func.
//
// ID: BC-4-10
func CalibrateUint32ToInt16Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToInt16FuncCalibrated(supplier)
	Uint32ToInt16Func(b, supplier, dummyUint32ToInt16Func)
}

// CalibrateUint64ToInt16Func calibrates the Uint64ToInt16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToInt16Func.
//
// ID: BC-4-11
func CalibrateUint64ToInt16Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToInt16FuncCalibrated(supplier)
	Uint64ToInt16Func(b, supplier, dummyUint64ToInt16Func)
}
