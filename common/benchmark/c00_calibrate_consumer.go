package benchmark

import "testing"

func dummyBoolConsumer(_ bool)     {}
func dummyIntConsumer(_ int)       {}
func dummyInt8Consumer(_ int8)     {}
func dummyInt16Consumer(_ int16)   {}
func dummyInt32Consumer(_ int32)   {}
func dummyInt64Consumer(_ int64)   {}
func dummyUintConsumer(_ uint)     {}
func dummyUint8Consumer(_ uint8)   {}
func dummyUint16Consumer(_ uint16) {}
func dummyUint32Consumer(_ uint32) {}
func dummyUint64Consumer(_ uint64) {}

// CalibrateBoolConsumer calibrates the BoolConsumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on BoolConsumer.
//
// ID: BC-0-1
func CalibrateBoolConsumer(b *testing.B, supplier func() bool) {
	if !isBoolSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setBoolConsumerCalibrated(supplier)
	BoolConsumer(b, supplier, dummyBoolConsumer)
}

// CalibrateIntConsumer calibrates the IntConsumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on IntConsumer.
//
// ID: BC-0-2
func CalibrateIntConsumer(b *testing.B, supplier func() int) {
	if !isIntSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setIntConsumerCalibrated(supplier)
	IntConsumer(b, supplier, dummyIntConsumer)
}

// CalibrateInt8Consumer calibrates the Int8Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int8Consumer.
//
// ID: BC-0-3
func CalibrateInt8Consumer(b *testing.B, supplier func() int8) {
	if !isInt8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt8ConsumerCalibrated(supplier)
	Int8Consumer(b, supplier, dummyInt8Consumer)
}

// CalibrateInt16Consumer calibrates the Int16Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int16Consumer.
//
// ID: BC-0-4
func CalibrateInt16Consumer(b *testing.B, supplier func() int16) {
	if !isInt16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt16ConsumerCalibrated(supplier)
	Int16Consumer(b, supplier, dummyInt16Consumer)
}

// CalibrateInt32Consumer calibrates the Int32Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int32Consumer.
//
// ID: BC-0-5
func CalibrateInt32Consumer(b *testing.B, supplier func() int32) {
	if !isInt32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt32ConsumerCalibrated(supplier)
	Int32Consumer(b, supplier, dummyInt32Consumer)
}

// CalibrateInt64Consumer calibrates the Int64Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Int64Consumer.
//
// ID: BC-0-6
func CalibrateInt64Consumer(b *testing.B, supplier func() int64) {
	if !isInt64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setInt64ConsumerCalibrated(supplier)
	Int64Consumer(b, supplier, dummyInt64Consumer)
}

// CalibrateUintConsumer calibrates the UintConsumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on UintConsumer.
//
// ID: BC-0-7
func CalibrateUintConsumer(b *testing.B, supplier func() uint) {
	if !isUintSupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUintConsumerCalibrated(supplier)
	UintConsumer(b, supplier, dummyUintConsumer)
}

// CalibrateUint8Consumer calibrates the Uint8Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint8Consumer.
//
// ID: BC-0-8
func CalibrateUint8Consumer(b *testing.B, supplier func() uint8) {
	if !isUint8SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint8ConsumerCalibrated(supplier)
	Uint8Consumer(b, supplier, dummyUint8Consumer)
}

// CalibrateUint16Consumer calibrates the Uint16Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint16Consumer.
//
// ID: BC-0-9
func CalibrateUint16Consumer(b *testing.B, supplier func() uint16) {
	if !isUint16SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint16ConsumerCalibrated(supplier)
	Uint16Consumer(b, supplier, dummyUint16Consumer)
}

// CalibrateUint32Consumer calibrates the Uint32Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint32Consumer.
//
// ID: BC-0-10
func CalibrateUint32Consumer(b *testing.B, supplier func() uint32) {
	if !isUint32SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint32ConsumerCalibrated(supplier)
	Uint32Consumer(b, supplier, dummyUint32Consumer)
}

// CalibrateUint64Consumer calibrates the Uint64Consumer function.
// It MUST be called at least once before all benchmarks if any benchmark
// depends on Uint64Consumer.
//
// ID: BC-0-11
func CalibrateUint64Consumer(b *testing.B, supplier func() uint64) {
	if !isUint64SupplierCalibrated(supplier) {
		panic("supplier function not calibrated")
	}
	setUint64ConsumerCalibrated(supplier)
	Uint64Consumer(b, supplier, dummyUint64Consumer)
}
