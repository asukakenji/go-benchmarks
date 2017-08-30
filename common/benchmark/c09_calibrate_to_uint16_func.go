package benchmark

import "testing"

func dummyBoolToUint16Func(_ bool) uint16     { return 0 }
func dummyIntToUint16Func(_ int) uint16       { return 0 }
func dummyInt8ToUint16Func(_ int8) uint16     { return 0 }
func dummyInt16ToUint16Func(_ int16) uint16   { return 0 }
func dummyInt32ToUint16Func(_ int32) uint16   { return 0 }
func dummyInt64ToUint16Func(_ int64) uint16   { return 0 }
func dummyUintToUint16Func(_ uint) uint16     { return 0 }
func dummyUint8ToUint16Func(_ uint8) uint16   { return 0 }
func dummyUint16ToUint16Func(_ uint16) uint16 { return 0 }
func dummyUint32ToUint16Func(_ uint32) uint16 { return 0 }
func dummyUint64ToUint16Func(_ uint64) uint16 { return 0 }

// CalibrateBoolToUint16Func calibrates the BoolToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolToUint16Func.
//
// ID: BC-9-1
func CalibrateBoolToUint16Func(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolToUint16FuncCalibrated(supplier)
	BoolToUint16Func(b, supplier, dummyBoolToUint16Func)
}

// CalibrateIntToUint16Func calibrates the IntToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntToUint16Func.
//
// ID: BC-9-2
func CalibrateIntToUint16Func(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntToUint16FuncCalibrated(supplier)
	IntToUint16Func(b, supplier, dummyIntToUint16Func)
}

// CalibrateInt8ToUint16Func calibrates the Int8ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8ToUint16Func.
//
// ID: BC-9-3
func CalibrateInt8ToUint16Func(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ToUint16FuncCalibrated(supplier)
	Int8ToUint16Func(b, supplier, dummyInt8ToUint16Func)
}

// CalibrateInt16ToUint16Func calibrates the Int16ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16ToUint16Func.
//
// ID: BC-9-4
func CalibrateInt16ToUint16Func(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ToUint16FuncCalibrated(supplier)
	Int16ToUint16Func(b, supplier, dummyInt16ToUint16Func)
}

// CalibrateInt32ToUint16Func calibrates the Int32ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32ToUint16Func.
//
// ID: BC-9-5
func CalibrateInt32ToUint16Func(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ToUint16FuncCalibrated(supplier)
	Int32ToUint16Func(b, supplier, dummyInt32ToUint16Func)
}

// CalibrateInt64ToUint16Func calibrates the Int64ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64ToUint16Func.
//
// ID: BC-9-6
func CalibrateInt64ToUint16Func(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ToUint16FuncCalibrated(supplier)
	Int64ToUint16Func(b, supplier, dummyInt64ToUint16Func)
}

// CalibrateUintToUint16Func calibrates the UintToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintToUint16Func.
//
// ID: BC-9-7
func CalibrateUintToUint16Func(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintToUint16FuncCalibrated(supplier)
	UintToUint16Func(b, supplier, dummyUintToUint16Func)
}

// CalibrateUint8ToUint16Func calibrates the Uint8ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8ToUint16Func.
//
// ID: BC-9-8
func CalibrateUint8ToUint16Func(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ToUint16FuncCalibrated(supplier)
	Uint8ToUint16Func(b, supplier, dummyUint8ToUint16Func)
}

// CalibrateUint16ToUint16Func calibrates the Uint16ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16ToUint16Func.
//
// ID: BC-9-9
func CalibrateUint16ToUint16Func(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ToUint16FuncCalibrated(supplier)
	Uint16ToUint16Func(b, supplier, dummyUint16ToUint16Func)
}

// CalibrateUint32ToUint16Func calibrates the Uint32ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32ToUint16Func.
//
// ID: BC-9-10
func CalibrateUint32ToUint16Func(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ToUint16FuncCalibrated(supplier)
	Uint32ToUint16Func(b, supplier, dummyUint32ToUint16Func)
}

// CalibrateUint64ToUint16Func calibrates the Uint64ToUint16Func function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64ToUint16Func.
//
// ID: BC-9-11
func CalibrateUint64ToUint16Func(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ToUint16FuncCalibrated(supplier)
	Uint64ToUint16Func(b, supplier, dummyUint64ToUint16Func)
}
