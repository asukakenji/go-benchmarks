package bitcount64_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkUintFuncUint64WithRandom(b *testing.B) {
	benchmark.CalibrateUintFuncUint64WithRandom(b)
}
