package benchmark

import "github.com/asukakenji/go-benchmarks/common/random"

var (
	genInt       = random.NewIntGenerator()
	genInt8      = random.NewInt8Generator()
	genInt16     = random.NewInt16Generator()
	genInt32     = random.NewInt32Generator()
	genInt64     = random.NewInt64Generator()
	genUint      = random.NewUintGenerator()
	genUint8     = random.NewUint8Generator()
	genUint16    = random.NewUint16Generator()
	genUint32    = random.NewUint32Generator()
	genUint64    = random.NewUint64Generator()
	genUintSlice = random.NewUintSliceGenerator()
)
