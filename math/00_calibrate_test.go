package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func BenchmarkBoolFuncIntWithRandom(b *testing.B) {
	benchmark.CalibrateBoolFuncIntWithRandom(b)
}
