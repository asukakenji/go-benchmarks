package onescount_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkIntFuncUintWithRandom(b *testing.B) {
	benchmark.CalibrateAnyFuncUintWithRandom(b)
}

func BenchmarkIntFuncUint32WithRandom(b *testing.B) {
	benchmark.CalibrateAnyFuncUint32WithRandom(b)
}

func BenchmarkIntFuncUint64WithRandom(b *testing.B) {
	benchmark.CalibrateAnyFuncUint64WithRandom(b)
}
