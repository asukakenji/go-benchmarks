package permutation_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func adaptor(x func(func([]uint), []uint)) func([]uint, func([]uint)) {
	return func(s []uint, f func([]uint)) {
		x(f, s)
	}
}

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"permutation.Permutation5Inc", adaptor(permutation.Permutation)},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}
