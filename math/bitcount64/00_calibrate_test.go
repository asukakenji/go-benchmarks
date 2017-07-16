package bitcount64_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkRandomUintFuncUint64(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint64(b)
}
