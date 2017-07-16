package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkUintFuncUintWithRandom(b *testing.B) {
	benchmark.CalibrateUintFuncUintWithRandom(b)
}

func BenchmarkUintFuncUint32WithRandom(b *testing.B) {
	benchmark.CalibrateUintFuncUint32WithRandom(b)
}

func BenchmarkUintFuncUint64(b *testing.B) {
	benchmark.CalibrateUintFuncUint64WithRandom(b)
}
