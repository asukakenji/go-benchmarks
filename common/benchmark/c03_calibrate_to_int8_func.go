package benchmark

import "testing"

func dummyBoolToInt8Func(_ bool) int8     { return 0 }
func dummyIntToInt8Func(_ int) int8       { return 0 }
func dummyInt8ToInt8Func(_ int8) int8     { return 0 }
func dummyInt16ToInt8Func(_ int16) int8   { return 0 }
func dummyInt32ToInt8Func(_ int32) int8   { return 0 }
func dummyInt64ToInt8Func(_ int64) int8   { return 0 }
func dummyUintToInt8Func(_ uint) int8     { return 0 }
func dummyUint8ToInt8Func(_ uint8) int8   { return 0 }
func dummyUint16ToInt8Func(_ uint16) int8 { return 0 }
func dummyUint32ToInt8Func(_ uint32) int8 { return 0 }
func dummyUint64ToInt8Func(_ uint64) int8 { return 0 }

// CalibrateBoolToInt8Func calibrates the BoolToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToInt8Func.
//
// ID: BC-3-1
func CalibrateBoolToInt8Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToInt8FuncCalibrated(supplier)
	BoolToInt8Func(b, supplier, dummyBoolToInt8Func)
}

// CalibrateIntToInt8Func calibrates the IntToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToInt8Func.
//
// ID: BC-3-2
func CalibrateIntToInt8Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToInt8FuncCalibrated(supplier)
	IntToInt8Func(b, supplier, dummyIntToInt8Func)
}

// CalibrateInt8ToInt8Func calibrates the Int8ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToInt8Func.
//
// ID: BC-3-3
func CalibrateInt8ToInt8Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToInt8FuncCalibrated(supplier)
	Int8ToInt8Func(b, supplier, dummyInt8ToInt8Func)
}

// CalibrateInt16ToInt8Func calibrates the Int16ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToInt8Func.
//
// ID: BC-3-4
func CalibrateInt16ToInt8Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToInt8FuncCalibrated(supplier)
	Int16ToInt8Func(b, supplier, dummyInt16ToInt8Func)
}

// CalibrateInt32ToInt8Func calibrates the Int32ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToInt8Func.
//
// ID: BC-3-5
func CalibrateInt32ToInt8Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToInt8FuncCalibrated(supplier)
	Int32ToInt8Func(b, supplier, dummyInt32ToInt8Func)
}

// CalibrateInt64ToInt8Func calibrates the Int64ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToInt8Func.
//
// ID: BC-3-6
func CalibrateInt64ToInt8Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToInt8FuncCalibrated(supplier)
	Int64ToInt8Func(b, supplier, dummyInt64ToInt8Func)
}

// CalibrateUintToInt8Func calibrates the UintToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToInt8Func.
//
// ID: BC-3-7
func CalibrateUintToInt8Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToInt8FuncCalibrated(supplier)
	UintToInt8Func(b, supplier, dummyUintToInt8Func)
}

// CalibrateUint8ToInt8Func calibrates the Uint8ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToInt8Func.
//
// ID: BC-3-8
func CalibrateUint8ToInt8Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToInt8FuncCalibrated(supplier)
	Uint8ToInt8Func(b, supplier, dummyUint8ToInt8Func)
}

// CalibrateUint16ToInt8Func calibrates the Uint16ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToInt8Func.
//
// ID: BC-3-9
func CalibrateUint16ToInt8Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToInt8FuncCalibrated(supplier)
	Uint16ToInt8Func(b, supplier, dummyUint16ToInt8Func)
}

// CalibrateUint32ToInt8Func calibrates the Uint32ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToInt8Func.
//
// ID: BC-3-10
func CalibrateUint32ToInt8Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToInt8FuncCalibrated(supplier)
	Uint32ToInt8Func(b, supplier, dummyUint32ToInt8Func)
}

// CalibrateUint64ToInt8Func calibrates the Uint64ToInt8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToInt8Func.
//
// ID: BC-3-11
func CalibrateUint64ToInt8Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToInt8FuncCalibrated(supplier)
	Uint64ToInt8Func(b, supplier, dummyUint64ToInt8Func)
}
