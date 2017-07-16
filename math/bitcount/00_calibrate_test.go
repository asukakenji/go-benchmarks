package bitcount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkRandomUintFuncUint(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint(b)
}
