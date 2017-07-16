package bitcount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkUintFuncUintWithRandom(b *testing.B) {
	benchmark.CalibrateUintFuncUintWithRandom(b)
}
