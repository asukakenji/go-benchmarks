package benchmark

import "github.com/asukakenji/go-benchmarks/common/random"

var (
	genUint      = random.NewUintGenerator()
	genUint32    = random.NewUint32Generator()
	genUint64    = random.NewUint64Generator()
	genUintSlice = random.NewUintSliceGenerator()
)
