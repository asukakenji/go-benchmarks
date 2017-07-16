package combination_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkUintSliceGeneratorWithRandom(b *testing.B) {
	benchmark.CalibrateUintSliceGeneratorWithRandom(b)
}
