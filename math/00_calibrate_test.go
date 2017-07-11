package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkRandomUintFuncUint(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint(b)
}

func BenchmarkRandomUintFuncUint32(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint32(b)
}

func BenchmarkRandomUintFuncUint64(b *testing.B) {
	benchmark.CalibrateRandomUintFuncUint64(b)
}
