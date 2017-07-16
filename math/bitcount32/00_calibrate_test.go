package bitcount32_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkUintFuncUint32WithRandom(b *testing.B) {
	benchmark.CalibrateUintFuncUint32WithRandom(b)
}
