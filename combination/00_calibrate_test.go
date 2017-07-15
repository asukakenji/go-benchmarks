package combination_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkRandomUintSliceGenerator(b *testing.B) {
	benchmark.CalibrateRandomUintSliceGenerator(b)
}
