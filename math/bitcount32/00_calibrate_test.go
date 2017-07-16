package bitcount32_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkRandomUintFuncUint32(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint32(b)
}
