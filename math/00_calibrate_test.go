package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/benchmark"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
)

var intSupplier = randomsupplier.NewInt()

func BenchmarkCalibrateSupplier(b *testing.B) {
	benchmark.IntSupplier(b, intSupplier.Next)
}

func BenchmarkCalibrateBenchmarker(b *testing.B) {
	benchmark.CalibrateIntPredicate(b, intSupplier.Next)
}
