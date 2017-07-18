package b4_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b4"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b4.Permutation4ParamOrder0", b4.Permutation4ParamOrder0},
		{"b4.Permutation4ParamOrder1", b4.Permutation4ParamOrder1},
		{"b4.Permutation4ParamOrder2", b4.Permutation4ParamOrder2},
		{"b4.Permutation4ParamOrder3", b4.Permutation4ParamOrder3},
		{"b4.Permutation4ParamOrder4", b4.Permutation4ParamOrder4},
		{"b4.Permutation4ParamOrder5", b4.Permutation4ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}
