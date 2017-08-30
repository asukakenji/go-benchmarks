package benchmark

import "testing"

func dummyBoolToUint8Func(_ bool) uint8     { return 0 }
func dummyIntToUint8Func(_ int) uint8       { return 0 }
func dummyInt8ToUint8Func(_ int8) uint8     { return 0 }
func dummyInt16ToUint8Func(_ int16) uint8   { return 0 }
func dummyInt32ToUint8Func(_ int32) uint8   { return 0 }
func dummyInt64ToUint8Func(_ int64) uint8   { return 0 }
func dummyUintToUint8Func(_ uint) uint8     { return 0 }
func dummyUint8ToUint8Func(_ uint8) uint8   { return 0 }
func dummyUint16ToUint8Func(_ uint16) uint8 { return 0 }
func dummyUint32ToUint8Func(_ uint32) uint8 { return 0 }
func dummyUint64ToUint8Func(_ uint64) uint8 { return 0 }

// CalibrateBoolToUint8Func calibrates the BoolToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToUint8Func.
//
// ID: BC-8-1
func CalibrateBoolToUint8Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToUint8FuncCalibrated(supplier)
	BoolToUint8Func(b, supplier, dummyBoolToUint8Func)
}

// CalibrateIntToUint8Func calibrates the IntToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToUint8Func.
//
// ID: BC-8-2
func CalibrateIntToUint8Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToUint8FuncCalibrated(supplier)
	IntToUint8Func(b, supplier, dummyIntToUint8Func)
}

// CalibrateInt8ToUint8Func calibrates the Int8ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToUint8Func.
//
// ID: BC-8-3
func CalibrateInt8ToUint8Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToUint8FuncCalibrated(supplier)
	Int8ToUint8Func(b, supplier, dummyInt8ToUint8Func)
}

// CalibrateInt16ToUint8Func calibrates the Int16ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToUint8Func.
//
// ID: BC-8-4
func CalibrateInt16ToUint8Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToUint8FuncCalibrated(supplier)
	Int16ToUint8Func(b, supplier, dummyInt16ToUint8Func)
}

// CalibrateInt32ToUint8Func calibrates the Int32ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToUint8Func.
//
// ID: BC-8-5
func CalibrateInt32ToUint8Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToUint8FuncCalibrated(supplier)
	Int32ToUint8Func(b, supplier, dummyInt32ToUint8Func)
}

// CalibrateInt64ToUint8Func calibrates the Int64ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToUint8Func.
//
// ID: BC-8-6
func CalibrateInt64ToUint8Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToUint8FuncCalibrated(supplier)
	Int64ToUint8Func(b, supplier, dummyInt64ToUint8Func)
}

// CalibrateUintToUint8Func calibrates the UintToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToUint8Func.
//
// ID: BC-8-7
func CalibrateUintToUint8Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToUint8FuncCalibrated(supplier)
	UintToUint8Func(b, supplier, dummyUintToUint8Func)
}

// CalibrateUint8ToUint8Func calibrates the Uint8ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToUint8Func.
//
// ID: BC-8-8
func CalibrateUint8ToUint8Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToUint8FuncCalibrated(supplier)
	Uint8ToUint8Func(b, supplier, dummyUint8ToUint8Func)
}

// CalibrateUint16ToUint8Func calibrates the Uint16ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToUint8Func.
//
// ID: BC-8-9
func CalibrateUint16ToUint8Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToUint8FuncCalibrated(supplier)
	Uint16ToUint8Func(b, supplier, dummyUint16ToUint8Func)
}

// CalibrateUint32ToUint8Func calibrates the Uint32ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToUint8Func.
//
// ID: BC-8-10
func CalibrateUint32ToUint8Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToUint8FuncCalibrated(supplier)
	Uint32ToUint8Func(b, supplier, dummyUint32ToUint8Func)
}

// CalibrateUint64ToUint8Func calibrates the Uint64ToUint8Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToUint8Func.
//
// ID: BC-8-11
func CalibrateUint64ToUint8Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToUint8FuncCalibrated(supplier)
	Uint64ToUint8Func(b, supplier, dummyUint64ToUint8Func)
}
